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

func (board *ChessBoard) MoveWhitePawn(currentPosition uint64, newPosition uint64) {
    allPieces := board.WhitePawns | board.WhiteKnights | board.WhiteBishops |
    board.WhiteRooks | board.WhiteQueens | board.WhiteKing | board.BlackPawns |
    board.BlackKnights | board.BlackBishops | board.BlackRooks |
    board.BlackQueens

    if newPosition & allPieces!= 0 {
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
func (board *ChessBoard) CapturePiece(newPosition uint64) {

    // Calculate the bit mask for the destination square
    destMask :=  newPosition
    fmt.Printf("destMask in binary: %b\n", newPosition)

    // Perform a bitwise AND operation between the destination square's bit mask and the bit masks
    // representing each piece on the board
    whitePawnsMask := board.WhitePawns & destMask
    whiteKnightsMask := board.WhiteKnights & destMask
    whiteBishopsMask := board.WhiteBishops & destMask
    whiteRooksMask := board.WhiteRooks & destMask
    whiteQueensMask := board.WhiteQueens & destMask
    whiteKingMask := board.WhiteKing & destMask
    blackPawnsMask := board.BlackPawns & destMask
    blackKnightsMask := board.BlackKnights & destMask
    blackBishopsMask := board.BlackBishops & destMask
    blackRooksMask := board.BlackRooks & destMask
    blackQueensMask := board.BlackQueens & destMask
    blackKingMask := board.BlackKing & destMask

    // Check which piece is in the destination square based on the result of the bitwise AND operation
    if whitePawnsMask != 0 {
        fmt.Println("There is a white pawn in the destination square")
        // Remove the white pawn from the destination square
        board.WhitePawns ^= destMask
    } else if whiteKnightsMask != 0 {
        fmt.Println("There is a white knight in the destination square")
        // Remove the white knight from the destination square
        board.WhiteKnights ^= destMask
    } else if whiteBishopsMask != 0 {
        fmt.Println("There is a white bishop in the destination square")
        // Remove the white bishop from the destination square
        board.WhiteBishops ^= destMask
    } else if whiteRooksMask != 0 {
        fmt.Println("There is a white rook in the destination square")
        // Remove the white rook from the destination square
        board.WhiteRooks ^= destMask
    } else if whiteQueensMask != 0 {
        fmt.Println("There is a white queen in the destination square")
        // Remove the white queen from the destination square
        board.WhiteQueens ^= destMask
    } else if whiteKingMask != 0 {
        fmt.Println("There is a white king in the destination square")
        // Remove the white king from the destination square
        board.WhiteKing ^= destMask
    } else if blackPawnsMask != 0 {
        fmt.Println("There is a black pawn in the destination square")
        // Remove the black pawn from the destination square
        board.BlackPawns ^= destMask
    } else if blackKnightsMask != 0 {
        fmt.Println("There is a black knight in the destination square")
        // Remove the black knight from the destination square
        board.BlackKnights ^= destMask
    } else if blackBishopsMask != 0 {
        fmt.Println("There is a black bishop in the destination square")
        // Remove the black bishop from the destination square
        board.BlackBishops ^= destMask
    } else if blackRooksMask != 0 {
        fmt.Println("There is a black rook in the destination square")
        // Remove the black rook from the destination square
        board.BlackRooks ^= destMask
    } else if blackQueensMask != 0 {
        fmt.Println("There is a black queen in the destination square")
        // Remove the black queen from the destination square
        board.BlackQueens ^= destMask
    } else if blackKingMask != 0 {
        fmt.Println("There is a black king in the destination square")
        // Remove the black king from the destination square
        board.BlackKing ^= destMask
    } else {
        fmt.Println("The destination square is empty")
    }
}
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
}

