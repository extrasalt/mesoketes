package main

//Each day has several attacks
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

//Alias int to direction to create Direction constants
type Direction int
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
