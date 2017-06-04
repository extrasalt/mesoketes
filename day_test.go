package main

import "testing"

func TestMax(t *testing.T) {
	if max(2, 3) != 3 {
		t.Errorf("Max function fails")
	}
}

func TestCalculateMaxDamage(t *testing.T) {
	d := CreateDay("Day 2$ T1 - N - X - 4: T2 - N - X - 3: T3 - S - X - 2")
	d.CalculateMaxDamage()
	if d.N != 4 {
		t.Errorf("Max damage fails")
	}

	if d.S != 2 {
		t.Errorf("Adding damage to day fails")
	}
}
