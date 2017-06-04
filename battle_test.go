package main

import "testing"

func TestBattleRageOn(t *testing.T) {
	battle := CreateBattle("Day 1$ T1 - N - X - 3: T2 - S - X - 4: T3 - W - X - 2; Day 2$ T1 - E - X - 4: T2 - N - X - 3: T3 - S - X - 2; Day 3$ T1 - W - X - 3: T2 - E - X - 5: T3 - N - X - 2")
	city := &City{}

	count := battle.RageOn(city)

	if count != 6 {
		t.Fatalf("Battle RageOn fails. expected 6 got %d instead", count)
	}

}
