package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/SamJohn04/dailynotes/internal/daily"
	"github.com/SamJohn04/dailynotes/internal/utils"
)

func main() {
	filename, err := utils.ParseArgs()
	if err != nil {
		printError(err)
		os.Exit(1)
	}

	err = utils.WriteNewFile(filename, daily.Boilerplate())
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func printError(err error) {
	if errors.Is(err, utils.ErrUsage) {
		fmt.Println("Usage: daily new [FILENAME]")
	} else if errors.Is(err, utils.ErrNotImplemented) {
		fmt.Println("Only new is currently implemented.")
	} else {
		fmt.Println(err)
	}
}
