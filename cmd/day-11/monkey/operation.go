package monkey

type Operation interface {
	Do(i Item) Item
}

type OperationAdd struct {
	addition int
}

func (o OperationAdd) Do(i Item) Item {
	return i + Item(o.addition)
}

type OperationMul struct {
	multiplication int
}

func (o OperationMul) Do(i Item) Item {
	return i * Item(o.multiplication)
}

type OperationMulSelf struct{}

func (o OperationMulSelf) Do(i Item) Item {
	return i * i
}
