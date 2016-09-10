package main

import (
	"log"
)

func main() {
	if err := ProcessFile("test.csv", "test_exported.csv"); err != nil {
		log.Fatal(err.Error())
	}
}
