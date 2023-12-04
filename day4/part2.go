package day4

import "fmt"

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

func printInstances(games []Game) {
	instances := countInstances(games)
	for i,game  := range games {
		fmt.Printf("Game %s: %d instances\n", game.Id, instances[i])
	}
}
