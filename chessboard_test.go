package main

import (
    "testing";
    "fmt"
)
func TestMovePawn(t *testing.T) {
    fmt.Println("Testing if Pawn can move one Square")
    board := NewChessBoard()
    board.MovePawn(A2, A3)

    expected := A3
    result := board.WhitePawns
    if (result & expected) == 0 {
        t.Errorf("Expected %d, got %d", expected, result)
    }

    fmt.Println("Done")
    fmt.Println("Testing if Pawn can move two squares")

    board2 := NewChessBoard()
    board2.MovePawn(A2, A4)

    expected = A4
    result = board2.WhitePawns

    if (result & expected) == 0 {
        t.Errorf("Expected %d, got %d", expected, result)
    }
    fmt.Println("Done")

    fmt.Println("Testing if Pawn can't move two squares while not in it's starting position, it should print Invalid Move")
    board3 := NewChessBoard()
    board3.MovePawn(A2, A3)
    board3.MovePawn(A3, A5)

    expected = A3
    result = board3.WhitePawns

    if (result & expected) == 0 {
        t.Errorf("Expected %d, got %d", expected, result)
    }
    fmt.Println("Done")
}
