package main

import (
	"github.com/codemicro/cs-battleships/internal/display"
	"github.com/codemicro/cs-battleships/internal/game"
)

const (
	oceanWidth  = 10
	oceanHeight = 10
)

func init() {
	game.Ocean = game.CreateOcean(oceanWidth, oceanHeight)
	display.DisplayOcean(game.Ocean)
}

func main() {

}
