package io

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var (
	scanner   = bufio.NewScanner(os.Stdin)
	cellRegex *regexp.Regexp
)

func init() {
	var err error
	cellRegex, err = regexp.Compile(`^\d+\w+$`)
	if err != nil {
		panic(err)
	}
}

func TakeInput(prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return scanner.Text()
}

func GetCell() string {
	for {
		input := TakeInput("Select a cell: ")
		fmt.Println(cellRegex.Match([]byte(input)))
		if cellRegex.Match([]byte(input)) {
			return input
		}
		fmt.Println("Invalid cell")
	}
}
