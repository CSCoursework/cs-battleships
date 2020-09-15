package game

import (
	"github.com/codemicro/cs-battleships/internal/models"
	"math/rand"
	"time"
)

var (
	Random *rand.Rand
	Ocean  [][]models.OceanCell
)

func init() {
	Random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func CreateOcean(oceanWidth, oceanHeight int) [][]models.OceanCell {
	var proto [][]models.OceanCell
	for y := 0; y < oceanHeight; y++ {
		var currentLine []models.OceanCell
		for x := 0; x < oceanWidth; x++ {
			num := Random.Intn(100)
			currentCell := models.OceanCell{}
			if num < 20 {
				currentCell.Occupied = true
			}
			currentLine = append(currentLine, currentCell)
		}
		proto = append(proto, currentLine)
	}
	return proto
}
