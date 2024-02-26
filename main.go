package main

import (
	"fmt"
)

// Constants to represent each square on the board
const (
	A1 = uint64(1) << iota
	B1
	C1
	D1
	E1
	F1
	G1
	H1
	A2
	B2
	C2
	D2
	E2
	F2
	G2
	H2
	A3
	B3
	C3
	D3
	E3
	F3
	G3
	H3
	A4
	B4
	C4
	D4
	E4
	F4
	G4
	H4
	A5
	B5
	C5
	D5
	E5
	F5
	G5
	H5
	A6
	B6
	C6
	D6
	E6
	F6
	G6
	H6
	A7
	B7
	C7
	D7
	E7
	F7
	G7
	H7
	A8
	B8
	C8
	D8
	E8
	F8
	G8
	H8
)

// ChessBoard represents the chess board using bitboards
type ChessBoard struct {
	WhitePawns   uint64
	WhiteKnights uint64
	WhiteBishops uint64
	WhiteRooks   uint64
	WhiteQueens  uint64
	WhiteKing    uint64
	BlackPawns   uint64
	BlackKnights uint64
	BlackBishops uint64
	BlackRooks   uint64
	BlackQueens  uint64
	BlackKing    uint64
}

// NewChessBoard creates a new instance of a ChessBoard with the default starting position
func NewChessBoard() *ChessBoard {
	board := &ChessBoard{
		WhitePawns:   A2 | B2 | C2 | D2 | E2 | F2 | G2 | H2,
		WhiteKnights: B1 | G1,
		WhiteBishops: C1 | F1,
		WhiteRooks:   A1 | H1,
		WhiteQueens:  D1,
		WhiteKing:    E1,
		BlackPawns:   A7 | B7 | C7 | D7 | E7 | F7 | G7 | H7,
		BlackKnights: B8 | G8,
		BlackBishops: C8 | F8,
		BlackRooks:   A8 | H8,
		BlackQueens:  D8,
		BlackKing:    E8,
	}
	return board
}

// PrintBoard prints the current state of the board to the console
func (board *ChessBoard) PrintBoard() {
	for row := 7; row >= 0; row-- {
		for col := 0; col <= 7; col++ {
			bit := getBit(board.WhitePawns|board.WhiteKnights|board.WhiteBishops|board.WhiteRooks|board.WhiteQueens|board.WhiteKing|
				board.BlackPawns|board.BlackKnights|board.BlackBishops|board.BlackRooks|board.BlackQueens|board.BlackKing, uint8(row), uint8(col))
			if bit == 1 {
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func getBit(bitboard uint64, row, col uint8) uint8 {
	return uint8((bitboard >> (row*8 + col)) & 1)
}

func (board *ChessBoard) MovePawn(currentPosition uint64, newPosition uint64) {
    if newPosition & board.WhitePawns != 0 {
        fmt.Println("There is already a piece in the destination")
        return
    }
    //Check if the pawn is in it's starting position
    if (currentPosition & board.WhitePawns & 0xFF00) != 0 {
        if ((currentPosition << 16) & newPosition) != 0 {
            board.WhitePawns ^= currentPosition
            board.WhitePawns |= newPosition
            return
        }
    }
    if (currentPosition & board.WhitePawns) != 0 {
        if((currentPosition << 8) &  newPosition) != 0 {
            board.WhitePawns ^= currentPosition
            board.WhitePawns |= newPosition
            return
        }
    }
    fmt.Println("Invalid Move")
}

func main() {
	board := NewChessBoard()
	board.PrintBoard()
    board.MovePawn(E2, E4)
    board.PrintBoard()
    board.MovePawn(E4, E5)
    board.PrintBoard()
}

