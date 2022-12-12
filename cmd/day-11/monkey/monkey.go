package monkey

import (
	"AOC2022/cmd/day-11/ringbuffer"
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var (
	startingItemsRegex    = regexp.MustCompile(`([0-9]+),? ?`)
	operationRegex        = regexp.MustCompile(`Operation: new = old (?:(\+) ([0-9]+))|(?:(\*) ([0-9]+)|(old))$`)
	ErrInvalidMonkeyParse = errors.New("invalid monkey parse")
)

type Monkey struct {
	MonkeyNum       int
	Items           ringbuffer.RingBuffer[Item]
	NItemsInspected int
	Operation       Operation
	Test            MonkeyTest
}

func ParseLinesToMonkey(lines []string) (*Monkey, error) {
	m := Monkey{
		Items: *ringbuffer.New[Item](100),
	}

	_, err := fmt.Sscanf(lines[0], "Monkey %d", &m.MonkeyNum)
	if err != nil {
		return nil, fmt.Errorf("monkey num: %w: %v", ErrInvalidMonkeyParse, err.Error())
	}

	matches := startingItemsRegex.FindAllStringSubmatch(lines[1], -1)
	if matches == nil || len(matches) < 1 || len(matches[0]) != 2 {
		return nil, ErrInvalidMonkeyParse
	}
	for i := 0; i < len(matches); i++ {
		item, err := strconv.ParseInt(matches[i][1], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("item: %w: %v", ErrInvalidMonkeyParse, err.Error())
		}
		err = m.Items.Enqueue(Item(item))
		if err != nil {
			return nil, fmt.Errorf("item enqueue: %w: %v", ErrInvalidMonkeyParse, err.Error())
		}
	}

	matches = operationRegex.FindAllStringSubmatch(lines[2], -1)
	if matches == nil || len(matches) != 1 || len(matches[0]) != 6 {
		return nil, ErrInvalidMonkeyParse
	}
	if matches[0][3] == "*" {
		mul, err := strconv.ParseInt(matches[0][4], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("operator: %w: %v", ErrInvalidMonkeyParse, err.Error())
		}
		m.Operation = OperationMul{
			multiplication: int(mul),
		}
	} else if matches[0][1] == "+" {
		add, err := strconv.ParseInt(matches[0][2], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("operator: %w: %v", ErrInvalidMonkeyParse, err.Error())
		}
		m.Operation = OperationAdd{
			addition: int(add),
		}
	} else if matches[0][0] == "old" && matches[0][5] == "old" { // I have no idea how regex capture groups work
		m.Operation = OperationMulSelf{}
	}

	var divisibleMonkeyTest DivisibleMonkeyTest
	_, err = fmt.Sscanf(lines[3], "  Test: divisible by %d", &divisibleMonkeyTest.divisibleBy)
	if err != nil {
		return nil, fmt.Errorf("divisible by: %w: %v", ErrInvalidMonkeyParse, err.Error())
	}

	_, err = fmt.Sscanf(lines[4], "    If true: throw to monkey %d", &divisibleMonkeyTest.trueTargetMonkeyNum)
	if err != nil {
		return nil, fmt.Errorf("true throw: %w: %v", ErrInvalidMonkeyParse, err.Error())
	}

	_, err = fmt.Sscanf(lines[5], "    If false: throw to monkey %d", &divisibleMonkeyTest.falseTargetMonkeyNum)
	if err != nil {
		return nil, fmt.Errorf("false throw: %w: %v", ErrInvalidMonkeyParse, err.Error())
	}
	m.Test = divisibleMonkeyTest

	return &m, nil // TODO
}
