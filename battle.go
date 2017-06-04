package main

func (b *Battle) RageOn(c *City) int {
	breachCount := 0

	for _, day := range b.Days {
		breachCount += c.Defend(day)
	}

	return breachCount
}
