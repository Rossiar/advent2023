package main

import (
	"cmp"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	aoc "github.com/rossiar/advent2023"
)

func main() {
	filename := os.Args[1]
	lines, err := aoc.ReadLinesFromFile(filename)
	if err != nil {
		panic(err.Error())
	}
	if err := Task1(lines); err != nil {
		panic(err.Error())
	}
}

var CardStrength = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'J': 1,
}

type HandType int32

const (
	_            = iota
	HighCard     = 0
	OnePair      = 1
	TwoPair      = 2
	ThreeOfAKind = 3
	FullHouse    = 4
	FourOfAKind  = 5
	FiveOfAKind  = 6
)

func GetHandType(hand string) HandType {
	charBuckets := make(map[rune]int, 0)
	for _, char := range hand {
		charBuckets[char]++
	}
	numJokers := charBuckets['J']
	delete(charBuckets, 'J')
	max := 0
	for _, count := range charBuckets {
		if count > max {
			max = count
		}
	}
	max += numJokers
	switch max {
	case 1:
		return HighCard
	case 2:
		if len(charBuckets) == 3 {
			return TwoPair
		}
		return OnePair
	case 3:
		if len(charBuckets) == 2 {
			return FullHouse
		}
		return ThreeOfAKind
	case 4:
		return FourOfAKind
	case 5:
		return FiveOfAKind
	}
	return HandType(0)
}

type hand struct {
	Bid   int
	Cards string
	Joker bool
	Type  HandType
}

func Task1(lines []string) error {
	hands := make([]hand, len(lines))
	for i, line := range lines {
		rawHand, rawBid, _ := strings.Cut(line, " ")
		bid, err := strconv.Atoi(rawBid)
		if err != nil {
			return err
		}
		hands[i] = hand{
			Cards: rawHand,
			Bid:   bid,
			Joker: strings.Contains(rawHand, "J"),
			Type:  GetHandType(rawHand),
		}
	}
	slices.SortFunc(hands, func(a, b hand) int {
		typeCmp := cmp.Compare(a.Type, b.Type)
		if typeCmp != 0 {
			return typeCmp
		}
		for i := 0; i < len(a.Cards); i++ {
			if a.Cards[i] != b.Cards[i] {
				aStrength := CardStrength[rune(a.Cards[i])]
				bStrength := CardStrength[rune(b.Cards[i])]
				//log.Printf("char %d: comparing %s (str %d) with %s (str %d)", i, a.Cards, aStrength, b.Cards, bStrength)
				if aStrength == bStrength {
					continue
				}
				if aStrength > bStrength {
					return 1
				}
				return -1
			}
		}
		log.Printf("hand %s equalled %s", a.Cards, b.Cards)
		return 0
	})
	//log.Print(hands)
	total := 0
	for i, hand := range hands {
		total += hand.Bid * (i + 1)
	}
	log.Print(total)
	return nil
}

func Task2(lines []string) error {
	return nil
}
