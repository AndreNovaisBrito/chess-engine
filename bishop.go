package main

import "fmt"

func (board *ChessBoard) MoveBishop(currentPosition, newPosition uint64) {
	blackPieces := board.BlackPawns | board.BlackKnights | board.BlackBishops | board.BlackRooks | board.BlackQueens
	whitePieces := board.WhitePawns | board.WhiteKnights | board.WhiteBishops | board.WhiteRooks | board.WhiteQueens | board.WhiteKing

	if isBishopMove(currentPosition, newPosition) {
		if !isPathClear(currentPosition, newPosition, whitePieces, blackPieces) {
			fmt.Println("There is a piece blocking the way")
			return
		}

		if currentPosition&board.WhiteBishops != 0 {
			if newPosition&whitePieces != 0 {
				fmt.Println("Can't capture own piece")
				return
			}
			if newPosition&blackPieces != 0 {
				board.CapturePiece(newPosition)
			}
			board.WhiteBishops ^= currentPosition
			board.WhiteBishops |= newPosition
			return
		}

		if currentPosition&board.BlackBishops != 0 {
			if newPosition&blackPieces != 0 {
				fmt.Println("Can't capture own piece")
				return
			}
			if newPosition&whitePieces != 0 {
				board.CapturePiece(newPosition)
			}
			board.BlackBishops ^= currentPosition
			board.BlackBishops |= newPosition
			return
		}
	}

	fmt.Println("Invalid Move")
}

func isBishopMove(currentPosition, newPosition uint64) bool {
	diff := int(newPosition - currentPosition)
	return diff%7 == 0 || diff%9 == 0
}

func isPathClear(currentPosition, newPosition uint64, whitePieces, blackPieces uint64) bool {
	pieces := whitePieces | blackPieces

	if (newPosition-currentPosition)%9 == 0 {
		step := int64(9)
		if newPosition < currentPosition {
			step = -9
		}
		for pos := currentPosition + uint64(step); pos != newPosition; pos += uint64(step) {
			if pieces&pos != 0 {
				return false
			}
		}
	} else if (newPosition-currentPosition)%7 == 0 {
		step := int64(7)
		if newPosition < currentPosition {
			step = -7
		}
		for pos := currentPosition + uint64(step); pos != newPosition; pos += uint64(step) {
			if pieces&pos != 0 {
				return false
			}
		}
	}
	return true
}
