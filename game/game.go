package game

import (
	"errors"
	"fmt"
)

const (
	draw      = "draw"
	win       = "win"
	lose      = "lose"
	inProcess = "in_process"
)

var (
	// errors
	errInvalidPosition = errors.New("error invalid position entered")
	errFilledPosition  = errors.New("error position is already filled")
	winCombo           = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
		{1, 4, 7}, {2, 5, 8}, {3, 6, 9},
		{1, 5, 9}, {3, 5, 7}}

	// grid
	grid = `
 ____ ____ ____ 
| %s  | %s  | %s  |
|____|____|____|
| %s  | %s  | %s  |
|____|____|____|
| %s  | %s  | %s  |
|____|____|____|

`
)

type xo struct {
	grid        []string
	players     []string
	turn        int
	playerMoves map[string][]int
}

// NewGame start the new game
func NewGame() XOInterface {
	return &xo{
		grid:        []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"},
		players:     []string{"X", "O"},
		turn:        1,
		playerMoves: make(map[string][]int),
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

// getInput accept the user input
func (g *xo) getInput() (int, error) {
	var position int
	fmt.Println("Enter the position")
	fmt.Scanln(&position)

	// validate the position.
	if position < 1 || position > 9 {
		return -1, errInvalidPosition
	}
	// validate the position is new
	if g.isPositionFilled(position) {
		return -1, errFilledPosition
	}

	return position, nil
}

// SetPlayerInput sets the playername at postion - 1 in the grid
func (g *xo) SetPlayerInput(playerName string, position int) {
	// set the player move
	g.playerMoves[playerName] = append(g.playerMoves[playerName], position)
	// set the grid position
	g.grid[position-1] = playerName
}

// StartGame() starts one game session of xo
func (g *xo) StartGame() {
	// 1. put the introductory text here
	fmt.Println("Welcome to the game")
	// 2. display the grid.
	g.DisplayGrid()
	for g.turn < 10 {
		// 3. set the current Player %2 operation.
		playerName := g.players[g.turn%2]
		fmt.Printf("%s Player turns \n", playerName)
		// 4. accept the input
		position, err := g.GetPlayerInput(3)
		if err != nil {
			fmt.Println(err)
			break
		}
		// 5. SetPlayerInput
		g.SetPlayerInput(playerName, position)
		// 6. display the grid.
		g.DisplayGrid()
		// 7. increment the turn
		g.turn++
		// 8. evalute the turn; if not 3(continue),start calculation of the result()
		g.EvaluateGameResult()
	}
}

// isPositionFilled checks if position is filled
func (g *xo) isPositionFilled(position int) bool {
	var posFilled = false
	for i := range g.players {
		if g.players[i] == g.grid[position-1] {
			posFilled = true
			break
		}
	}
	return posFilled
}

// EvaluateResult docs
func (g *xo) EvaluateGameResult() string {
	if g.turn < 6 {
		return inProcess
	}
	fmt.Println(g.playerMoves)
	// 1. calculate the number of turns. if turn is !3 return nothing.
	// 2. evaluate the current player Move.
	// g.playerMoves[]
	return ""
}

func evalutePlayerWin(moves []int) {}
