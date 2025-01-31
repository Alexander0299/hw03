package main

import (
	"fmt"
	"strings"
)

func chessboard(size int) string {
	var board strings.Builder

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			switch (i + j) % 2 {
			case 0:
				board.WriteString(" ")
			default:
				board.WriteString("#")
			}
		}
		board.WriteString("\n")
	}

	return board.String()
}

func main() {
	var size int
	fmt.Print("Введите размер шахматной доски: ")
	fmt.Scanf("%d", &size)

	cb := chessboard(size)
	fmt.Println(cb)
}
