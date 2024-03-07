
package main

import "fmt"

func (board *ChessBoard) MoveRook(currentPosition, newPosition uint64) {
    //blackPieces := board.BlackPawns |
    //board.BlackKnights | board.BlackBishops | board.BlackRooks |
    //board.BlackQueens
    //whitePieces:= board.WhitePawns | board.WhiteKnights | board.WhiteBishops |
    //board.WhiteRooks | board.WhiteQueens | board.WhiteKing 
    //Check if it's a valid rook move

    for i := 8 ; i <= 56  ; i+=8 {
        if currentPosition << i & newPosition != 0 {
           fmt.Println("Valid move Poggers") 
          //Check if there is any piece in the way


          //Check if there is a piece in the destination
              //Capture Piece
          //
        }
    }
    for i:=1 ; i>=8 ; i++ {
        if currentPosition << i & newPosition != 0 {
           fmt.Println("Valid move Poggers") 
        }
    }
//            //WhiteKnight logic
//            if currentPosition & board.WhiteKnights != 0 {
//                //Capture logic
//                if newPosition & whitePieces != 0 {
//                    fmt.Println("Can't capture own piece")
//                    return
//                }
//                if newPosition & blackPieces != 0 {
//                    board.CapturePiece(newPosition)
//                }
//                board.WhiteKnights ^= currentPosition
//                board.WhiteKnights |= newPosition
//                return
//            }
//            //BlackKnight logic
//            if currentPosition & board.BlackKnights != 0 {
//                if newPosition & blackPieces != 0 {
//                    fmt.Println("Can't capture own piece")
//                    return
//                }
//                if newPosition & whitePieces != 0 {
//                    board.CapturePiece(newPosition)
//                }
//                board.BlackKnights ^= currentPosition
//                board.BlackKnights |= newPosition
//                return
//            }
//        }
    fmt.Println("Invalid Move")
    }
