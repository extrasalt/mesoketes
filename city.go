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
	Counter []int
}

//Defend takes a day as input and returns
//the number of breaches for that particular day
//while also changing the wall's properties

//Defend gets an array of attacks from the day
//and returns the number of breaches
//if the city has to defend 3 attacks continously and successfully
//in a particular direction it resets the counter to 0

func (c *City) Defend(d Day) int {
	breachCount := 0
	c.Counter = make([]int, 4)

	for _, attack := range d.Attacks {
		switch attack.Dir {
		case north:
			if IncrIfBreach(c.N, attack.Strength) {
				c.N = attack.Strength
				breachCount += 1
			} else {
				c.Counter[north] += 1

				if c.Counter[north] == 3 {
					c.Counter[north] = 0
					c.N = 0
				}

			}

		case south:
			if IncrIfBreach(c.S, attack.Strength) {
				breachCount += 1
				c.S = attack.Strength
			} else {
				c.Counter[south] += 1

				if c.Counter[south] == 3 {
					c.Counter[south] = 0
					c.S = 0
				}
			}

		case east:
			if IncrIfBreach(c.E, attack.Strength) {
				breachCount += 1
				c.E = attack.Strength
			} else {
				c.Counter[east] += 1
				if c.Counter[east] == 3 {
					c.Counter[east] = 0
					c.E = 0
				}
			}

		case west:
			if IncrIfBreach(c.W, attack.Strength) {
				breachCount += 1
				c.W = attack.Strength
			} else {
				c.Counter[west] += 1

				if c.Counter[west] == 3 {
					c.Counter[west] = 0
					c.W = 0
				}
			}
		}
	}

	//	c.N = max(c.N, d.N)
	//	c.S = max(c.S, d.S)
	//	c.E = max(c.E, d.E)
	//	c.W = max(c.W, d.W)

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
