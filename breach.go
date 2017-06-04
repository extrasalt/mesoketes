package main

func (c *City) Defend(d Day) {
	c.BreachCount += IncrIfBreach(c.N, d.N)
	c.BreachCount += IncrIfBreach(c.E, d.E)
	c.BreachCount += IncrIfBreach(c.S, d.S)
	c.BreachCount += IncrIfBreach(c.W, d.W)
	c.N = max(c.N, d.N)
	c.S = max(c.S, d.S)
	c.E = max(c.E, d.E)
	c.W = max(c.W, d.W)
}

func IncrIfBreach(current, atkStr int) int {
	if current < atkStr {
		return 1
	}

	return 0
}
