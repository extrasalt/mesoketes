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

	return d
}

func CreateAttack(input string) Attack {
	atk := Attack{}
	attributes := strings.Split(input, "-")
	atk.Dir = ParseDirection(attributes[1])
	atk.Strength = ParseStrength(attributes[3])
	return atk
}

func ParseDirection(s string) Direction {
	var d Direction

	switch s {
	case " N ":
		d = North
	case " E ":
		d = East
	case " S ":
		d = South
	case " W ":
		d = West
	}

	return d
}

func ParseStrength(s string) int {
	trimmedString := strings.Trim("s", " ")
	str, err := strconv.Atoi(trimmedString)
	if err != nil {
		panic(err)
	}

	return str
}
