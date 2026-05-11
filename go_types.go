package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total      int
	lastUpdate time.Time
}

func (c Counter) String() string {
	return fmt.Sprintf("The values are %d and %v", c.total, c.lastUpdate)
}

func (c *Counter) Increment() {

	c.total++
	c.lastUpdate = time.Now()

}

func main() {

	// local variable value type
	var c Counter

	output := c.String()
	fmt.Println(output)

	c.Increment()

	fmt.Println(c.String())

}
