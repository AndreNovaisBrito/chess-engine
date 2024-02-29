package main

import (
	"fmt"
)

func (board *ChessBoard) PromoteWhitePawn(newPosition uint64){
    if newPosition & rowMask(7) != 0 {
        board.WhitePawns^= newPosition
        board.WhiteQueens |= newPosition
        return
    }
    fmt.Println("White Pawn didn't promote D:")
}

func (board *ChessBoard) PromoteBlackPawn(newPosition uint64){
    if newPosition & rowMask(0) != 0 {
        board.BlackPawns ^= newPosition
        board.BlackPawns |= newPosition
    }
    fmt.Println("Black Pawn didn't promote D:")
}

func (board *ChessBoard) MoveWhitePawn(currentPosition uint64, newPosition uint64) {
    allPieces := board.WhitePawns | board.WhiteKnights | board.WhiteBishops |
    board.WhiteRooks | board.WhiteQueens | board.WhiteKing | board.BlackPawns |
    board.BlackKnights | board.BlackBishops | board.BlackRooks |
    board.BlackQueens

    if newPosition & allPieces!= 0 {
        fmt.Println("There is already a piece in the destination")
        board.WhitePawnCapture(currentPosition, newPosition)
        board.PromoteWhitePawn(newPosition)
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
            board.PromoteWhitePawn(newPosition)
            return
        }
    }
    fmt.Println("Invalid Move")
}

func (board *ChessBoard) MoveBlackPawn(currentPosition uint64, newPosition uint64) {
    allPieces := board.WhitePawns | board.WhiteKnights | board.WhiteBishops |
    board.WhiteRooks | board.WhiteQueens | board.WhiteKing | board.BlackPawns |
    board.BlackKnights | board.BlackBishops | board.BlackRooks |
    board.BlackQueens

    if newPosition & allPieces!= 0 {
        fmt.Println("There is already a piece in the destination")
        return
    }
    //Check if the pawn is in it's starting position
    if (currentPosition & board.BlackPawns & 0xFF000000000000) != 0 {
        if ((currentPosition >> 16) & newPosition) != 0 {
            board.BlackPawns^= currentPosition
            board.BlackPawns |= newPosition
            return
        }
    }
    if (currentPosition & board.BlackPawns) != 0 {
        if((currentPosition >> 8) &  newPosition) != 0 {
            board.BlackPawns ^= currentPosition
            board.BlackPawns |= newPosition
            board.PromoteBlackPawn(newPosition)
            return
        }
    }
    fmt.Println("Invalid Move")
}



func (board *ChessBoard) WhitePawnCapture(currentPosition uint64, newPosition uint64) {
    allPieces := board.WhitePawns | board.WhiteKnights | board.WhiteBishops |
    board.WhiteRooks | board.WhiteQueens | board.WhiteKing | board.BlackPawns |
    board.BlackKnights | board.BlackBishops | board.BlackRooks |
    board.BlackQueens

    if newPosition & allPieces == 0 {
        fmt.Println("There is no piece at the destination")
        return
    }
    if (currentPosition & board.WhitePawns) != 0 {
        if((currentPosition << 7) &  newPosition) != 0 || ((currentPosition << 9) &  newPosition) != 0  {
            board.WhitePawns^= currentPosition
            board.CapturePiece(newPosition)
            board.WhitePawns |= newPosition
            return
        }
    }
    fmt.Println("Invalid Move")
}


func (board *ChessBoard) BlackPawnCapture(currentPosition uint64, newPosition uint64) {
    allPieces := board.WhitePawns | board.WhiteKnights | board.WhiteBishops |
    board.WhiteRooks | board.WhiteQueens | board.WhiteKing | board.BlackPawns |
    board.BlackKnights | board.BlackBishops | board.BlackRooks |
    board.BlackQueens

    if newPosition & allPieces == 0 {
        fmt.Println("There is no piece at the destination")
        return
    }
    if (currentPosition & board.BlackPawns) != 0 {
        if((currentPosition >> 7) &  newPosition) != 0 || ((currentPosition >> 9) &  newPosition) != 0  {
            board.BlackPawns^= currentPosition
            board.CapturePiece(newPosition)
            board.BlackPawns |= newPosition
            return
        }
    }
    fmt.Println("Invalid Move")
}

