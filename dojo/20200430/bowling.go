package bowling

import "fmt"

type Frame struct {
	isLast bool
	pins   []int
}

func NewFrame(isLast bool) *Frame {
	if isLast {
		return &Frame{isLast: isLast, pins: []int{}}
	}
	return &Frame{isLast: isLast, pins: []int{}}
}

func (f *Frame) Roll(r int) {
	fmt.Printf("roll: %d\n", r)
	f.pins = append(f.pins, r)
}

func (f *Frame) IsSpare() bool {
	if !f.IsFinish() {
		return false
	}
	if f.IsStrike() {
		return false
	}
	if f.pins[0]+f.pins[1] != 10 {
		return false
	}
	return true
}

func (f *Frame) IsStrike() bool {
	if len(f.pins) == 0 {
		return false
	}
	if f.pins[0] == 10 {
		return true
	}
	return false
}

func (f *Frame) IsFinish() bool {
	if f.IsStrike() {
		return true
	}
	if !f.isLast && len(f.pins) == 2 {
		return true
	}
	if f.isLast && len(f.pins) == 3 {
		return true
	}
	return false
}

type Game struct {
	frame       []*Frame
	current     int
	strikeCount int
	rollCount   int
	rolls       []int
	score       int
}

func NewGame() *Game {
	fmt.Println("--- newGame")
	frm := make([]*Frame, 0)
	frm = append(frm, NewFrame(false))
	rolls := []int{}
	return &Game{frame: frm, current: 0, strikeCount: 0, rollCount: 0, rolls: rolls, score: 0}
}

func (g *Game) Roll(r int) {
	if g.rollCount+1 > len(g.rolls) {
		g.rolls = append(g.rolls, 1)
	}
	g.rolls[g.rollCount] *= r
	frame := g.frame[g.current]
	frame.Roll(r)
	if frame.IsFinish() {
		if frame.IsSpare() {
			g.rolls = append(g.rolls, 1)
			g.rolls[g.rollCount+1]++
		} else if frame.IsStrike() {
			g.rolls = append(g.rolls, 1)
			g.rolls = append(g.rolls, 1)
			g.rolls[g.rollCount+1]++
			g.rolls[g.rollCount+2]++
		}
		g.current++
		g.frame = append(g.frame, NewFrame(false))
	}
	g.rollCount++
}

func (g *Game) Score() int {
	score := 0
	for i, s := range g.rolls {
		fmt.Printf("[%d] %d\n", i, s)
		if i >= g.rollCount {
			break
		}
		score += s
	}
	return score
}
