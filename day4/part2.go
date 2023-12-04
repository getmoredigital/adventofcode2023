package day4

func countInstances(games []Game) []int {
	n := len(games)
	if n == 0 {
		return nil
	}

	instances := make([]int, n)
	for i := range instances {
		instances[i] = 1
	}

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n && j <= i+len(games[i].Matches); j++ {
			instances[j] += instances[i]
		}
	}

	return instances
}
