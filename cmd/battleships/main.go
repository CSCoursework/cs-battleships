package main

import (
	"github.com/codemicro/cs-battleships/internal/game"
)

func init() {
	game.SetupNewGame()
}

func main() {
	game.Start()
}
