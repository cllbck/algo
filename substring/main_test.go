package main

import "testing"

type testCase struct {
	text, mask    string
	expectedIndex int
}

var testCases = []testCase{
	{
		text:          "123456789",
		mask:          "123",
		expectedIndex: 0,
	},
	{
		text:          "123456789",
		mask:          "456",
		expectedIndex: 3,
	},
	{
		text:          "123456789",
		mask:          "789",
		expectedIndex: 6,
	},
	{
		text:          "123456789",
		mask:          "123456789",
		expectedIndex: 0,
	},
	{
		text:          "Hello World",
		mask:          "World",
		expectedIndex: 6,
	},
	{
		text:          "Hello World",
		mask:          "Hello",
		expectedIndex: 0,
	},
	{
		text:          "Hello World",
		mask:          "World Hello",
		expectedIndex: -1,
	},
	{
		text:          "afsfwefwvwefgerhrtjrtasfva",
		mask:          "fwef",
		expectedIndex: 3,
	},
	{
		text:          "afsfwefwvwefgerhrtjrtasfva",
		mask:          "fwefrw0",
		expectedIndex: -1,
	},
}

func TestFindSubstringFullScan(t *testing.T) {
	for _, testStr := range testCases {
		index := FindSubstringFullScan(testStr.text, testStr.mask)
		if index != testStr.expectedIndex {
			t.Errorf("FindSubstringFullScan(%s, %s) = %d, expected %d", testStr.text, testStr.mask, index, testStr.expectedIndex)
		}
	}
}

func TestFindSubstringBM(t *testing.T) {
	for _, testStr := range testCases {
		index := FindSubstringBM(testStr.text, testStr.mask)
		if index != testStr.expectedIndex {
			t.Errorf("FindSubstringBM(%s, %s) = %d, expected %d", testStr.text, testStr.mask, index, testStr.expectedIndex)
		}
	}
}
