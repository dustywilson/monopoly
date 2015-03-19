package classic

import "math/rand"

// Dice is the pair of standard 6-sided dice.
type Dice struct {
	dice []Die
}

// dice2d6 is the standard pair of 6-sided dice.
var dice2d6 = Dice{dice: []Die{die1d6, die1d6}}

// Roll rolls the pair of dice and returns the sum
func (d Dice) Roll() int {
	v := 0
	for _, die := range d.dice {
		v = v + die.Roll()
	}
	return v
}

// Die is a single 6-sided die.
type Die struct {
	sides []int
}

// die1d6 is the standard 6-sided die.
var die1d6 = Die{sides: []int{1, 2, 3, 4, 5, 6}}

// Roll rolls the die and returns the result
func (d Die) Roll() int {
	return d.sides[rand.Int31n(int32(len(d.sides)))]
}
