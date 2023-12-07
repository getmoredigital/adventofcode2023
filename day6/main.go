package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseData(filepath string) ([]int, []int) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Error opening input file: %v\n", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var distances []int
	var times []int
	distanceMode := false
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		for _, word := range words {
			if word == "Time:" {
				continue
			} else if word == "Distance:" {
				distanceMode = true
			} else {
				if num, err := strconv.Atoi(word); err == nil {
					if distanceMode {
						distances = append(distances, num)
					} else {
						times = append(times, num)
					}
				}
			}
		}
	}

	return distances, times
}

func races(time int, distance int) []int {
	var result []int
	for i := 1; i < time-1; i++ {
		d := (time - i) * i
		if d > distance {
			result = append(result, i)
		}
	}
	return result
}

func Main() {
	dists, times := parseData("day6/sample.txt")

	var waysToWin int
	for i := 0; i < len(times); i++ {
		t := races(times[i], dists[i])
		if waysToWin == 0 {
			waysToWin = len(t)
		} else {
			waysToWin *= len(t)
		}
	}

	fmt.Println("Ways to win part 1:", waysToWin)

	var longTimeStr string
	var longDistStr string
	for i := 0; i < len(times); i++ {
		str := strconv.Itoa(times[i])
		str2 := strconv.Itoa(dists[i])
		longTimeStr = longTimeStr + str
		longDistStr = longDistStr + str2
	}

	longTime, _ := strconv.Atoi(longTimeStr)
	longDist, _ := strconv.Atoi(longDistStr)
	t := races(longTime, longDist)
	fmt.Println("Ways to win part 2:", len(t))
}
