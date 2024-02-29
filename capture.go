
package main

import "fmt"

func (board *ChessBoard) CapturePiece(newPosition uint64) {

    // Calculate the bit mask for the destination square
    destMask :=  newPosition

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
