package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Pull struct {
	Red   int
	Blue  int
	Green int
}

type Game struct {
	Id    int
	Pulls []Pull
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
		words := strings.Fields(line)
		var currentGame Game
		var value int
		currentPull := Pull{}

		for i, word := range words {
			if num, err := strconv.Atoi(word); err == nil {
				value = num
			} else if strings.HasSuffix(word, ":") {
				word = strings.TrimSuffix(word, ":")
				if gameNumber, err := strconv.Atoi(word); err == nil {
					currentGame.Id = gameNumber
				}
			} else if strings.HasSuffix(word, ",") || strings.HasSuffix(word, ";") {
				color := strings.Trim(word, ",;")
				switch color {
				case "red":
					currentPull.Red = value
				case "blue":
					currentPull.Blue = value
				case "green":
					currentPull.Green = value
				}

				if strings.HasSuffix(word, ";") {
					currentGame.Pulls = append(currentGame.Pulls, currentPull)
					currentPull = Pull{}
				}
			} else if i == len(words)-1 {
				switch word {
				case "red":
					currentPull.Red = value
				case "blue":
					currentPull.Blue = value
				case "green":
					currentPull.Green = value
				}
				currentGame.Pulls = append(currentGame.Pulls, currentPull)
			}
		}

		all = append(all, currentGame)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error reading file: %v\n", err)
	}

	return all
}

func printGames(games []Game) {
	for _, g := range games {
		fmt.Printf("Game ID: %d\n", g.Id)
		fmt.Println("Pulls:")
		for _, p := range g.Pulls {
			fmt.Printf("Red: %d, Blue: %d, Green: %d\n", p.Red, p.Blue, p.Green)
		}
		fmt.Println()
	}
}

func checkGame(g Game, max Pull) bool {
	for _, p := range g.Pulls {
		if p.Blue > max.Blue || p.Red > max.Red || p.Green > max.Green {
			return false
		}
	}

	return true
}

func Main() {
	games := parseData("day2/day2.txt")
	sum := 0
	checkPossible := Pull{12, 14, 13}
	for _, g := range games {
		if checkGame(g, checkPossible) {
			sum += g.Id
		}
	}
	fmt.Println("The part 1 sum is", sum)
	sum = 0
	for _, g := range games {
		min := minTotal(g)
		sum += powerSet(min)
	}
	fmt.Println("The part 2 sum is", sum)
}
