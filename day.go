package main

//Day has several attacks
//The N, E, S, W fields are used to store the max
//attack for a given day
type Day struct {
	Attacks []Attack
	N       int
	E       int
	S       int
	W       int
}

const (
	north = iota
	east
	south
	west
)

//Direction constants are aliases to int
//Makes it more readable
type Direction int

//Attack specifies the attack direction and strength
//Other details like the tribe ID and weapon strength seem
//to be of little use. Should the weapon value change, it should
//get a new field here
type Attack struct {
	Dir      Direction
	Strength int
}

//CalculateMaxDamage calculates the max damage in any
//Given day for each side of the wall
//if there are two attacks on the same side, it sets the max
func (d *Day) CalculateMaxDamage() {
	for _, attack := range d.Attacks {
		switch attack.Dir {
		case north:
			d.N = max(d.N, attack.Strength)
		case east:
			d.E = max(d.E, attack.Strength)
		case south:
			d.S = max(d.S, attack.Strength)
		case west:
			d.W = max(d.W, attack.Strength)
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
