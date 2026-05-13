package main

import "fmt"

// Interface
type paymentMethod interface {
	Pay(amount float64)
}

//-------------------
// Credit card
//-------------------

type CreditCard struct {
	name string
}

func (c CreditCard) Pay(amount float64) {
	fmt.Printf(
		"Paid %.2f using credit card belonging to %s\n",
		amount,
		c.name,
	)
}

//-------------------
// Paypal
//-------------------

type Paypal struct {
	email string
}

func (p Paypal) Pay(amount float64) {
	fmt.Printf(
		"Paid %.2f using Pay pal account belonging to %s\n",
		amount,
		p.email,
	)
}

//----------------
// Bank transfer
//----------------

type BankTransfer struct {
	bankName string
}

func (b BankTransfer) Pay(amount float64) {
	fmt.Printf(
		"paid %.2f from bankname %s\n",
		amount,
		b.bankName,
	)
}

//-----------------
// Function using interface
//-----------------

func Checkout(p paymentMethod, amount float64) {
	fmt.Println("Processing payments......")
	p.Pay(amount)
	fmt.Println("Payment Complete")
	fmt.Println("")
}

func main() {

	card := CreditCard{
		"Gagan",
	}

	paypal := Paypal{
		"gasingla07@gmail.com",
	}

	bank := BankTransfer{
		"ABC",
	}

	Checkout(card, 100)
	Checkout(paypal, 200)
	Checkout(bank, 300)

}

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
