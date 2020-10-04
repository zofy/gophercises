package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zofy/adventure/internal/bookAPI"
	"github.com/zofy/adventure/internal/server"
)

func exit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func main() {
	port := flag.Int("port", 8000, "Server port")
	filename := flag.String("file", "gopher.json", "Actual story")
	flag.Parse()

	story, err := bookAPI.LoadBook(*filename)
	if err != nil {
		exit(err)
	}
	fmt.Printf("Starting an adventure! %s\n", *filename)
	server.New(*port).Start(story)
}
