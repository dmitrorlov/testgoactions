package src

import (
	"testing"
)

func TestSum(t *testing.T) {
	tables := []struct {
		a   int
		b   int
		sum int
	}{
		{
			1,
			2,
			3,
		},
		{
			3,
			4,
			7,
		},
		{
			3,
			5,
			8,
		},
	}

	for _, table := range tables {
		sum := Sum(table.a, table.b)
		if table.sum != sum {
			t.Errorf("unexpected sum. a: %v, b: %v, expected: %v, actual: %v", table.a, table.b, table.sum, sum)
		}
	}
}
