package main

//Battle consists of several days
type Battle struct {
	Days []Day
}

//RageOn imposes the battle over a city
//This function returns the number of times
//the city wall was breached throughout the battle
func (b *Battle) RageOn(c *City) int {
	breachCount := 0

	for _, day := range b.Days {
		breachCount += c.Defend(day)
	}

	return breachCount
}
