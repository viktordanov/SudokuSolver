package main

import (
	"fmt"
	"strings"
)

type Sudoku struct {
	Board [][]byte
}

func (s *Sudoku) ValidNumbersAt(x, y int) []byte {
	found := [10]bool{}

	// Cols
	for i := 0; i < len(s.Board); i++ {
		if x == i {
			continue
		}
		if s.Board[i][y] != 0 {
			found[s.Board[i][y]] = true // meaning it's taken
		}
	}
	// Rows
	for i := 0; i < len(s.Board); i++ {
		if y == i {
			continue
		}
		if s.Board[x][i] != 0 {
			found[s.Board[x][i]] = true // meaning it's taken
		}
	}

	// Subsquare
	xx, yy := x/3*3, y/3*3
	for i := xx; i < xx+3; i++ {
		for j := yy; j < yy+3; j++ {
			if i != x && j != y && s.Board[i][j] != 0 {
				found[s.Board[i][j]] = true // meaning it's taken
			}
		}
	}

	free := []byte{}
	for i, num := range found {
		if !num && i != 0 {
			free = append(free, byte(i))
		}
	}
	return free
}

func (s *Sudoku) Solve() *Sudoku {
	return s.solveUtil(0)
}

func (s *Sudoku) solveUtil(i int) *Sudoku {
	if i == 9*9 {
		return s
	}
	x, y := i/9, i%9
	if s.Board[x][y] != 0 {
		return s.solveUtil(i + 1)
	}
	for _, val := range s.ValidNumbersAt(x, y) {
		s.Board[x][y] = val

		if s.solveUtil(i+1) != nil {
			return s
		}

		s.Board[x][y] = 0
	}
	return nil
}

// Print prints the sudoku
func (s Sudoku) String() string {
	b := &strings.Builder{}
	for i, line := range s.Board {
		if i%3 == 0 {
			fmt.Fprintln(b, "+---------+---------+---------+")
		}
		for j, num := range line {
			if j%3 == 0 {
				fmt.Fprint(b, "|")
			}
			if num == 0 {
				fmt.Fprint(b, " - ")
			} else {
				fmt.Fprintf(b, " %d ", num)
			}
			if j == 8 {
				fmt.Fprint(b, "|")
			}
		}
		if i == 8 {
			fmt.Fprint(b, "\n+---------+---------+---------+")
		}
		fmt.Fprintln(b)
	}

	return b.String()
}
