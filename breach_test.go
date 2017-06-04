package main

import "testing"

func TestIncrIfBreach(t *testing.T) {
	current := 5
	atkStr := 7
	breachCount := 0

	breachCount += IncrIfBreach(current, atkStr)
	if breachCount != 1 {
		t.Errorf("Positive breach doesn't increment the counter")
	}

	current, atkStr = atkStr, current
	if breachCount != 1 {
		t.Errorf("Counter is incremented even if there is no breach")
	}
}

func TestDefend(t *testing.T) {
	day := CreateDay("Day 1$ T1 - S - X - 4: T1 - S - X - 2: T3 - W - X - 3")

	city := City{}

	b := city.Defend(day)

	if b != 2 {
		t.Errorf("Defend returns bad count: Expected 2, got %d", b)
	}

	if city.S != 4 {
		t.Errorf("city.S expected 4, got %d", city.S)
	}

	if city.W != 3 {
		t.Errorf("city.W expected 3, got %d", city.W)
	}
}
