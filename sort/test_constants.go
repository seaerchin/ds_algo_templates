package sort

var sorted  = []int{1, 2, 3, 4, 5}
var one = []int{1}
var none = []int{}
var scrambled = []int{2, 4, 3, 5, 1}

type testEntry struct { 
	input []int
	expected []int
}

// testTable is the data for sort
var testTable = []testEntry{
	{sorted, sorted},
	{one, one},
	{none, none},
	{scrambled, sorted},
}
