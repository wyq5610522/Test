package main

import "testing"


func TestDel(t *testing.T) {
	var tests = []struct {
		a int
		b int
		expect int
	}{
		{2, 1, 1},
		{4, 2, 2},
		{6, 3, 3},
		{8, 4, 4},
		{10, 5, 5},
	}
	for _, test := range tests {
        if res := del(test.a, test.b); res != test.expect {
            t.Errorf("del(%d, %d) = %d", test.a, test.b, res)
        }
    }
}