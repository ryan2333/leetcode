package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	var (
		i   int
		j   int
		ok  bool
		row int
		col int
	)

	for i = 0; i < 9; i++ {
		var mp1, mp2, mp3 map[byte]bool
		mp1 = make(map[byte]bool)
		mp2 = make(map[byte]bool)
		mp3 = make(map[byte]bool)
		for j = 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			if _, ok = mp1[board[i][j]]; ok {
				return false
			}
			mp1[board[i][j]] = true
			if board[j][i] == '.' {
				continue
			}
			if _, ok = mp2[board[j][i]]; ok {
				return false
			}
			mp2[board[j][i]] = true
			row = (i%3)*3 + j%3
			col = (i/3)*3 + j/3
			if board[row][col] == '.' {
				continue
			}
			if _, ok = mp3[board[row][col]]; ok {
				return false
			}
			return true
		}
	}
	return true
}

func main() {
	var b1 [][]string
	b1 = [][]string{[]string{"1", "2", "3"}, []string{"2", "3", "4"}}
	fmt.Println("b1: ", b1)
}
