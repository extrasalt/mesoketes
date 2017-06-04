package main

import "testing"

func TestMax(t *testing.T) {
	if max(2, 3) != 3 {
		t.Errorf("Max function fails")
	}
}

func TestCalculateMaxDamage(t *testing.T) {
	d := CreateDay("Day 3$ T1 - W - X - 2: T1 - W - X - 4: T2 - N - X - 3: T2 - S - X - 4")

	if d.W != 4 {
		t.Errorf("Max damage fails. Expected 4 got %d", d.W)
	}

	if d.N != 3 {
		t.Errorf("Max Damage fails expected 3, got %d", d.N)
	}
	if d.S != 4 {
		t.Errorf("Adding damage to day fails")
	}
}
