package classic

type GoSpace struct{}

func (g GoSpace) Name() string { return "Go" }

func (g GoSpace) Reward() int { return 200 }

func (g GoSpace) PassReward() int { return 200 }
