package main

import (
	"fmt"
	"github.com/peterh/liner"
	"github.com/suzuken/minidb"
	"io"
	"os"
)

var (
	HistoryFile = "/tmp/.liner_history"
)

func main() {
	db := minidb.New()

	// liner state
	ls := liner.NewLiner()
	defer ls.Close()

	if f, err := os.Open(HistoryFile); err == nil {
		ls.ReadHistory(f)
		f.Close()
	}

	// repl
	for {
		in, err := ls.Prompt("minidb> ")
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Fprintf(os.Stderr, "fatal: %s", err)
			os.Exit(1)
		}

		if in == "" {
			continue
		}

		err = db.Eval(in)
		if err != nil {
			fmt.Println(err)
		}
	}
}
