package main

import "testing"

// "half ticket numbers len": "count lucky ticket"
var testCases = map[int]int{
	1:  10,
	2:  670,
	3:  55252,
	4:  4816030,
	5:  432457640,
	6:  39581170420,
	7:  3671331273480,
	8:  343900019857310,
	9:  32458256583753952,
	10: 3081918923741896840,
}

func TestLuckyTicketsCount(t *testing.T) {
	for n, expectedLuckyTicketsCount := range testCases {
		gotLuckyTicketsCount := LuckyTicketsCount(n)
		if gotLuckyTicketsCount != expectedLuckyTicketsCount {
			t.Errorf("case(%d) got <%d> wanted <%d>", n, gotLuckyTicketsCount, expectedLuckyTicketsCount)
		}
	}
}
