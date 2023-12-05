package main

import (
	"fmt"
	"os"
	"strconv"

	cp "github.com/otiai10/copy"
)

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Println("Not enough arguments. Try \"yyyy dd\"")
		return
	}

	year := args[0]
	day := args[1]

	val, err := strconv.Atoi(day)
	if err != nil {
		fmt.Println("Wrong format for the day argument. Try an integer.")
		return
	}

	if val < 10 {
		day = "0" + day
	}

	folderName := year + "/" + year + "-" + day

	cp.Copy("day-example", folderName)

	f, _ := os.Create(folderName + "/test.txt")
	f, _ = os.Create(folderName + "/input.txt")

	f.Close()

	// create files for input.txt and test.txt
}
