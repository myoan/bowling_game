package bowling

import (
	"testing" // テストで使える関数・構造体が用意されているパッケージをimport
)

func TestFrameIsSpare(t *testing.T) {
	testcases := []struct {
		input  []int
		expect bool
	}{
		{
			[]int{0},
			false,
		},
		{
			[]int{1},
			false,
		},
		{
			[]int{1, 9},
			true,
		},
		{
			[]int{10, 0},
			false,
		},
	}
	for _, tc := range testcases {
		frm := NewFrame(false)
		for _, r := range tc.input {
			frm.Roll(r)
		}
		if frm.IsSpare() != tc.expect {
			t.Errorf("failed test\nexpect: %t, actual: %t\n", tc.expect, frm.IsSpare())
		}
	}
}

func TestFrameIsStrike(t *testing.T) {
	testcases := []struct {
		input  []int
		expect bool
	}{
		{
			[]int{},
			false,
		},
		{
			[]int{0},
			false,
		},
		{
			[]int{10},
			true,
		},
	}
	for _, tc := range testcases {
		frm := NewFrame(false)
		for _, r := range tc.input {
			frm.Roll(r)
		}
		if frm.IsStrike() != tc.expect {
			t.Errorf("failed test\nexpect: %t, actual: %t\n", tc.expect, frm.IsStrike())
		}
	}
}

func TestFrameIsFinish(t *testing.T) {
	testcases := []struct {
		isLast bool
		input  []int
		expect bool
	}{
		{
			false,
			[]int{0},
			false,
		},
		{
			false,
			[]int{1, 1},
			true,
		},
		{
			false,
			[]int{10},
			true,
		},
	}
	for _, tc := range testcases {
		frm := NewFrame(tc.isLast)
		for _, r := range tc.input {
			frm.Roll(r)
		}
		if frm.IsFinish() != tc.expect {
			t.Errorf("failed test\nexpect: %t, actual: %t\n", tc.expect, frm.IsFinish())
		}
	}
}

func TestBowlingPlay(t *testing.T) {
	testcases := []struct {
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
			[]int{1, 1, 1},
			3,
		},
		{
			[]int{1, 9, 2},
			14, // spare: (1 + 9) + 2 + 2
		},
		{
			[]int{10, 2},
			14, // strike: (10 + 0) + 2 + 2
		},
		{
			[]int{10, 2, 3, 4},
			24, // strike: (10 + 2 + 3) + 2 + 3 + 4
		},
		{
			[]int{10, 10, 3, 4},
			47, // strike: (10 + 10 + 3) + (10 + 3 + 4) + 3 + 4
		},
	}
	for _, tc := range testcases {
		game := NewGame()
		for _, r := range tc.input {
			game.Roll(r)
		}
		if game.Score() != tc.expect {
			t.Errorf("failed test\nexpect: %d, actual: %d\n", tc.expect, game.Score())
		}
	}
}
