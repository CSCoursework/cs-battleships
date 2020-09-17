package game

import (
	"fmt"
	"github.com/codemicro/cs-battleships/internal/io"
	"github.com/codemicro/cs-battleships/internal/models"
	"math/rand"
	"time"
)

var (
	Random *rand.Rand
	Ocean  [][]models.OceanCell
	shipsToPlace = []int{5, 4, 3, 3, 2}
)

func init() {
	Random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func CreateOcean(oceanWidth, oceanHeight int) (proto [][]models.OceanCell) {
	for y := 0; y < oceanHeight; y++ {
		var currentLine []models.OceanCell
		for x := 0; x < oceanWidth; x++ {
			currentLine = append(currentLine, models.OceanCell{})
		}
		proto = append(proto, currentLine)
	}

	for _, shipLen := range shipsToPlace {
		// pick random orientation
		isShipHorizontal := Random.Intn(2) % 2 == 0

		var x int
		var y int
		// check if ship has enough space
		for {
			// pick a random coordinate
			x = Random.Intn(oceanWidth)
			y = Random.Intn(oceanHeight)

			if isShipHorizontal {
				// Prevent overflow
				for x + shipLen > oceanWidth {
					x--
				}

				// Check for collisions
				for i := 0; i < shipLen; i++ {
					if proto[x + i][y].Occupied {
						continue
					}
				}

			} else {
				// Prevent overflow in the other direction
				for y + shipLen > oceanHeight {
					y--
				}

				// Check for collisions
				for i := 0; i < shipLen; i++ {
					if proto[x][y + i].Occupied {
						continue
					}
				}
			}
			break
		}

		if isShipHorizontal {
			for i := 0; i < shipLen; i++ {
				thing := proto[x + i][y]
				thing.Occupied = true
				proto[x + i][y] = thing
			}
		} else {
			for i := 0; i < shipLen; i++ {
				thing := proto[x][y + i]
				thing.Occupied = true
				proto[x][y + i] = thing
			}
		}
	}

	/*for y := 0; y < oceanHeight; y++ {
		for x := 0; x < oceanWidth; x++ {
			num := Random.Intn(100)
			currentCell := proto[x][y]
			if num < 20 {
				currentCell.Occupied = true
			}
			proto[x][y] = currentCell
		}
	}*/

	return
}

func Start() {
	io.ShowOcean(Ocean)
	cell := io.GetCell()
	fmt.Println(cell)
}
