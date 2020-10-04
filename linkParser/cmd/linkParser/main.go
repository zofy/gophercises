package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zofy/linkParser/cmd/linkParser/link"
)

func exit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func main() {
	path := flag.String("--path", "source/test.html", "Path of HTML file to be parsed")

	flag.Parse()
	fmt.Printf("Running parser on: %s\n", *path)
	file, err := os.Open(*path)
	if err != nil {
		exit(err)
	}
	links, err := link.Parse(file)
	if err != nil {
		exit(err)
	}
	fmt.Printf("%+v\n", links)
}
