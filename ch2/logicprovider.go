package main

import "fmt"

type LogicProvider struct{}

func (lp *LogicProvider) process(data *string) string {

	return "I'm the business logic " + *data
}

type Logic interface {
	process(data *string) string
}

type Client struct {
	L Logic
}

func (c Client) Program() string {

	data := "Hello Go"
	result := c.L.process(&data)

	return result

}

func main() {
	c := Client{
		L: &LogicProvider{},
	}
	finalresult := c.Program()

	fmt.Println("finalresult is " + finalresult)
	fmt.Println("Final result retuned by interface is " + finalresult)

}
