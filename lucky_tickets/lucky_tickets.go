package main

import "fmt"

func LuckyTicketsCount(n int) int {
	table := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	count := 0
	for i := 1; i < n; i++ {
		table = fillTableFromPrev(table)
	}
	for _, item := range table {
		count += item * item
	}
	return count
}

func fillTableFromPrev(prev []int) []int {
	oldTableLen := len(prev)
	newTableLen := oldTableLen + 9
	newTable := make([]int, 0, newTableLen)
	for i := 0; i < newTableLen; i++ {
		newTableValue := 0
		for j := 0; j < 10; j++ {
			if i-j >= 0 && i-j < oldTableLen {
				newTableValue += prev[i-j]
			}
		}
		newTable = append(newTable, newTableValue)
	}
	return newTable
}

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d lucky tickets for %d number ticket\n", LuckyTicketsCount(i), i*2)
	}
}
