package main

import (
	"bytes"
	"fmt"
	"regexp"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "gopher"
	password = "gopher"
	dbname   = "phone"
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

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Phone Normalizer started!")
	phones := []string{
		"1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"(123)456-7892",
	}
	resetDB(dbname)
	must(initDB(phones))
}
