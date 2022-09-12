package piscine

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 && nb == 0 {
		return 1
	}
	if nb == 1 || nb == 0 {
		return nb
	} else if power == 0 {
		return 1
	}
	a := RecursivePower(nb, power/2)
	if power%2 == 0 {
		return a * a
	} else {
		return a * a * nb
	}
}
