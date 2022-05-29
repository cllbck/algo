package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

type caseFunc struct {
	title string
	f     func(number []int)
}

var cases = []caseFunc{
	{"bubbleSort", bubbleSort},
	{"insertionSort", insertionSort},
	{"shellSort", shellSort},
	{"selectionSort", selectionSort},
	{"heapSort", heapSort},
}

func bubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}

func shellSort(arr []int) {
	for gap := len(arr) / 2; gap > 0; gap = gap / 2 {
		for j := gap; j < len(arr); j++ {
			for i := j - gap; i >= 0; i = i - gap {
				if arr[i+gap] > arr[i] {
					break
				} else {
					arr[i+gap], arr[i] = arr[i], arr[i+gap]
				}
			}
		}
	}
}

func selectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
}

func heapSort(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		heapify(arr, i, len(arr))
	}
	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, 0, i)
	}
}

func heapify(arr []int, i, n int) {
	for i >= 0 {
		largest := i
		l := 2*i + 1
		r := 2*i + 2
		if l < n && arr[l] > arr[largest] {
			largest = l
		}
		if r < n && arr[r] > arr[largest] {
			largest = r
		}
		if largest != i {
			arr[i], arr[largest] = arr[largest], arr[i]
			i = largest
		} else {
			break
		}
	}
}

func generateRandomArray(maxSize, maxValue int) []int {
	arr := make([]int, maxSize)
	for i := 0; i < maxSize; i++ {
		arr[i] = rand.Intn(maxValue)
	}
	return arr
}

func generateRandomArrays(maxSize, maxValue int) [][]int {
	arr1 := generateRandomArray(maxSize, maxValue)
	arr2 := make([]int, maxSize)
	arr3 := make([]int, maxSize)
	arr4 := make([]int, maxSize)
	arr5 := make([]int, maxSize)
	copy(arr2, arr1)
	copy(arr3, arr1)
	copy(arr4, arr1)
	copy(arr5, arr1)
	return [][]int{arr1, arr2, arr3, arr4, arr5}
}

func timeTrackWithResp(f func(number []int), n []int) string {
	start := time.Now()
	f(n)
	return fmt.Sprintf("%s", time.Since(start))
}

func deadLineWrapper(f func(number []int), n []int) string {
	ch := make(chan string)
	ticker := time.NewTicker(2 * time.Minute)

	go func() {
		ch <- timeTrackWithResp(f, n)
	}()
	for {
		select {
		case ret := <-ch:
			return ret
		case <-ticker.C:
			return "timeout"
		}
	}
}

func main() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	header := table.Row{"sort func/n elements"}
	for _, c := range cases {
		header = append(header, c.title)
	}
	rows := []table.Row{}
	for _, i := range []int{100, 1000, 10000, 100000, 1000000} {
		row := table.Row{i}
		arrays := generateRandomArrays(i, 1000000)
		useArray := 0
		for _, c := range cases {
			row = append(row, fmt.Sprintf("%s", deadLineWrapper(c.f, arrays[useArray])))
			useArray++
		}
		rows = append(rows, row)
	}
	t.AppendHeader(header)
	t.AppendRows(rows)
	t.Render()
}
