package main

import "testing"

func TestBattleRageOn(t *testing.T) {
	battle1 := CreateBattle("Day 1$ T1 - N - X - 3: T2 - S - X - 4: T3 - W - X - 2; Day 2$ T1 - E - X - 4: T2 - N - X - 3: T3 - S - X - 2; Day 3$ T1 - W - X - 3: T2 - E - X - 5: T3 - N - X - 2")

	battle2 := CreateBattle("Day 1$ T1 - S - X - 4: T1 - N - X - 2: T3 - W - X - 3; Day 2$ T2 - S - X - 5: T2 - N - X - 1: T3 - N - X - 3; Day 3$ T1 - W - X - 2: T1 - W - X - 4: T2 - N - X - 3: T2 - S - X - 4")

	battle3 := CreateBattle("Day 1$ T1 - E - X - 4; Day 2$ T1 - W - X - 3 : T2 - E - X - 3; Day 3$ T3 - N - X - 2: T2 - W - X - 4")

	city := &City{}
	breachCounts := []struct {
		out int
		exp int
	}{
		{battle1.RageOn(city), 6},
		{battle2.RageOn(city), 6},
		{battle3.RageOn(city), 4},
	}

	if len(battle3.Days) != 3 {
		t.Errorf("Expected 3 days in test case 3, got %d", len(battle3.Days))
	}

	if battle3.Days[1].E != 3 {
		t.Errorf("Expected 4, got %d", battle3.Days[0].E)
	}

	for _, c := range breachCounts {
		if c.exp != c.out {
			t.Errorf("Battle RageOn: Expected %d got %d", c.exp, c.out)
		}

	}
}
