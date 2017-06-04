package main

func (d *Day) CalculateMaxDamage() {
	for _, attack := range d.Attacks {
		switch attack.Dir {
		case North:
			d.N = max(d.N, attack.Strength)
		case East:
			d.E = max(d.E, attack.Strength)
		case South:
			d.S = max(d.S, attack.Strength)
		case West:
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
