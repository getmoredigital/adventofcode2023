package day3

func getGearRatios() int {
	sum := 0
	for _, symb := range symbols {
		if symb.Value == "*" && len(symb.Numbers) == 2 {
			sum += (symb.Numbers[0] * symb.Numbers[1])
		}
	}
	return sum
}
