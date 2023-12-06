package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Seed struct {
	Id          int
	Soil        int
	Fertilizer  int
	Water       int
	Light       int
	Temperature int
	Humidity    int
	Location    int
}

type Span struct {
	Start  int
	End    int
	Length int
}

type ConvertMetric struct {
	Target      Span
	Transformer Span
}

type Mode int

const (
	SeedMode Mode = iota
	AssignMapMode
	MapMode
)

func parseData(filepath string) ([]int, map[string][]ConvertMetric) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Error opening input file: %v\n", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var all []int
	convertMetrics := make(map[string][]ConvertMetric)
	mode := MapMode
	var currentMap string
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		target := Span{}
		transformer := Span{}
		if len(words) > 1 && words[0] == "seeds:" {
			mode = SeedMode
		} else if len(words) == 2 && words[1] == "map:" {
			mode = AssignMapMode
		} else {
			mode = MapMode
		}
		for i, word := range words {
			switch mode {
			case SeedMode:
				if num, err := strconv.Atoi(word); err == nil {
					all = append(all, num)
				}
			case AssignMapMode:
				if i == 0 {
					currentMap = word
				}
			case MapMode:
				if num, err := strconv.Atoi(word); err == nil {
					switch i {
					case 0:
						target.Start = num
					case 1:
						transformer.Start = num
					case 2:
						target.End = target.Start + num
						transformer.End = transformer.Start + num
						target.Length = num
						transformer.Length = num
						metric := ConvertMetric{Target: target, Transformer: transformer}
						convertMetrics[currentMap] = append(convertMetrics[currentMap], metric)
					}
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error reading file: %v\n", err)
	}

	return all, convertMetrics
}

func printSeeds(seeds []Seed) {
	for _, seed := range seeds {
		fmt.Printf("Seed %d: Soil: %d, Fertilizer: %d, Water: %d, Light: %d, Temperature: %d, Humidity: %d, Location: %d\n", seed.Id, seed.Soil, seed.Fertilizer, seed.Water, seed.Light, seed.Temperature, seed.Humidity, seed.Location)
	}
}

func PrintConvertMetrics(convertMetrics map[string][]ConvertMetric) {
	for key, metrics := range convertMetrics {
		fmt.Printf("Map %s\n", key)
		for _, metric := range metrics {
			fmt.Printf("Target: %d-%d, Transformer: %d-%d\n", metric.Target.Start, metric.Target.End, metric.Transformer.Start, metric.Transformer.End)
		}
	}
}

func transformItem(source int, metrics []ConvertMetric) int {
	for _, metric := range metrics {
		if source >= metric.Transformer.Start && source < metric.Transformer.End {
			diff := source - metric.Transformer.Start
			return metric.Target.Start + diff

		}
	}
	return source
}

func makeSeeds(ids []int, convertMetrics map[string][]ConvertMetric) []Seed {
	var seeds []Seed
	for _, id := range ids {
		seed := Seed{Id: id}
		//Soil
		metrics := convertMetrics["seed-to-soil"]
		seed.Soil = transformItem(id, metrics)
		//Fertilizer
		metrics = convertMetrics["soil-to-fertilizer"]
		seed.Fertilizer = transformItem(seed.Soil, metrics)
		//Water
		metrics = convertMetrics["fertilizer-to-water"]
		seed.Water = transformItem(seed.Fertilizer, metrics)
		//Light
		metrics = convertMetrics["water-to-light"]
		seed.Light = transformItem(seed.Water, metrics)
		//Temperature
		metrics = convertMetrics["light-to-temperature"]
		seed.Temperature = transformItem(seed.Light, metrics)
		//Humidity
		metrics = convertMetrics["temperature-to-humidity"]
		seed.Humidity = transformItem(seed.Temperature, metrics)
		//Location
		metrics = convertMetrics["humidity-to-location"]
		seed.Location = transformItem(seed.Humidity, metrics)
		seeds = append(seeds, seed)
	}
	return seeds
}

func Main() {
	seedIds, convertMetrics := parseData("day5/day5.txt")
	seeds := makeSeeds(seedIds, convertMetrics)
	var lowest Seed
	for _, seed := range seeds {
		if seed.Location < lowest.Location || lowest == (Seed{}) {
			lowest = seed
		}
	}
	fmt.Println("The lowest location is", lowest.Location)

	if len(seedIds)%2 == 0 {
		results := make(chan int, len(seedIds)/2)

		for i := 0; i < len(seedIds); i += 2 {
			go findLowestLocFromRange(seedIds[i], seedIds[i+1], convertMetrics, results)
		}

		for i := 0; i < len(seedIds); i += 2 {
			fmt.Printf("Result %d: %d\n", i+1, <-results)
		}

		// Close the channel
		close(results)

	}

}
