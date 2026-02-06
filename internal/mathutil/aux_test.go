package mathutil

import (
	"testing"
)

func TestRoundTo2Decimals(t *testing.T) {
	cases := []struct {
		value    float64
		expected float64
	}{
		{123.456, 123.46},
		{123.454, 123.45},
		{123.0, 123.00},
		{-123.456, -123.46},
	}

	for _, c := range cases {
		got := RoundTo2Decimals(c.value)
		if got != c.expected {
			t.Errorf("RoundTo2Decimals(%v) == %v, expected %v", c.value, got, c.expected)
		}
	}
}
