package day5

var convertMetrics map[string][]ConvertMetric

func processIds(ids []int) Seed {
	seeds := makeSeeds(ids, convertMetrics)
	var lowest Seed
	for _, seed := range seeds {
		if seed.Location < lowest.Location || lowest == (Seed{}) {
			lowest = seed
		}
	}
	return lowest
}

func processRange(start, overallEnd, currentEnd, chunkSize int, currentLowest Seed) Seed {
	if start > overallEnd {
		return currentLowest
	}

	if currentEnd > overallEnd {
		currentEnd = overallEnd
	}

	idChunk := make([]int, currentEnd-start+1)
	for i := range idChunk {
		idChunk[i] = start + i
	}

	l := processIds(idChunk)
	if l.Location < currentLowest.Location || currentLowest == (Seed{}) {
		currentLowest = l
	}

	return processRange(currentEnd+1, overallEnd, currentEnd+chunkSize, chunkSize, currentLowest)
}

func findLowestLocFromRange(start int, length int, c chan int) {
	const chunkSize = 1000
	end := start + length

	lowest := processRange(start, end, (start+chunkSize)-1, chunkSize, Seed{})

	c <- lowest.Location

}
