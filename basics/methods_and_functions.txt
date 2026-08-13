// package main

// import (
// 	"fmt"
// 	"time"
// )

// type Counter struct {
// 	counter    int
// 	lastUpdate time.Time
// }

// func (c Counter) String() string {
// 	return fmt.Sprintf("The value of counter is %d at time %s", c.counter, c.lastUpdate)
// }

// func (c *Counter) Increment() {
// 	c.counter++
// 	c.lastUpdate = time.Now()
// }

// func main() {

// 	c := &Counter{}
// 	output := c.String()
// 	fmt.Println(output)
// 	c.Increment()
// 	fmt.Println(c.String())

// }
