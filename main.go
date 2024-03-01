package main

import (
	"fmt"
)

func main() {
//    fmt.Println("Testing Pawn Moves")
//    board := NewChessBoard()
//    board.PrintBoard()
//    board.MoveWhitePawn(E2, E4)
//    board.PrintBoard()
//    board.MoveBlackPawn(D7, D5)
//    board.PrintBoard()
//    board.MoveWhitePawn(E4, D5)
//    board.PrintBoard()
//    board.MoveBlackPawn(D5, D4)
//    board.PrintBoard()
    //row1Mask := rowMask(1)

    // Print row1Mask in binary
    //  fmt.Printf("Row 1 mask in binary: %064b\n", row1Mask)
    // fmt.Printf("White queens at %064b\n", board.WhiteQueens)
//    fmt.Println("Testing Pawn Promotion to Queen")
//    board2 := NewChessBoard()
//    board2.WhitePawns |= A7
//    board2.BlackRooks ^= A8
//    board2.PrintBoard()
//    board2.MoveWhitePawn(A7, A8)
//    fmt.Printf("White queens at %064b\n", board2.WhiteQueens)

//    fmt.Println("Test Pawn Capture Promotion")
//    board3 := NewChessBoard()
//    board3.PrintBoard()
//    board3.WhitePawns |= A7
//    board3.PrintBoard()
//    board3.MoveWhitePawn(A7, B8)
//    fmt.Printf("White queens at %064b\n", board3.WhiteQueens)
    //fmt.Println("How does the knight move?")
    //board4 := NewChessBoard()
    //board4.MoveKnight(B1, C3)
    //board4.PrintBoard()
    //board4.MoveKnight(C3, E4)
    //board4.PrintBoard()
    //board4.MoveKnight(E4, D6)
    //board4.PrintBoard()
    //board4.MoveKnight(D6, C4)
    //board4.PrintBoard()
    //board4.MoveKnight(C4, B6)
    //board4.PrintBoard()
    //board4.MoveKnight(B6, D5)
    //board4.PrintBoard()
    fmt.Println("Testing Knight capture")
    board5 := NewChessBoard()
    board5.MoveKnight(G1, F3)
    board5.PrintBoard()
    board5.MoveBlackPawn(E7, E5)
    board5.PrintBoard()
    board5.MoveKnight(F3, E5)
    board5.PrintBoard()

}

