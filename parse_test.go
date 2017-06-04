package main

import "testing"

func TestParseDirection(t *testing.T) {
	directionTests := []struct {
		in  string
		out Direction
	}{
		{" N ", North},
		{" S ", South},
		{" E ", East},
		{" W ", West},
	}

	for _, c := range directionTests {
		if ParseDirection(c.in) != c.out {
			t.Errorf("Error in ParseDirection:%s is not %s", c.in, string(c.out))
		}
	}
}

func TestParseStrength(t *testing.T) {
	if ParseStrength(" 3") != 3 {
		t.Errorf("ParseStrength fails")
	}
}
