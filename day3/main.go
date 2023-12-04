package day3

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

type schematicCoordinate struct {
	isDigit  bool
	isSymbol bool
	Value    interface{}
}

type Coordinate struct {
	x int
	y int
}

type schematicSymbol struct {
	Value    string
	Numbers  []int
	Location Coordinate
}

type schematicNumber struct {
	Value    int
	Location []Coordinate
	Symbols  []string
}

var numbers []schematicNumber
var symbols []schematicSymbol

func parseData(filepath string) [][]schematicCoordinate {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Error opening input file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var matrix [][]schematicCoordinate
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		var row []schematicCoordinate
		number := ""
		for i, char := range line {
			//isDigit and not last char
			if unicode.IsDigit(char) && i != len(line)-1 {
				number += string(char)
				row = append(row, schematicCoordinate{isDigit: true, isSymbol: false, Value: nil})
				continue
			} else if unicode.IsDigit(char) {
				number += string(char)
				row = append(row, schematicCoordinate{isDigit: true, isSymbol: false, Value: nil})
			}

			if number != "" {
				if v, err := strconv.Atoi(number); err == nil {
					numbers = append(numbers, schematicNumber{Value: v})
					for c := len(number); c != 0; c-- {
						lastIndex := len(row) - c
						numbers[len(numbers)-1].Location = append(numbers[len(numbers)-1].Location, Coordinate{lastIndex, y})
						row[lastIndex].Value = v
					}
				}
				number = ""
			}

			if char != '.' && !unicode.IsDigit(char) {
				stringified := string(char)
				s := schematicSymbol{Value: stringified, Location: Coordinate{i, y}}
				symbols = append(symbols, s)
				row = append(row, schematicCoordinate{isDigit: false, isSymbol: true, Value: stringified})
			} else if !unicode.IsDigit(char) {
				row = append(row, schematicCoordinate{isDigit: false, isSymbol: false, Value: nil})
			}
		}
		y++
		matrix = append(matrix, row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error reading file: %v\n", err)
	}
	return matrix
}

func printSchematic(grid [][]schematicCoordinate) {
	const (
		ColorReset = "\033[0m"
		ColorRed   = "\033[31m"
		ColorGreen = "\033[32m"
	)

	for y, row := range grid {
		for x, coord := range row {
			if coord.isDigit {
				if _, num, err := FindNumber(Coordinate{x, y}); err == nil {
					if len(num.Symbols) > 0 {
						fmt.Printf("%s%v %s", ColorGreen, coord.Value, ColorReset)
						continue
					}
				}
				fmt.Printf("%s%v %s", ColorRed, coord.Value, ColorReset)
			} else if coord.isSymbol {
				fmt.Printf("%s ", coord.Value)
			} else {
				fmt.Print("nil ")
			}
		}
		fmt.Println()
	}
}

func checkForNumber(x int, y int, grid [][]schematicCoordinate) (bool, int, schematicNumber) {
	maxY := len(grid) - 1
	maxX := len(grid[0]) - 1
	if x < 0 || y < 0 || y > maxY || x > maxX {
		return false, -1, schematicNumber{}
	}
	coord := grid[y][x]
	if coord.isDigit {
		if index, num, err := FindNumber(Coordinate{x, y}); err == nil {
			return true, index, num
		}
	}
	return false, -1, schematicNumber{}
}

func FindNumber(coord Coordinate) (int, schematicNumber, error) {
	for index, num := range numbers {
		for _, loc := range num.Location {
			if loc.x == coord.x && loc.y == coord.y {
				return index, num, nil
			}
		}
	}
	return -1, schematicNumber{}, errors.New("No number match")
}

func addAdjacentNumbers(grid [][]schematicCoordinate) {
	symIndex := 0
	for y, row := range grid {
		for x, coord := range row {
			if coord.isSymbol {
				var nums []int
				var symbl string
				if s, ok := coord.Value.(string); ok {
					symbl = s
				}
				// Check topRight
				if found, i, num := checkForNumber(x-1, y-1, grid); found {
					numbers[i].Symbols = append(numbers[i].Symbols, symbl)
					if !isDuplicate(nums, num.Value) {
						nums = append(nums, num.Value)
					}
				}
				// check topCenter
				if found, i, num := checkForNumber(x, y-1, grid); found {
					numbers[i].Symbols = append(numbers[i].Symbols, symbl)
					if !isDuplicate(nums, num.Value) {
						nums = append(nums, num.Value)
					}
				}
				// Check toLeft
				if found, i, num := checkForNumber(x+1, y-1, grid); found {
					numbers[i].Symbols = append(numbers[i].Symbols, symbl)
					if !isDuplicate(nums, num.Value) {
						nums = append(nums, num.Value)
					}
				}
				// Check left
				if found, i, num := checkForNumber(x-1, y, grid); found {
					numbers[i].Symbols = append(numbers[i].Symbols, symbl)
					if !isDuplicate(nums, num.Value) {
						nums = append(nums, num.Value)
					}
				}
				// Check right
				if found, i, num := checkForNumber(x+1, y, grid); found {
					numbers[i].Symbols = append(numbers[i].Symbols, symbl)
					if !isDuplicate(nums, num.Value) {
						nums = append(nums, num.Value)
					}
				}
				// check bottomRight
				if found, i, num := checkForNumber(x-1, y+1, grid); found {
					numbers[i].Symbols = append(numbers[i].Symbols, symbl)
					if !isDuplicate(nums, num.Value) {
						nums = append(nums, num.Value)
					}
				}
				// Check bottomCenter
				if found, i, num := checkForNumber(x, y+1, grid); found {
					numbers[i].Symbols = append(numbers[i].Symbols, symbl)
					if !isDuplicate(nums, num.Value) {
						nums = append(nums, num.Value)
					}
				}
				// Check bottomLeft
				if found, i, num := checkForNumber(x+1, y+1, grid); found {
					numbers[i].Symbols = append(numbers[i].Symbols, symbl)
					if !isDuplicate(nums, num.Value) {
						nums = append(nums, num.Value)
					}
				}
				symbols[symIndex].Numbers = nums
				symIndex++
			}
		}
	}
}

func isDuplicate(slice []int, element int) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

func Main() {
	schematic := parseData("day3/day3.txt")
	addAdjacentNumbers(schematic)
	sum := 0
	for _, num := range numbers {
		if len(num.Symbols) > 0 {
			sum += num.Value
		}
	}
	fmt.Println("The sum of part 1 is ", sum)
	sum = getGearRatios()
	fmt.Println("The sum of part 2 is ", sum)

}
