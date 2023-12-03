package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

var writtenDigits []string = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var writtenDigitsBackward []string = []string{"eno", "owt", "eerht", "ruof", "evif", "xis", "neves", "thgie", "enin"}
var writtenDigitDict map[string]string = map[string]string{
	"one": "1", "eno": "1", "two": "2", "owt": "2", "three": "3", "eerht": "3",
	"four": "4", "ruof": "4", "five": "5", "evif": "5", "six": "6", "xis": "6",
	"seven": "7", "neves": "7", "eight": "8", "thgie": "8", "nine": "9", "enin": "9",
}

func checkList(list []string, crt string) (bool, string) {
	for _, str := range list {
		if strings.Contains(crt, str) {
			return true, str
		}
	}
	return false, ""
}

func parseDigitsAndWords(filepath string) []Calibration {
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
		var current string

		for _, char := range line {
			if unicode.IsDigit(char) {
				lineCalibration.First = string(char)
				current = ""
				break
			}
			current += string(char)
			if check, value := checkList(writtenDigits, current); check {
				lineCalibration.First = writtenDigitDict[value]
				current = ""
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				lineCalibration.Last = string(line[i])
				current = ""
				break
			}
			current += string(line[i])
			if check, value := checkList(writtenDigitsBackward, current); check {
				lineCalibration.Last = writtenDigitDict[value]
				current = ""
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
