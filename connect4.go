package main

import (
	"fmt"
	
	"github.com/thethirdone/connect4-solver/data"
	"github.com/thethirdone/connect4-solver/game"
)

func main() {
	y := game.Init(2)
	y.Player = 1
	
	y = y.Drop(3)
	y = y.Drop(4)
	y = y.Drop(3)
	y = y.Drop(4)
	y = y.Drop(3)
	y = y.Drop(4)
 
	y = y.Drop(0)
	y = y.Drop(0)
	y = y.Drop(0)
	y = y.Drop(0)
	y = y.Drop(0)
	y = y.Drop(0)
	
	fmt.Println(y.Evaluate())
	g, s, d := data.GetVals()
	fmt.Println(g, s, d)
}