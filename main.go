package main

import (
	"fmt"
)

func main() {
	board := NewChessBoard()
	board.PrintBoard()
    board.MoveWhitePawn(E2, E4)
    board.PrintBoard()
    board.MoveBlackPawn(D7, D5)
    board.PrintBoard()
    board.WhitePawnCapture(E4, D5)
    board.PrintBoard()
    board.MoveBlackPawn(D5, D4)
    board.PrintBoard()
    row1Mask := rowMask(1)

    // Print row1Mask in binary
    fmt.Printf("Row 1 mask in binary: %064b\n", row1Mask)
  // fmt.Printf("White queens at %064b\n", board.WhiteQueens)

    board2 := NewChessBoard()
    board2.WhitePawns |= A7
    board2.BlackRooks ^= A8
    board2.MoveWhitePawn(A7, A8)
    fmt.Printf("White queens at %064b\n", board2.WhiteQueens)
    board2.PrintPiece() 
}

