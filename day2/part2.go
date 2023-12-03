package day2

func minTotal(g Game) Pull {
	red := 0
	blue := 0
	green := 0

	for _, p := range g.Pulls {
		if p.Red > red {
			red = p.Red
		}
		if p.Blue > blue {
			blue = p.Blue
		}
		if p.Green > green {
			green = p.Green
		}
	}

	return Pull{red, blue, green}

}

func powerSet(p Pull) int {
	return p.Red * p.Blue * p.Green
}
