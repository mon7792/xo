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
	GetPlayerInput(attempt int) (int, error)
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
func (g *xo) GetPlayerInput(attempt int) (int, error) {
	var (
		position int
		err      error
	)
inputLoop:
	for i := 1; i <= attempt; i++ {
		position, err = g.getInput()
		switch {
		case err != nil && i < attempt:
			fmt.Printf("Error:%v \n", err)
		case err != nil && i == attempt:
			fmt.Println("invalid input")
			break
		default:
			break inputLoop
		}
	}
	return position, err
}

func (g *xo) getInput() (int, error) {
	var position int
	fmt.Println("Enter the position")
	fmt.Scanln(&position)

	// validate the position.
	if position < 1 || position > 9 {
		return -1, errInvalidPosition
	}
	// validate the position is new
	if elementInSlice(g.players, g.grid[position-1]) {
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
	for g.turn < 10 {
		fmt.Println("Welcome to the game")
		// 1. put the introductory text here
		// 2. display the grid.
		g.DisplayGrid()
		// 4. accept the input
		position, err := g.GetPlayerInput(3)
		if err != nil {
			fmt.Println(err)
			break
		}
		// 4. set the current Player %2 operation.
		playerName := g.players[g.turn%2]
		// 5. SetPlayerInput
		g.SetPlayerInput(playerName, position)
		// 6. display the grid
		g.DisplayGrid()
		// 7. evalute the turn; if not 3(continue)
		// 8. start calculation of the result()

		g.turn++
	}
}

// helper function
func elementInSlice(str []string, elem string) bool {
	var exist = false
	for i := range str {
		if str[i] == elem {
			exist = true
			break
		}
	}
	return exist
}
