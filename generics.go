package main

import "fmt"

// interface
type Orderable interface {
	Order(any) int
}

type Tree struct {
	val         Orderable
	left, right *Tree
}

func (t *Tree) Insert(val Orderable) *Tree {
	if t == nil {
		return &Tree{val: val}
	}

	switch comp := val.Order(t.val); {

	case comp < 0:
		t.left = t.left.Insert(val)
	case comp > 0:
		t.right = t.right.Insert(val)

	}

	return t
}

type OrderableInt int

func (oi OrderableInt) Order(val any) int {
	return int(oi - val.(OrderableInt))
}

func main() {

	var it *Tree

	it = it.Insert(OrderableInt(5))
	it = it.Insert(OrderableInt(3))

	fmt.Println(it.left.val)
	//fmt.Println(it.right.val)

}
