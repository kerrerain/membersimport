package main

import (
	"log"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatal("The file to import has not been provided. Please run './membersimport <filename>.csv'")
	}

	if err := ProcessFile(args[0], exportFileName()); err != nil {
		log.Fatal(err.Error())
	}
}

func exportFileName() string {
	return EXPORT_FOLDER + "members_exported" + timestamp() + ".csv"
}

func timestamp() string {
	return time.Now().Format("20060102150405")
}
