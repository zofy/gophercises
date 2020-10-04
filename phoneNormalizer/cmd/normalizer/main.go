package main

import (
	"bytes"
	"fmt"
	"regexp"

	_ "github.com/lib/pq"
)

func normalizeSimple(phone string) string {
	var buf bytes.Buffer
	for _, ch := range phone {
		if ch >= '0' && ch <= '9' {
			buf.WriteRune(ch)
		}
	}
	return buf.String()
}

func normalizeRegexp(phoneNumber string) string {
	return regexp.MustCompile("\\D").ReplaceAllString(phoneNumber, "")
}

func main() {
	fmt.Println("Phone Normalizer started!")
}
