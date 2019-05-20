package main

import "testing"


func TestAdd(t *testing.T) {
	var tests = []struct {
		a int
		b int
		expect int
	}{
		{1, 1, 2},
		{2, 2, 4},
		{3, 3, 6},
		{4, 4, 8},
		{5, 5, 10},
	}
	for _, test := range tests {
        if sum := add(test.a, test.b); sum != test.expect {
            t.Errorf("add(%d, %d) = %d", test.a, test.b, sum)
        }
    }
}