package io

import (
	"bufio"
	"fmt"
	"github.com/codemicro/cs-battleships/internal/helpers"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	scanner   = bufio.NewScanner(os.Stdin)
	cellRegex *regexp.Regexp
)

func init() {
	var err error
	cellRegex, err = regexp.Compile(`^\w\d$`)
	if err != nil {
		panic(err)
	}
}

func TakeInput(prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return scanner.Text()
}

func GetCell() (x int, y int) {
	for {
		input := TakeInput("Select a cell: ")
		if cellRegex.Match([]byte(input)) {

			x = helpers.GetCharNumber(strings.ToUpper(string(input[0])))
			y, _ = strconv.Atoi(string(input[1])) // Hopefully the regex means we can ignore this error

			if !(x > OceanWidth || y > OceanHeight) {
				return
			}
		}
		fmt.Println("Invalid cell")
	}
}
