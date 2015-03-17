package classic

import (
	"math/rand"
	"testing"
)

func TestRollDice(t *testing.T) {
	rand.Seed(1)
	expected := []int{10, 12, 3, 5, 6}
	for i, expect := range expected {
		if got := dice2d6.Roll(); got != expect {
			t.Errorf("Expected %d but got %d on dice roll #%d.", expect, got, i+1)
		}
	}
}

func TestRollDie(t *testing.T) {
	rand.Seed(1)
	expected := []int{6, 4, 6, 6, 2, 1}
	for i, expect := range expected {
		if got := die1d6.Roll(); got != expect {
			t.Errorf("Expected %d but got %d on dice roll #%d.", expect, got, i+1)
		}
	}
}
