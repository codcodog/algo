package main

import "fmt"

func main() {
	input := [][]byte{[]byte{'A', 'B', 'C', 'E'}, []byte{'S', 'F', 'C', 'S'}, []byte{'A', 'D', 'E', 'E'}}
	fmt.Println(exist(input, "ABCCED")) // expected: true
	fmt.Println(exist(input, "SEE"))    // expected: true
	fmt.Println(exist(input, "ABCB"))   // expected: false
}

//			(x, y-1)
// (x-1, y)  (x, y)   (x+1, y)
//			(x, y+1)
func exist(board [][]byte, word string) bool {
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			if isExist(board, y, x, word, 0) {
				return true
			}
		}
	}
	return false
}

func isExist(board [][]byte, y, x int, word string, i int) bool {
	if i == len(word) {
		return true
	}
	if y < 0 || y == len(board) || x < 0 || x == len(board[y]) {
		return false
	}
	if board[y][x] != word[i] {
		return false
	}
	var tmp byte
	board[y][x], tmp = '#', board[y][x] // avoid visit again
	exist := isExist(board, y, x+1, word, i+1) ||
		isExist(board, y, x-1, word, i+1) ||
		isExist(board, y+1, x, word, i+1) ||
		isExist(board, y-1, x, word, i+1)
	board[y][x] = tmp
	return exist
}
