package main

import "testing"

func TestMax(t *testing.T) {
	if max(2, 3) != 3 {
		t.Errorf("Max function fails")
	}
}

func TestCalculateMaxDamage(t *testing.T) {
	d := CreateDay("Day 3$ T1 - W - X - 2: T1 - W - X - 4: T2 - N - X - 3: T2 - S - X - 4")

	cases := []struct {
		In  int
		Out int
	}{
		{d.W, 4},
		{d.N, 3},
		{d.S, 4},
	}

	for _, c := range cases {
		if c.In != c.Out {
			t.Errorf("Max damage fails. Expected %d got %d", c.Out, c.In)
		}
	}
}
