# Battleships

Golang 1.14, `github.com/codemicro/cs-battleships`

#### To build

```bash
go get -u github.com/codemicro/cs-battleships
go build github.com/codemicro/cs-battleships/cmd/battleships
```

#### To run

```bash
sudo chmod +x ./battleships
./battleships
```

## Structure

The vast majority of the code in this project is stored in packages in the `/internal` directory, which contains the following packages:

* `game`: contains logic related to playing the game and managing the scoreboard.
* `helpers`: contains functions required in different places in the game that don't really fit into any of the other packages
  * For example, functions to clear the console and functions to convert letters into numbers are contained in this package.
* `io`: contains functions to aid with displaying content to and collecting input from the user.
* `models`: contains the structs used throughout the app
  * Specifically, it's just one model - `OceanCell`

## Creating and storing the ocean state

The heart of a battleships game is the ocean, which is in this case a 10x10 grid. Each cell in the grid has three different attributes attached to it: `Occupied`, `Hit` and `Guessed`.

This is achieved using a struct called `OceanCell`.

```go
type OceanCell struct {
	Hit      bool
	Occupied bool
	Guessed  bool
}
```

This alone isn't much help - to organise the ocean as a whole, the `Ocean` variable is used. It's of type `[][]models.OceanCell`, an array of arrays of our `OceanCell` struct, but each cell requires initialisation, since the array is currently made up of nil values. To do this, we have to manually create each ocean cell.

```go
func CreateOcean(oceanWidth, oceanHeight int) (proto [][]models.OceanCell) {
	for y := 0; y < oceanHeight; y++ {
		var currentLine []models.OceanCell
		for x := 0; x < oceanWidth; x++ {
			currentLine = append(currentLine, models.OceanCell{})
		}
		proto = append(proto, currentLine)
	}

	// ...

	return
}
```

Since Go doesn't have native support for two-dimensional arrays, we first create subarrays and then append those to the main ocean array.

However - this just creates an empty ocean. Battleships needs boats, right? We need to write some boat placement code.

```go
func CreateOcean(oceanWidth, oceanHeight int) (proto [][]models.OceanCell) {
    
    // ...
    
    // shipsToPlace is an array of integers that represent ships to be placed in the ocean.
	for _, shipLen := range shipsToPlace {
		// First, a random boat orientation is chosen.
		isShipHorizontal := Random.Intn(2) == 0

		var x int
		var y int
        
		// Enter an infinte loop
		for {
			// Pick a random coordinate on the board
			x = Random.Intn(oceanWidth)
			y = Random.Intn(oceanHeight)

            // Now we adjust those coordinates to ensure that the current ship won't overflow the board
            // In addition, we also check for collisions with other ships already on the board.
			if isShipHorizontal {
				// Prevent overflow by moving the ship left
				for x+shipLen > oceanWidth {
					x--
				}

				// Check for collisions. If there is a collision, return to the beginning of the loop
                // and generate new coordinates.
				for i := 0; i < shipLen; i++ {
					if proto[x+i][y].Occupied {
						continue
					}
				}

			} else {
				// Now we do the same thing, but for the y direction
                
				for y+shipLen > oceanHeight {
					y--
				}

				// Check for collisions
				for i := 0; i < shipLen; i++ {
					if proto[x][y+i].Occupied {
						continue
					}
				}
			}
            
            // If we've reached this point, we've adjusted the ship location so it's not overflowed
            // and we've ensured that there're no collisions.
			break
		}

		// Finally, we can set the cells of the ship to occupied in the ocean array.
		if isShipHorizontal {
			for i := 0; i < shipLen; i++ {
				thing := proto[x+i][y]
				thing.Occupied = true
				proto[x+i][y] = thing
			}
		} else {
			for i := 0; i < shipLen; i++ {
				thing := proto[x][y+i]
				thing.Occupied = true
				proto[x][y+i] = thing
			}
		}
	}
    
    // Because of Go magic, this is equivalent of "return proto"
    return 
}
```

## Displaying the ocean

Creating and storing the ocean is all well and good, but at the moment there's no way for the user to see it. For this, I think we need a new function.

```go
func ShowOcean(ocean [][]models.OceanCell) {
	helpers.ClearConsole() // We'll get to this later

	fmt.Print("  ")

    // Print the letters at the top of the board
	for i := 0; i < len(ocean); i++ {
		fmt.Printf(" %s ", helpers.GetAlphabetChar(i))
	}

	fmt.Println()

    // Iterate over the ocean (y first)
	for y := 0; y < len(ocean); y++ {
		fmt.Printf(" %d", y) // Print the row number
        
        // Iterate over the ocean in the x direction
		for x := 0; x < len(ocean[0]); x++ {
			
            // Get the current cell, and depending on if the cell is hit, or has been guessed before,
            // print out an "!" or a "x" respectively.
			currentCell := ocean[x][y]
			var marker string
			if currentCell.Hit {
				marker = "!"
			} else if currentCell.Guessed {
				marker = "x"
			} else {
				marker = "-"
			}
            // Print out the marker (with some padding in)
			fmt.Printf(" %s ", marker)

		}
		fmt.Println()
	}
}

```

## Clearing the console

To prevent the console becoming cluttered with outdated copies of the ocean, it's necessary to clear the console periodically.

Unlike in C# and other languages, Go has no built in way to do this. If we want to clear the console, we have to write a function that manually calls the command.

```go
func ClearConsole() {
	var cmd *exec.Cmd
    // Depending on the platform we're running on, we need to choose a different command.
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
    // The output of the command is set to the current command line that we're playing the game on.
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}
```

