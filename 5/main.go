package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	aoc "github.com/rossiar/advent2023"
)

func main() {
	filename := os.Args[1]
	lines, err := aoc.ReadLinesFromFile(filename)
	if err != nil {
		panic(err.Error())
	}
	input, err := parse(lines)
	if err != nil {
		panic(err.Error())
	}
	if err := task(input); err != nil {
		panic(err.Error())
	}
}

type data struct {
	Seeds                 []int
	SeedsToSoil           mapping
	SoilToFertilizer      mapping
	FertilizerToWater     mapping
	WaterToLight          mapping
	LightToTemperature    mapping
	TemperatureToHumidity mapping
	HumidityToLocation    mapping
}

type mapping []entry

type entry struct {
	Dest  int
	Src   int
	Range int
}

func (m mapping) Map(src int) int {
	for _, e := range m {
		if src >= e.Src && src <= e.Src+e.Range {
			res := e.Dest + (src - e.Src)
			//fmt.Printf("mapped %d to %d\n", src, res)
			return res
		}
	}
	return src
}

func parse(lines []string) (*data, error) {
	var err error
	d := &data{
		Seeds:                 make([]int, 0),
		SeedsToSoil:           make([]entry, 0),
		SoilToFertilizer:      make([]entry, 0),
		FertilizerToWater:     make([]entry, 0),
		WaterToLight:          make([]entry, 0),
		LightToTemperature:    make([]entry, 0),
		TemperatureToHumidity: make([]entry, 0),
		HumidityToLocation:    make([]entry, 0),
	}
	seeds := strings.TrimPrefix(lines[0], "seeds: ")
	d.Seeds, err = aoc.ReadIntsFromString(seeds)
	if err != nil {
		return nil, err
	}
	var toAdd *mapping
	toAdd = &d.SeedsToSoil
	for i := 3; i < len(lines); i++ {
		line := lines[i]
		switch line {
		case "soil-to-fertilizer map:":
			toAdd = &d.SoilToFertilizer
			continue
		case "fertilizer-to-water map:":
			toAdd = &d.FertilizerToWater
			continue
		case "water-to-light map:":
			toAdd = &d.WaterToLight
			continue
		case "light-to-temperature map:":
			toAdd = &d.LightToTemperature
			continue
		case "temperature-to-humidity map:":
			toAdd = &d.TemperatureToHumidity
			continue
		case "humidity-to-location map:":
			toAdd = &d.HumidityToLocation
			continue
		case "":
			continue
		}
		nums, err := aoc.ReadIntsFromString(lines[i])
		if err != nil {
			return nil, err
		}
		*toAdd = append(*toAdd, entry{
			Dest:  nums[0],
			Src:   nums[1],
			Range: nums[2],
		})
	}

	return d, nil
}

func task(in *data) error {
	locations := make([]int, len(in.Seeds))
	for i, seed := range in.Seeds {
		soil := in.SeedsToSoil.Map(seed)
		fert := in.SoilToFertilizer.Map(soil)
		water := in.FertilizerToWater.Map(fert)
		light := in.WaterToLight.Map(water)
		temp := in.LightToTemperature.Map(light)
		hum := in.TemperatureToHumidity.Map(temp)
		loc := in.HumidityToLocation.Map(hum)
		fmt.Printf("seed %d, soil %d, fertilizer %d, water: %d, light: %d, temperature: %d, humidity: %d, location %d\n",
			seed, soil, fert, water, light, temp, hum, loc)
		locations[i] = loc
	}
	fmt.Printf("lowest location: %d\n", slices.Min(locations))
	return nil
}
