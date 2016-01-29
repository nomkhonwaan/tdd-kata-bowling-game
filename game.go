package bowling

type Game struct {
	i      int
	Frames []Frame
}

func NewGame() *Game {
	return &Game{
		Frames: make([]Frame, 10),
	}
}

func (g *Game) Roll(pins int) {
	f := &g.Frames[g.i]
	f.ID = g.i + 1
	f.Roll(pins)

	if f.IsFinished() {
		if g.i > 0 {
			f.SetPrevious(&g.Frames[g.i-1])
		}
		if g.i < len(g.Frames)-1 {
			f.SetNext(&g.Frames[g.i+1])
		}

		g.i++
	}
}

func (g *Game) Score() int {
	var totalScore int
	for _, frame := range g.Frames {
		totalScore += frame.Score()
	}

	return totalScore
}
