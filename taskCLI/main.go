package main

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/zofy/task/cmd"
	"github.com/zofy/task/db"
)

const projectPath = "/go/src/github.com/zofy/task/source/tasks.db"

func must(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	dir, _ := homedir.Dir()
	dbPath := dir + projectPath
	must(db.Init(dbPath))
	must(cmd.Execute())
}
