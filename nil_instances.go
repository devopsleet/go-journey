package main

type IntTree struct {
	val         int
	left, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}

	if val < it.val {
		it.left = it.left.Insert(val)

	} else if val > it.val {
		it.right = it.right.Insert(val)
	}

	return it

}

func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return
	}
}

func main() {

	var it *IntTree
	it = it.Insert(5)

}
