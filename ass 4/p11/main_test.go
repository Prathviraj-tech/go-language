package main

import "testing"

func TestSubtract(t *testing.T) {
	tbl := []struct {
		a, b, r int
	}{
		{20, 10, 10},
		{50, 20, 30},
	}

	for _, v := range tbl {
		if Subtract(v.a, v.b) != v.r {
			t.Fail()
		}
	}
}
