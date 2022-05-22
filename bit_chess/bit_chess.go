package main

import "fmt"

func getKingBitBoardMoves(pos uint) uint {
	var k uint = 1 << pos
	var kNoLeftMoves uint = 0xfefefefefefefefe & k
	var kNoRightMoves uint = 0x7f7f7f7f7f7f7f7f & k
	return (kNoLeftMoves << 7) | (k << 8) | (kNoRightMoves << 9) |
		(kNoLeftMoves >> 1) | (kNoRightMoves << 1) |
		(kNoLeftMoves >> 9) | (k >> 8) | (kNoRightMoves >> 7)
}

func getKnightBitBoardMoves(pos uint) uint {
	var k uint = 1 << pos
	var nA uint = 0xFeFeFeFeFeFeFeFe
	var nAB uint = 0xFcFcFcFcFcFcFcFc
	var nH uint = 0x7f7f7f7f7f7f7f7f
	var nGH uint = 0x3f3f3f3f3f3f3f3f
	return nGH&(k<<6|k>>10) |
		nH&(k<<15|k>>17) |
		nA&(k<<17|k>>15) |
		nAB&(k<<10|k>>6)
}

func getRookBitBoardMoves(pos uint) uint {
	var k uint = 1 << pos
	var mask uint
	var maskH uint = 0xff
	for i := 8; i < 64; i += 8 {
		if maskH&k > 0 {
			mask |= maskH ^ k // fill horizontal moves
		}
		maskH <<= 8
		mask |= k<<i | k>>i //fill vertical moves
	}
	return mask
}

func main() {
	var pos uint = 49
	fmt.Printf("bit mask for king moves %v from position %v\n", getKingBitBoardMoves(pos), pos)
	fmt.Printf("bit mask for knigth moves %v from position %v\n", getKnightBitBoardMoves(pos), pos)
	fmt.Printf("bit mask for rook moves %v from position %v\n", getRookBitBoardMoves(pos), pos)
}
