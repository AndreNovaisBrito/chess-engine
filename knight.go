package main
import "fmt"

func (board *ChessBoard) MoveKnight(currentPosition, newPosition uint64) {
    blackPieces := board.BlackPawns |
    board.BlackKnights | board.BlackBishops | board.BlackRooks |
    board.BlackQueens
    whitePieces:= board.WhitePawns | board.WhiteKnights | board.WhiteBishops |
    board.WhiteRooks | board.WhiteQueens | board.WhiteKing 
    //Check if it's a valid knight move
    if currentPosition << 17 & newPosition != 0 ||
        currentPosition << 15 & newPosition != 0 || 
        currentPosition << 10 & newPosition != 0 || 
        currentPosition << 6 & newPosition != 0 ||
        currentPosition >> 17 & newPosition != 0 || 
        currentPosition >> 15 & newPosition != 0 || 
        currentPosition >> 10 & newPosition != 0 || 
        currentPosition >> 6 & newPosition != 0 {

            //WhiteKnight logic
            if currentPosition & board.WhiteKnights != 0 {
                //Capture logic
                if newPosition & whitePieces != 0 {
                    fmt.Println("Can't capture own piece")
                    return
                }
                if newPosition & blackPieces != 0 {
                    board.CapturePiece(newPosition)
                }
                board.WhiteKnights ^= currentPosition
                board.WhiteKnights |= newPosition
                return
            }
            //BlackKnight logic
            if currentPosition & board.BlackKnights != 0 {
                if newPosition & blackPieces != 0 {
                    fmt.Println("Can't capture own piece")
                    return
                }
                if newPosition & whitePieces != 0 {
                    board.CapturePiece(newPosition)
                }
                board.BlackKnights ^= currentPosition
                board.BlackKnights |= newPosition
                return
            }
        }
        fmt.Println("Invalid Move")
    }
