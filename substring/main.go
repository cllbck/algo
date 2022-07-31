package main

import "fmt"

func FindSubstringFullScan(text, mask string) int {
	textRune := []rune(text)
	maskRune := []rune(mask)
	for i := 0; i <= len(textRune)-len(maskRune); i++ {
		for j := len(maskRune) - 1; j >= 0 && textRune[i+j] == maskRune[j]; j-- {
			if j == 0 {
				return i
			}
		}
	}
	return -1
}

func FindSubstringBM(text, mask string) int {
	textLen := len(text)
	maskLen := len(mask)
	switch {
	case maskLen == 0:
		return 0
	case maskLen > textLen:
		return -1
	case textLen == maskLen:
		if text == mask {
			return 0
		}
	}

	textRune := []rune(text)
	maskRune := []rune(mask)
	shiftTable := makeShiftTable(mask)

	for i := 0; i <= len(textRune)-len(maskRune); {
		j := len(maskRune) - 1
		for ; j >= 0 && textRune[i+j] == maskRune[j]; j-- {
		}
		if j == -1 {
			return i
		}
		i += shiftTable[textRune[i+len(maskRune)-1]]
	}
	return -1
}

func makeShiftTable(mask string) [256]int {
	maskRune := []rune(mask)
	shiftTable := [256]int{}
	for i := 0; i < 256; i++ {
		shiftTable[i] = len(maskRune)
	}
	for i := 0; i < len(maskRune)-1; i++ {
		shiftTable[maskRune[i]] = len(maskRune) - i - 1
	}
	return shiftTable
}

func main() {
	fmt.Println(FindSubstringFullScan("STRONGSTRING", "STRING"))
	fmt.Println(FindSubstringFullScan("Hello world", "d"))
	fmt.Println(FindSubstringFullScan("Hello world", "v"))
	fmt.Println(FindSubstringBM("STRONGSTRING", "STRING"))
	fmt.Println(FindSubstringBM("Hello world", "d"))
	fmt.Println(FindSubstringBM("Hello world", "v"))
}
