package bowling

import "testing"

func TestGame(t *testing.T) {
	testcase := []struct {
		input  []int
		expect int
	}{
		{
			[]int{0},
			0,
		},
		{
			[]int{1},
			1,
		},
		{
			[]int{1, 3},
			4,
		},
		{
			[]int{1, 9, 1},
			12, // (1 + 9 + 1) + 1
		},
		{
			[]int{1, 9, 1, 2},
			14, // (1 + 9 + 1) + 1 + 2
		},
		{
			[]int{10, 1, 2},
			16, // (10 + 1 + 2) + 1 + 2
		},
		{
			[]int{10, 10, 2, 3},
			42, // (10 + 10 + 2) + (10 + 2 + 3) + 2 + 3
		},
		{ // It's not correct (last game is no bonus)
			[]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
			270,
		},
		// { // perfect game
		// 	[]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		// 	300,
		// },
	}

	for _, tc := range testcase {
		game := NewGame()
		for _, r := range tc.input {
			game.Roll(r)
		}
		if tc.expect != game.Score() {
			t.Errorf("test failed! want: %d, actual: %d\n", tc.expect, game.Score())
		}
	}
}
