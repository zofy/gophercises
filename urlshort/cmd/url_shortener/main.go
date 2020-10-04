package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/zofy/urlshort/internal/dbApi"
)

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func readFile(filePath string) ([]byte, error) {
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func createHandler(filePath string, db bool) (http.Handler, error) {
	fallback := defaultMux()
	if db == true {
		err := dbApi.InitDB()
		if err != nil {
			return nil, err
		}
		return DBHandler(fallback)
	}
	content, err := readFile(filePath)
	if err != nil {
		return nil, err
	}
	switch format := strings.Split(filePath, ".")[1]; format {
	case "yml":
		return YAMLHandler(content, MapHandler(nil, fallback))
	case "json":
		return JSONHandler(content, MapHandler(nil, fallback))
	default:
		return MapHandler(nil, fallback), nil
	}
}

func main() {
	filePath := flag.String("file", "source/paths.yml", "specify file path for pathURLs mapping")
	db := flag.Bool("db", false, "use url mapping from database")
	flag.Parse()

	httpHandler, err := createHandler(*filePath, *db)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", httpHandler)
}
