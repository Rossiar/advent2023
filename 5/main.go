package main

import (
	"fmt"
	"os"
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
	Seeds                 []seedRange
	SeedsToSoil           mapping
	SoilToFertilizer      mapping
	FertilizerToWater     mapping
	WaterToLight          mapping
	LightToTemperature    mapping
	TemperatureToHumidity mapping
	HumidityToLocation    mapping
}

type seedRange struct {
	Start int
	End   int
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
		Seeds:                 make([]seedRange, 0),
		SeedsToSoil:           make([]entry, 0),
		SoilToFertilizer:      make([]entry, 0),
		FertilizerToWater:     make([]entry, 0),
		WaterToLight:          make([]entry, 0),
		LightToTemperature:    make([]entry, 0),
		TemperatureToHumidity: make([]entry, 0),
		HumidityToLocation:    make([]entry, 0),
	}
	seeds := strings.TrimPrefix(lines[0], "seeds: ")
	seedRanges, err := aoc.ReadIntsFromString(seeds)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(seedRanges); i += 2 {
		d.Seeds = append(d.Seeds,
			seedRange{
				Start: seedRanges[i],
				End:   seedRanges[i] + seedRanges[i+1],
			},
		)
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
			Dest:  nums[1],
			Src:   nums[0],
			Range: nums[2],
		})
	}

	return d, nil
}

func task(in *data) error {
	location := 0
	found := false
	for {
		hum := in.HumidityToLocation.Map(location)
		temp := in.TemperatureToHumidity.Map(hum)
		light := in.LightToTemperature.Map(temp)
		water := in.WaterToLight.Map(light)
		fert := in.FertilizerToWater.Map(water)
		soil := in.SoilToFertilizer.Map(fert)
		seed := in.SeedsToSoil.Map(soil)
		for _, seedRange := range in.Seeds {
			if seed >= seedRange.Start && seed <= seedRange.End {
				found = true
			}
		}
		if found {
			break
		}
		location++
	}
	fmt.Printf("lowest location: %d\n", location)
	return nil
}
