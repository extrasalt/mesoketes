package main

func (c *City) Defend(d Day) int {
	breachCount := 0
	breachCount += IncrIfBreach(c.N, d.N)
	breachCount += IncrIfBreach(c.E, d.E)
	breachCount += IncrIfBreach(c.S, d.S)
	breachCount += IncrIfBreach(c.W, d.W)
	c.N = max(c.N, d.N)
	c.S = max(c.S, d.S)
	c.E = max(c.E, d.E)
	c.W = max(c.W, d.W)

	return breachCount
}

func IncrIfBreach(current, atkStr int) int {
	if current < atkStr {
		return 1
	}

	return 0
}
