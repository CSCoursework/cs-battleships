package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

var (
	ocean  [][]oceanCell
	random *rand.Rand
)

const (
	oceanWidth  = 10
	oceanHeight = 10
)

type oceanCell struct {
	Hit      bool
	Occupied bool
}

func clearConsole() {
	out, _ := exec.Command("cls").Output() // Windows only
	os.Stdout.Write(out)
}

func initialiseOcean() {
	var proto [][]oceanCell
	for y := 0; y < oceanHeight; y++ {
		var currentLine []oceanCell
		for x := 0; x < oceanWidth; x++ {
			num := random.Intn(100)
			currentCell := oceanCell{}
			if num < 20 {
				currentCell.Occupied = true
			}
			currentLine = append(currentLine, currentCell)
		}
		proto = append(proto, currentLine)
	}
	ocean = proto
}

func getAlphabetChar(i int) string {
	return string(rune(int('A') + i))
}

func displayOcean() {
	clearConsole()

	fmt.Print("  ")

	for i := 0; i < oceanWidth; i++ {
		fmt.Printf(" %s ", getAlphabetChar(i))
	}

	fmt.Println()

	for y := 0; y < oceanHeight; y++ {
		fmt.Printf(" %d", y)
		for x := 0; x < oceanWidth; x++ {

			currentCell := ocean[x][y]
			var marker string
			if currentCell.Hit {
				marker = "*"
			} else {
				marker = "-"
			}
			fmt.Printf(" %s ", marker)

		}
		fmt.Println()
	}
}

func init() {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
	initialiseOcean()
	displayOcean()
}

func main() {
	
}
