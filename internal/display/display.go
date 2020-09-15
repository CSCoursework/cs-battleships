package display

import (
	"fmt"
	"github.com/codemicro/cs-battleships/internal/helpers"
)

func DisplayOcean(ocean [][]oceanCell) {
	helpers.ClearConsole()

	fmt.Print("  ")

	for i := 0; i < oceanWidth; i++ {
		fmt.Printf(" %s ", helpers.GetAlphabetChar(i))
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