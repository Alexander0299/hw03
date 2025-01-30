package main

import (
	"fmt"
)

func chessboard(size int) string {
	var board string
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			switch (i + j) % 2 {
			case 0:
				board += " "
			case 1:
				board += "#"
			}
		}
		board += "\n"
	}
	return board
}
func main() {

	size := 8

	cb := chessboard(size)
	fmt.Println(cb)
}
