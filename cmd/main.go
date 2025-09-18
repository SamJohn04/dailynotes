package main

import (
	"fmt"

	"github.com/SamJohn04/dailynotes/internal/daily"
	"github.com/SamJohn04/dailynotes/internal/utils"
)

func main() {
	filename, err := utils.ParseArgs()
	if err != nil {
		printError(err)
		return
	}

	err = utils.WriteNewFile(filename, daily.Boilerplate())
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func printError(err error) {
	if err.Error() == "usage error" {
		fmt.Println("Usage: daily new [FILENAME]")
	} else if err.Error() == "unimplemented error" {
		fmt.Println("Only new is currently implemented.")
	} else {
		fmt.Println(err)
	}
}
