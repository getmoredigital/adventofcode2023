package main

import (
	"fmt"
	"getmoredigital/adventofcode2023/day1"
	"getmoredigital/adventofcode2023/day2"
	"getmoredigital/adventofcode2023/day3"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Specify your day e.g. go run main.go day1")
		return
	}

	switch os.Args[1] {
	case "day1":
		day1.Main()
	case "day2":
		day2.Main()
	case "day3":
		day3.Main()
	}
}
