package cmd

import (
	"test_github_actions/cmd"
	"testing"
)

func TestSum(t *testing.T) {
	var x, y int
	var want int
	x = 1
	y = 2
	want = 3

	result := cmd.Sum(x, y)

	if result == want {
		t.Errorf("test failed.")
	}
}

func TestMul(t *testing.T) {
	var x, y int
	var want int

	x = 10
	y = 20
	want = 200

	result := cmd.Mul(x, y)

	if result != want {
		t.Errorf("test failed.")
	}
}
