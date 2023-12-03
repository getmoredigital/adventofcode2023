package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

type Calibration struct {
	First string
	Last  string
}

func printTuples(tuples []Calibration) {
	for _, tuple := range tuples {
		fmt.Println(tuple.First, tuple.Last)
	}
}

func parseDigits(filepath string) []Calibration {
	file, err := os.Open(filepath)
	var all []Calibration

	if err != nil {
		log.Fatalf("Error opening input file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var lineCalibration Calibration
		line := scanner.Text()
		for _, char := range line {
			if unicode.IsDigit(char) {
				lineCalibration.First = string(char)
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				lineCalibration.Last = string(line[i])
				break
			}
		}

		all = append(all, lineCalibration)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error reading file: %v\n", err)
	}

	return all
}

func Main() {
	calibrationsJustDigits := parseDigits("day1/day1.txt")
	calibrationsWithWords := parseDigitsAndWords("day1/day1.txt")
	sum := 0

	for _, cal := range calibrationsJustDigits {
		join := cal.First + cal.Last
		if combinedInt, err := strconv.Atoi(join); err == nil {
			sum += combinedInt
		}
	}

	fmt.Println("The part 1 sum is", sum)
	sum = 0

	for _, cal := range calibrationsWithWords {
		join := cal.First + cal.Last
		if combinedInt, err := strconv.Atoi(join); err == nil {
			sum += combinedInt
		}
	}
	fmt.Println("The part 2 sum is", sum)
}
