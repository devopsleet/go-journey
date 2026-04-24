package main

import (
	"log"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

}
