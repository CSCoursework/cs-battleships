package main

import (
	"github.com/codemicro/cs-battleships/internal/game"
	"github.com/codemicro/cs-battleships/internal/io"
)

func init() {
	game.Ocean = game.CreateOcean(io.OceanWidth, io.OceanHeight)
}

func main() {
	game.Start()
}
