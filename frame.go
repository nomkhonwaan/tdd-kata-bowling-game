package bowling

type Frame struct {
	i          []int
	totalScore int
	ID         int
	IsSpare    bool
	IsStrike   bool

	previous *Frame
	next     *Frame
}

func (f *Frame) Roll(pins int) {
	f.i = append(f.i, pins)
	f.totalScore += pins

	// spare
	if len(f.i) == 2 && f.totalScore == 10 {
		f.IsSpare = true
	}

	// strike
	if len(f.i) == 1 && f.totalScore == 10 {
		f.IsStrike = true
	}
}

func (f *Frame) Score() int {
	// spare
	if f.IsSpare && f.next != nil && f.ID != 10 {
		return f.totalScore + f.next.FirstScore()
	}

	// strike
	if f.IsStrike && f.next != nil && f.ID != 10 {
		if f.next.IsStrike && f.ID != 9 {
			return f.totalScore +
				f.next.FirstScore() +
				f.next.next.FirstScore()
		}
		return f.totalScore +
			f.next.FirstScore() +
			f.next.SecondScore()
	}

	return f.totalScore
}

func (f *Frame) FirstScore() int {
	return f.i[0]
}

func (f *Frame) SecondScore() int {
	if len(f.i) >= 2 {
		return f.i[1]
	}
	return 0
}

func (f *Frame) ThirdScore() int {
	if len(f.i) >= 3 {
		return f.i[2]
	}
	return 0
}

func (f *Frame) IsFinished() bool {
	if (f.IsSpare || f.IsStrike) &&
		len(f.i) == 3 &&
		f.ID == 10 {
		return true
	}

	if (f.IsStrike || len(f.i) == 2) &&
		f.ID != 10 {
		return true
	}

	return false
}

func (f *Frame) SetPrevious(prev *Frame) {
	f.previous = prev
}

func (f *Frame) SetNext(next *Frame) {
	f.next = next
}
