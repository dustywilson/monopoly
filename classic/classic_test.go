package classic

import (
	"math/rand"
	"monopoly/common"
	"reflect"
	"testing"
)

func TestBoardDice(t *testing.T) {
	rand.Seed(1)
	b := Board
	expected := []int{10, 12, 3, 5, 6}
	for i, expect := range expected {
		if got := b.Dice().Roll(); got != expect {
			t.Errorf("Expected %d but got %d on dice roll #%d.", expect, got, i+1)
		}
	}
}

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

func TestBoardStartSpace(t *testing.T) {
	b := Board
	expect := 200
	start := b.StartSpace()
	if start == nil {
		t.Error("Got a nil StartSpace.")
		return
	}
	if rs, ok := start.(common.RewardSpace); ok {
		got := rs.Reward()
		if got != 200 {
			t.Errorf("Expected RewardSpace for $%d but got $%d instead.", expect, got)
		}
	} else {
		t.Errorf("Expected a RewardSpace but %s isn't a RewardSpace.", reflect.TypeOf(start).String())
	}
}
