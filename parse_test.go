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

func TestCreateAttack(t *testing.T) {
	attack := CreateAttack("T1 - N - X - 2")
	if attack.Dir != North {
		t.Errorf("Error setting direction")
	}

	if attack.Strength != 2 {
		t.Errorf("Error setting strength")
	}
}

func TestCreateDay(t *testing.T) {
	day := CreateDay("Day 1$ T1 - S - X - 4: T1 - N - X - 2: T3 - W - X - 3")
	if len(day.Attacks) != 3 {
		t.Errorf("Unable to add attacks to day")
	}
}

func TestCreateBattle(t *testing.T) {
	battleInput := "Day 1$ T1 - E - X - 4; Day 2$ T1 - W - X - 3 : T2 - E - X - 3; Day 3$ T3 - N - X - 2: T2 - W - X - 4"
	battle := CreateBattle(battleInput)

	if len(battle.Days) != 3 {
		t.Errorf("Unable to add days to battle")
	}
}
