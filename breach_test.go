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
