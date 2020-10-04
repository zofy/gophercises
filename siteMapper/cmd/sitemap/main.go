package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zofy/sitemap/internal/mapper"
)

func exit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func main() {
	domain := flag.String("domain", "https://gophercises.com", "Base URL to be site-mapped")
	maxDepth := flag.Int("max-depth", 3, "Max depth of the sitemap tree")

	flag.Parse()
	// fmt.Printf("Mapping site: %s\n", *domain)
	sm, err := mapper.BuildMap(*domain, *maxDepth)
	if err != nil {
		exit(err)
	}
	// sm.Print()
	sm.ToXML()
}
