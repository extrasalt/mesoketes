package main

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
