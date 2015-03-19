package classic

import "monopoly/common"

// board is the game board
type board struct {
	dice       Dice
	startSpace common.Space
}

// Board is the classic Monopoly Board
var Board = &board{
	dice:       dice2d6,
	startSpace: GoSpace{},
}

// Dice are the dice used for this board, a pair of typical 6-sided dice.
func (b *board) Dice() Dice {
	return b.dice
}

// StartSpace is the spot that starts the game (not necessarily the Go spot on all boards)
func (b *board) StartSpace() common.Space {
	return b.startSpace
}
