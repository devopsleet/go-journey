package main

import (
	"fmt"
	goo "net/http"
)

func main() {
	fmt.Println("Hello, Go standard library")

	resp, err := goo.Get("https://google.com")
	if err != nil {
		fmt.Println("Error ", err)
		return
	}kkkk


	

	defer resp.Body.Close()
	fmt.Println("HTTP reponse stauts", resp.Status)
}
