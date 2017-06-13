package main

//City has a wall
type City struct {
	wall
}

type strength struct {
	N int //Northern Wall's strength
	E int //Eastern
	S int //Southern
	W int //Western
}

type wall struct {
	strength
	//Counter []int
}

//Defend takes a day as input and returns
//the number of breaches for that particular day
//while also changing the wall's properties
func (c *City) Defend(d Day) int {
	breachCount := 0

	dirs := []struct {
		CityStrength int
		DayStrength  int
	}{
		{c.N, d.N},
		{c.E, d.E},
		{c.S, d.S},
		{c.W, d.W},
	}

	for _, dir := range dirs {
		if IncrIfBreach(dir.CityStrength, dir.DayStrength) {
			breachCount += 1
		}
	}
	c.N = max(c.N, d.N)
	c.S = max(c.S, d.S)
	c.E = max(c.E, d.E)
	c.W = max(c.W, d.W)

	return breachCount
}

//IncrIfBreach compares two values and returns
//whether or not the current value is smaller
//than the next value. Ints used instead of bool
//To facilitate incrementing the counter.
func IncrIfBreach(current, atkStr int) bool {
	if current < atkStr {
		return true
	}

	return false
}
