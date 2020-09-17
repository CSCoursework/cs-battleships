package io

import (
	"fmt"
	"github.com/codemicro/cs-battleships/internal/helpers"
	"github.com/codemicro/cs-battleships/internal/models"
)

const (
	OceanWidth  = 10
	OceanHeight = 10
)

func ShowOcean(ocean [][]models.OceanCell) {
	helpers.ClearConsole()

	fmt.Print("  ")

	for i := 0; i < len(ocean); i++ {
		fmt.Printf(" %s ", helpers.GetAlphabetChar(i))
	}

	fmt.Println()

	for y := 0; y < len(ocean); y++ {
		fmt.Printf(" %d", y)
		for x := 0; x < len(ocean[0]); x++ {

			currentCell := ocean[x][y]
			var marker string
			if currentCell.Occupied {
				marker = "*"
			} else {
				marker = "-"
			}
			fmt.Printf(" %s ", marker)

		}
		fmt.Println()
	}
}
