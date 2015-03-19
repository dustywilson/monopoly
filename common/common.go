package common

// Board is an implementation of a Monopoly-style board
type Board interface {
	Dice() Dice
	StartSpace() Space
	Players() []Player
}

// Space is an implementation of a Monopoly board space
type Space interface {
	Name() string
}

// RewardSpace is a space that provides a reward for landing on it (Go)
type RewardSpace interface {
	Space
	Reward() int
	PassReward() int
}

// PenaltySpace is a space that charges a penalty for landing on it (taxes)
type PenaltySpace interface {
	Space
}

// CardSpace is a space that has an associated card deck (Chance / Community Chest)
type CardSpace interface {
	Space
}

// PropertySpace is a space that is an ownable property
type PropertySpace interface {
	Space
	Owned() bool
}

// GotoSpace is a space that causes you to move your token elsewhere (Go To Jail)
type GotoSpace interface {
	Space
}

// CaptiveSpace is a space that causes the player to be held captive until matching certain criteria (Jail)
type CaptiveSpace interface {
	Space
}

// Player is the player token and player attributes
type Player interface {
	Space() Space
}

// Dice is the set of dice used for rolling
type Dice interface {
	Roll() int
}

// Die is a member of the Dice for rolling
type Die interface {
	Roll() int
	Sides() []int
}
