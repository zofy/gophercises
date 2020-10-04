package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func readInput(r *bufio.Reader) string {
	text, _ := r.ReadString('\n')
	return strings.Replace(text, "\n", "", -1)
}

func readLines(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		exit(fmt.Sprintf("Could not open file: %s", path))
	}
	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Error occurred while reading file: %s", path))
	}
	return lines
}
