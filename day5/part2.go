package day5

func findLowestLocFromRange(start int, length int, convertMetrics map[string][]ConvertMetric, c chan int) {
	seedRange := make([]int, length)
	for i := 0; i < length; i++ {
		seedRange[i] = start + i
	}

	seeds := makeSeeds(seedRange, convertMetrics)
	var lowest Seed
	for _, seed := range seeds {
		if seed.Location < lowest.Location || lowest == (Seed{}) {
			lowest = seed
		}
	}

	c <- lowest.Location

}
