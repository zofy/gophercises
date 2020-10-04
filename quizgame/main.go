package main

import (
	"bufio"
	"flag"
	f "fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type problem struct {
	question, answer string
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return problems
}

func Start(path string, limit int, shuffle bool) int {
	var score int

	lines := readLines(path)

	if shuffle == true {
		rand.Shuffle(len(lines), func(i, j int) { lines[i], lines[j] = lines[j], lines[i] })
	}

	r := bufio.NewReader(os.Stdin)
	timer := time.NewTimer(time.Duration(limit) * time.Second)
	answerChan := make(chan string)

	for i, p := range parseLines(lines) {
		f.Printf("Problem #%d: %s = \n", i+1, p.question)
		go func() {
			answerChan <- readInput(r)
		}()
		select {
		case <-timer.C:
			f.Println("You ran out of time!")
			return score
		case answer := <-answerChan:
			f.Println(answer)
			if answer == p.answer {
				score++
			} else {
				f.Println("Wrong!")
				return score
			}
		}
	}
	return score
}

func main() {
	f.Println("Starting quiz game")
	csvFilename := flag.String("csv", "problems.csv", "a csv file containing problems in format of question,answer")
	timeLimit := flag.Int("limit", 30, "time limit for a quiz in seconds")
	shuffle := flag.Bool("shuffle", false, "shuffle order of quiz problems")
	flag.Parse()

	score := Start(*csvFilename, *timeLimit, *shuffle)
	f.Printf("Your score: %d\n", score)
}
