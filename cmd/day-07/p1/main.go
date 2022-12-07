package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type File struct {
	Name     string
	Size     int64
	Parent   *File
	IsDir    bool
	Children []*File
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

}

func parseCommand(currentDir **File, line string) {

}

var commandRegex = regexp.MustCompile(`$ ([a-zA-Z]+) ([a-zA-Z0-9]+)`)

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
		Children: make([]*File, 0),
	}

	var currentDir *File = &root

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if strings.HasPrefix(line, "$") {

		} else {
			parseCommandOutput(currentDir, line)
		}
	}
	if err = fileScanner.Err(); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
