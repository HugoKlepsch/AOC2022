package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type File struct {
	Name     string
	Size     int64
	Parent   *File
	IsDir    bool
	Children map[string]*File
}

func (f *File) getRoot() *File {
	var currentFile *File = f
	for {
		if currentFile.Name == "/" {
			return currentFile
		} else if currentFile.Parent == nil {
			panic("Parent nil and did not find root")
		} else {
			currentFile = currentFile.Parent
		}
	}
}

func (f *File) FindFile(name string) *File {
	if name == "/" {
		return f.getRoot()
	}
	if name == ".." {
		return f.Parent
	}
	for _, file := range f.Children {
		if file.Name == name {
			return file
		}
	}
	return nil
}

func parseCommandOutput(currentDir *File, line string) {
	matches := dirRegex.FindAllStringSubmatch(line, -1)
	if matches != nil && len(matches) == 1 && len(matches[0]) == 3 {
		// this line is a "dir X" line
		newDir := File{
			Name:     matches[0][2],
			Size:     0,
			Parent:   currentDir,
			IsDir:    true,
			Children: make(map[string]*File),
		}
		if _, ok := currentDir.Children[newDir.Name]; !ok {
			// Only add a new file if we haven't seen it already
			currentDir.Children[newDir.Name] = &newDir
		}
		return
	}
	matches = fileRegex.FindAllStringSubmatch(line, -1)
	if matches != nil && len(matches) == 1 && len(matches[0]) == 3 {
		// this line is a "<num> <file>" line
		size, err := strconv.ParseInt(matches[0][1], 10, 64)
		if err != nil {
			panic(err)
		}
		newFile := File{
			Name:     matches[0][2],
			Size:     size,
			Parent:   currentDir,
			IsDir:    false,
			Children: nil,
		}
		if _, ok := currentDir.Children[newFile.Name]; !ok {
			// Only add a new file if we haven't seen it already
			currentDir.Children[newFile.Name] = &newFile
		}
		return
	}

	panic("no valid parse")
}

func parseCommand(currentDir **File, line string) {
	matches := commandRegex.FindAllStringSubmatch(line, -1)
	if matches == nil || len(matches) != 1 || len(matches[0]) != 3 {
		err := fmt.Errorf("failed to parse command. matches: %v, line %s", matches, line)
		panic(err)
	}
	command := matches[0][1]

	if command == "cd" {
		arg := matches[0][2]
		newDir := (*currentDir).FindFile(arg)
		*currentDir = newDir
	}
}

func (f *File) PrintTreeNode(indent int, recursive bool) {
	for i := 0; i < indent; i++ {
		fmt.Print("  ")
	}
	if f.IsDir {
		fmt.Printf("%s (dir)\n", f.Name)
		if recursive {
			for _, child := range f.Children {
				child.PrintTreeNode(indent+1, recursive)
			}
		}
	} else {
		fmt.Printf("%s (file, size=%d)\n", f.Name, f.Size)
	}
}

func (f *File) recursiveSize() int64 {
	if f.IsDir {
		var totalSize int64 = 0
		for _, child := range f.Children {
			totalSize += child.recursiveSize()
		}
		return totalSize
	} else {
		return f.Size
	}
}

func (f *File) dirsUnderSize(size int64) []*File {
	files := make([]*File, 0)

	if f.IsDir {
		dirSize := f.recursiveSize()
		if dirSize <= size {
			f.PrintTreeNode(0, false)
			files = append(files, f)
		}

		for _, child := range f.Children {
			files = append(files, child.dirsUnderSize(size)...)
		}
	}
	return files
}

var commandRegex = regexp.MustCompile(`^\$ ([a-zA-Z]+) ?([a-zA-Z0-9/.]+)?`)

var dirRegex = regexp.MustCompile(`^(dir) ([^ ]+)`)

var fileRegex = regexp.MustCompile(`^([0-9]+) ([^ ]+)`)

func main() {

	readFile, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	root := File{
		Name:     "/",
		Size:     0,
		Parent:   nil,
		IsDir:    true,
		Children: make(map[string]*File),
	}

	var currentDir *File = &root

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if strings.HasPrefix(line, "$") {
			parseCommand(&currentDir, line)
		} else {
			parseCommandOutput(currentDir, line)
		}
	}
	if err = fileScanner.Err(); err != nil {
		fmt.Println(err)
		panic(err)
	}

	root.PrintTreeNode(0, true)

	dirsUnderSize := root.dirsUnderSize(100000)
	var total int64 = 0
	for _, dir := range dirsUnderSize {
		size := dir.recursiveSize()
		fmt.Printf("%v (recursive size=%d)\n", dir.Name, size)
		total += size
	}
	fmt.Printf("%d\n", total)
}
