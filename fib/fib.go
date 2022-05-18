package main

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

const (
	timeout = 10
)

type caseFunc struct {
	title string
	f     func(number int) int
}

var PHI = (1 + math.Sqrt(5)) / 2
var cases = []caseFunc{
	{"fibonacciRecursion", fibonacciRecursion},
	{"fibonacciIteration", fibonacciIteration},
	{"fibonacciBinet", fibonacciBinet},
	{"fibonacciMatrix", fibonacciMatrix},
}

func fibonacciRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacciRecursion(n-1) + fibonacciRecursion(n-2)
}

func fibonacciIteration(n int) int {
	if n == 0 {
		return 0
	}
	f1 := 1
	f2 := 1
	for i := 2; i < n; i++ {
		f3 := f1 + f2
		f1 = f2
		f2 = f3
	}
	return f2
}

func fibonacciBinet(n int) int {
	var num float64 = float64(n)
	return int(((math.Pow(PHI, num) - math.Pow(1-PHI, num)) / math.Sqrt(5)) + 0.5)
}

type Matrix [4]int

func (m *Matrix) Multiply(n Matrix) {
	m[0], m[1], m[2], m[3] = m[0]*n[0]+m[1]*n[2], m[0]*n[1]+m[1]*n[3], m[2]*n[0]+m[3]*n[2], m[2]*n[1]+m[3]*n[3]
}

func fibonacciMatrix(n int) int {
	if n == 0 {
		return 0
	}
	res := Matrix{1, 0, 1, 0}
	base := Matrix{1, 1, 1, 0}
	for n > 1 {
		if (n & 1) == 1 {
			res.Multiply(base)
		}
		base.Multiply(base)
		n >>= 1
	}

	res.Multiply(base)
	return res[3]
}

func timeTrackWithResp(f func(number int) int, n int) string {
	start := time.Now()
	resp := f(n)
	return fmt.Sprintf("%v(%s)", resp, time.Since(start))
}

func deadLineWrapper(f func(number int) int, n int) string {
	ch := make(chan string)
	ticker := time.NewTicker(timeout * time.Second)

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
	header := table.Row{"fib numbers"}
	for _, c := range cases {
		header = append(header, c.title)
	}
	rows := []table.Row{}
	for i := 0; i <= 90; i += 10 {
		row := table.Row{i}
		for _, c := range cases {
			row = append(row, fmt.Sprintf("%s", deadLineWrapper(c.f, i)))
		}
		rows = append(rows, row)
	}
	t.AppendHeader(header)
	t.AppendRows(rows)
	t.Render()
}
