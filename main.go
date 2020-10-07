package main

import (
	"example.com/game"
)

func main() {
	// var position int
	// var err error
	// var attempt = 3
	gme := game.NewGame()
	gme.StartGame()

	// for i := 1; i <= attempt; i++ {
	// 	position, err = gme.GetPlayerInput()
	// 	switch {
	// 	case err != nil && i < attempt:
	// 		fmt.Printf("Error:%v \n", err)
	// 	case err != nil && i == attempt:
	// 		fmt.Println("invalid input")
	// 		os.Exit(0)
	// 	default:
	// 		break
	// 	}
	// }
	// gme.DisplayGrid()
	// gme.SetPlayerInput("X", position)
	// gme.DisplayGrid()

}
