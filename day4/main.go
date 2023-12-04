package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"math"
)

type Game struct {
	Id    string
	Winners []int
	MyNumbers []int
	Matches []int
	Points	int
}


func getMatches(winners []int, myNumbers []int) []int {
	var matches []int
	for _, winner := range winners {
		for _, myNumber := range myNumbers {
			if winner == myNumber {
				matches = append(matches, winner)
			}
		}
	}
	return matches
}

func getPoints(common []int) int {
	if len(common) < 2 {
		return len(common)
	}

	return int(math.Pow(2, float64(len(common)-1)))
}

func parseData(filepath string) []Game {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Error opening input file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var all []Game
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Fields(line)
		var id string
		var winners []int
		var myNumbers []int
		firstSet := true
		for _, item := range items {
			if strings.HasSuffix(item, ":") {
				id = strings.TrimSuffix(item, ":")
			} else if item == "|" {
				firstSet = false
			} else if num, err := strconv.Atoi(item); err == nil && firstSet {
				winners = append(winners, num)
			} else if num, err := strconv.Atoi(item); err == nil {
				myNumbers = append(myNumbers, num)
			}
		}
		matches := getMatches(winners, myNumbers)
		points := getPoints(matches)
		all = append(all, Game{Id: id, Winners: winners, MyNumbers: myNumbers, Points: points, Matches: matches})
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error reading file: %v\n", err)
	}

	return all
}

func printGames(games []Game) {
	for _, game := range games {
		fmt.Printf("Game %s: %d points\n", game.Id, game.Points)
	}
}

func Main(){
	games := parseData("day4/day4.txt")
	sum := 0
	for _, game := range games {
		sum += game.Points
	}
	fmt.Println("The part 1 sum is", sum)
	total := 0
	allInstancesOfGames :=  countInstances(games)
	for _, instanceCount := range allInstancesOfGames {
		total += instanceCount
	}
	fmt.Println("The part 2 total is", total)
}