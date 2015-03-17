package main

import (
	"fmt"
	"math/rand"
	"monopoly/classic"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	board := classic.Board
	fmt.Printf("Rolled: %d\n", board.Dice().Roll())
}
