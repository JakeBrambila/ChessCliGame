package main

import "fmt"

type Board struct {
	//board is a 2D array of pieces
	tiles [8][8]string
}

// NewBoard creates a new board with the default piece positions
func (b *Board) StartingBoard() {
	for i := 0; i < 8; i++ {
		if i == 0 {
			b.tiles[i][0] = "[♜]"
			b.tiles[i][1] = "[♞]"
			b.tiles[i][2] = "[♝]"
			b.tiles[i][3] = "[♛]"
			b.tiles[i][4] = "[♚]"
			b.tiles[i][5] = "[♝]"
			b.tiles[i][6] = "[♞]"
			b.tiles[i][7] = "[♜]"
		} else if i == 1 {
			for j := 0; j < 8; j++ {
				b.tiles[i][j] = "[♟]"
			}
		} else if i == 6 {
			for j := 0; j < 8; j++ {
				b.tiles[i][j] = "[♙]"
			}
		} else if i == 7 {
			b.tiles[i][0] = "[♖]"
			b.tiles[i][1] = "[♘]"
			b.tiles[i][2] = "[♗]"
			b.tiles[i][3] = "[♕]"
			b.tiles[i][4] = "[♔]"
			b.tiles[i][5] = "[♗]"
			b.tiles[i][6] = "[♘]"
			b.tiles[i][7] = "[♖]"
		} else {
			for j := 0; j < 8; j++ {
				b.tiles[i][j] = "[ ]"
			}
		}
	}
}

func (b *Board) PrintBoard() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Printf("%s ", b.tiles[i][j])
		}
		fmt.Println()
	}
}
