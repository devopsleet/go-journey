package main

import "strings"

// Create an interface

type Orderable interface {
	Order(any) int
}

type OrderableInt int

func (oi OrderableInt) Order(val any) int {
	return int(oi - val.(OrderableInt))
}

type OrderableString string

func (os OrderableString) Order(val any) int {
	return strings.Compare(string(os), val.(string))
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

func main() {

	var t *Tree
	t = t.Insert(OrderableInt(5))
	t = t.Insert(OrderableInt(7))
	t = t.Insert(OrderableString("nope"))

}
