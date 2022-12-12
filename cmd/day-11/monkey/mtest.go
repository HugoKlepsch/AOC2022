package monkey

type MonkeyTest interface {
	Route(i Item) int
	Divisor() int
}

type DivisibleMonkeyTest struct {
	divisibleBy          int
	trueTargetMonkeyNum  int
	falseTargetMonkeyNum int
}

func (d DivisibleMonkeyTest) Route(i Item) int {
	if i%Item(d.divisibleBy) == 0 {
		return d.trueTargetMonkeyNum
	} else {
		return d.falseTargetMonkeyNum
	}
}
func (d DivisibleMonkeyTest) Divisor() int {
	return d.divisibleBy
}
