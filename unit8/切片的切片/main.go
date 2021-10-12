package main

import (
	"fmt"
	"strings"
)

func main() {
	//创建一个井字板游戏
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[1][1] = "O"
	board[0][1] = "X"
	board[0][2] = "O"
	board[1][0] = "X"
	board[2][0] = "O"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
