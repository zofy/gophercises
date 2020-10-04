package bookAPI

import (
	"encoding/json"
	"os"
)

type Story map[string]Chapter

type Chapter struct {
	Title      string
	Paragraphs []string
	Options    []Option
}

type Option struct {
	Text    string
	Chapter string
}

func LoadBook(path string) (Story, error) {
	var story Story
	jsonFile, err := os.Open(path)
	if err != nil {
		return story, err
	}
	if json.NewDecoder(jsonFile).Decode(&story) != nil {
		return story, err
	}
	return story, nil
}
