package main

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

const numberForPow = 2

type caseFunc struct {
	title string
	f     func(number, pow int) float64
}

var cases = []caseFunc{
	{"builtInPow", builtInPow},
	{"simplePow", simplePow},
	{"simplePow2", simplePow2},
	{"simpleBinaryPow", simpleBinaryPow},
}

func builtInPow(number, pow int) float64 {
	return math.Pow(float64(number), float64(pow))
}

// without checking overflow
func simplePow(number, pow int) float64 {
	ret := 1
	for i := 1; i <= pow; i++ {
		ret *= number
	}
	return float64(ret)
}

// without checking overflow
func simplePow2(number, pow int) float64 {
	if pow == 0 {
		return 1
	}
	tempNumber := number
	i := 1
	for ; i <= pow/2; i *= 2 {
		number *= number
	}
	for ; i < pow; i++ {
		number *= tempNumber
	}
	return float64(number)
}

// without checking overflow
func simpleBinaryPow(number, pow int) float64 {
	if pow == 0 {
		return 1
	}
	temp := 1
	for ; pow > 1; pow /= 2 {
		if pow%2 == 1 {
			temp *= number
		}
		number *= number
	}
	temp *= number
	return float64(temp)
}

func main() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	header := table.Row{"pow of 2"}
	for _, c := range cases {
		header = append(header, c.title)
	}
	rows := []table.Row{}
	for i := 0; i <= 62; i += 2 {
		row := table.Row{i}
		for _, c := range cases {
			start := time.Now()
			resp := c.f(numberForPow, i)
			row = append(row, fmt.Sprintf("%e(%s)", resp, time.Since(start)))
		}
		rows = append(rows, row)
	}
	t.AppendHeader(header)
	t.AppendRows(rows)
	t.Render()
}
