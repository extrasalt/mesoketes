package main

import (
	"strconv"
	"strings"
)

func CreateBattle(input string) Battle {
	b := Battle{}
	days := strings.Split(input, ";")
	for _, s := range days {
		d := CreateDay(s)
		b.Days = append(b.Days, d)
	}

	return b
}

func CreateDay(input string) Day {
	d := Day{}
	attacks := strings.Split(input, ":")
	for _, s := range attacks {
		a := CreateAttack(s)
		d.Attacks = append(d.Attacks, a)
	}
	d.CalculateMaxDamage()
	return d
}

func CreateAttack(input string) Attack {
	atk := Attack{}
	attributes := strings.Split(input, "-")
	atk.Dir = parseDirection(attributes[1])
	atk.Strength = parseStrength(attributes[3])
	return atk
}

func parseDirection(s string) Direction {
	var d Direction

	switch s {
	case " N ":
		d = north
	case " E ":
		d = east
	case " S ":
		d = south
	case " W ":
		d = west
	}

	return d
}

func parseStrength(s string) int {
	trimmedString := strings.Trim(s, " ")
	str, err := strconv.Atoi(trimmedString)
	if err != nil {
		panic(err)
	}

	return str
}
