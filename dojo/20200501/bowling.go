package bowling

import "fmt"

type Game struct {
	score int
	bonus []int
	pins  []int
	idx   int
}

func NewGame() *Game {
	fmt.Println("--- create NewGame")
	pins := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	bonus := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	return &Game{score: 0, bonus: bonus, pins: pins, idx: 0}
}

func (g *Game) Roll(r int) {
	fmt.Println(g.bonus[g.idx])
	g.pins[g.idx] = r
	g.score += r * g.bonus[g.idx]
	if g.IsSpare() {
		g.bonus[g.idx+1]++
		fmt.Println(g.bonus[g.idx+1])
	}
	if g.IsStrike(r) {
		g.bonus[g.idx+1]++
		g.bonus[g.idx+2]++
		fmt.Println(g.bonus[g.idx+1])
	}
	g.idx++
}

func (g *Game) IsSpare() bool {
	fmt.Printf("is spare?: %d\n", g.idx)
	if (g.idx+1)%2 == 0 && g.pins[g.idx-1]+g.pins[g.idx] == 10 {
		fmt.Println("spare")
		return true
	}
	return false
}

func (g *Game) IsStrike(r int) bool {
	fmt.Printf("is strike?: %d\n", g.idx)
	if r == 10 {
		fmt.Println("strike")
		return true
	}
	return false
}

func (g *Game) Score() int {
	return g.score
}
