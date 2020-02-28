package main

import (
	"bytes"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		desc     string
		input    Sudoku
		expected Sudoku
	}{
		{
			desc: "solver",
			input: Sudoku{[][]byte{
				{0, 0, 8 /**/, 0, 0, 2 /**/, 0, 5, 0},
				{0, 4, 0 /**/, 0, 0, 5 /**/, 0, 0, 8},
				{0, 3, 5 /**/, 6, 0, 0 /**/, 2, 0, 7},
				/*											 						*/
				{3, 0, 1 /**/, 0, 2, 0 /**/, 0, 0, 0},
				{5, 0, 0 /**/, 0, 7, 0 /**/, 0, 0, 1},
				{0, 0, 0 /**/, 5, 9, 0 /**/, 8, 0, 3},
				/*											 						*/
				{7, 0, 3 /**/, 0, 0, 4 /**/, 0, 9, 0},
				{0, 0, 0 /**/, 2, 0, 0 /**/, 0, 8, 0},
				{0, 5, 0 /**/, 0, 0, 0 /**/, 6, 0, 0},
			}},
			expected: Sudoku{[][]byte{
				{9, 7, 8 /**/, 1, 4, 2 /**/, 3, 5, 6},
				{2, 4, 6 /**/, 7, 3, 5 /**/, 9, 1, 8},
				{1, 3, 5 /**/, 6, 8, 9 /**/, 2, 4, 7},
				/*											 						*/
				{3, 8, 1 /**/, 4, 2, 6 /**/, 5, 7, 9},
				{5, 9, 2 /**/, 3, 7, 8 /**/, 4, 6, 1},
				{4, 6, 7 /**/, 5, 9, 1 /**/, 8, 2, 3},
				/*											 						*/
				{7, 2, 3 /**/, 8, 6, 4 /**/, 1, 9, 5},
				{6, 1, 9 /**/, 2, 5, 3 /**/, 7, 8, 4},
				{8, 5, 4 /**/, 9, 1, 7 /**/, 6, 3, 2},
			}},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			assert.Equal(t, tC.expected, tC.input.Solve())
		})
	}
}

type pos struct {
	x int
	y int
}

func TestValidNumbersAt(t *testing.T) {

	sudoku := Sudoku{[][]byte{
		{0, 0, 8 /**/, 0, 0, 2 /**/, 0, 5, 0},
		{0, 4, 0 /**/, 0, 0, 5 /**/, 0, 0, 8},
		{0, 3, 5 /**/, 6, 0, 0 /**/, 2, 0, 7},
		/*											 						*/
		{3, 0, 1 /**/, 0, 2, 0 /**/, 0, 0, 0},
		{5, 0, 0 /**/, 0, 7, 0 /**/, 0, 0, 1},
		{0, 0, 0 /**/, 5, 9, 0 /**/, 8, 0, 3},
		/*											 						*/
		{7, 0, 3 /**/, 0, 0, 4 /**/, 0, 9, 0},
		{0, 0, 0 /**/, 2, 0, 0 /**/, 0, 8, 0},
		{0, 5, 0 /**/, 0, 0, 0 /**/, 6, 0, 0},
	}}

	testCases := []struct {
		desc     string
		input    pos
		expected []byte
	}{
		{
			desc:     "simple",
			input:    pos{4, 5},
			expected: []byte{3, 6, 8},
		},
		{
			desc:     "simple",
			input:    pos{8, 8},
			expected: []byte{2, 4},
		},
		{
			desc:     "simple",
			input:    pos{7, 0},
			expected: []byte{1, 4, 6, 9},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := sudoku.ValidNumbersAt(tC.input.x, tC.input.y)
			if !bytes.Equal(got, tC.expected) {
				t.Errorf("ValidNumbersAt for %v; got %v expected %v", tC.input, got, tC.expected)
			}
		})
	}
}
