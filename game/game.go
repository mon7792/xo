package game

import (
	"errors"
	"fmt"
)

var errInvalidPosition = errors.New("error invalid position entered")
var grid = `
 ____ ____ ____ 
| %s  | %s  | %s  |
|____|____|____|
| %s  | %s  | %s  |
|____|____|____|
| %s  | %s  | %s  |
|____|____|____|

`

type xo struct {
	grid    []string
	players []string
	turn    int
}

// NewGame start the new game
func NewGame() XOInterface {
	return &xo{
		grid:    []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"},
		players: []string{"X", "O"},
		turn:    1,
	}
}

// XOInterface exposes the game functionality
type XOInterface interface {
	DisplayGrid()
	GetPlayerInput() (int, error)
	SetPlayerInput(playerName string, position int)
	StartGame()
}

// DisplayGrid display the current game grid
func (g *xo) DisplayGrid() {
	fmt.Printf(grid, g.grid[0], g.grid[1], g.grid[2],
		g.grid[3], g.grid[4], g.grid[5],
		g.grid[6], g.grid[7], g.grid[8])
}

// GetPlayerInput gets the input from the player and validate it.
func (g *xo) GetPlayerInput() (int, error) {
	var position int
	fmt.Println("Enter the position")
	fmt.Scanln(&position)
	// validate the input.
	if position < 1 || position > 9 {
		return -1, errInvalidPosition
	}
	return position, nil
}

// SetPlayerInput sets the playername at postion - 1 in the grid
func (g *xo) SetPlayerInput(playerName string, position int) {
	position = position - 1
	g.grid[position] = playerName
}

// StartGame() starts one game session of xo
func (g *xo) StartGame() {
	for g.turn < 6 {
		fmt.Println("Welcome to the game")
		// 1. put the introductory text here
		// 2. display the grid.
		// 3. set the current Player %2 operation.
		// 4. accept the input
		// 5. display the grid
		// 6. evalute the turn; if not 3(continue)
		// 7. start calculation of the result()

		g.turn++
	}
}
