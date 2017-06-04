package main

import (
	"strconv"
	"strings"
)

//CreateBattle creates a battle from parsing the string
//It delegates the cration of days to underlying functions
//and populates the slice by appending returned structs
func CreateBattle(input string) Battle {
	b := Battle{}
	days := strings.Split(input, ";")
	for _, s := range days {
		d := CreateDay(s)
		b.Days = append(b.Days, d)
	}

	return b
}

//CreateDay creates a day from parsing the given string
//It delegates the creation of Attack structs and then
//populates the slice by appending the returned structs
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

//CreateAttack creates a Attack struct from the given string
//and populates it with parsed Strength and Direction
func CreateAttack(input string) Attack {
	atk := Attack{}
	attributes := strings.Split(input, "-")
	atk.Dir = parseDirection(attributes[1])
	atk.Strength = parseStrength(attributes[3])
	return atk
}

//Sets the direction from the given string
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

//Cleans the string and parses the integer
func parseStrength(s string) int {
	trimmedString := strings.Trim(s, " ")
	str, err := strconv.Atoi(trimmedString)
	if err != nil {
		panic(err)
	}

	return str
}
