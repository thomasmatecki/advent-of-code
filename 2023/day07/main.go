package main

import (
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	cards [5]int
	bid   int
	str   *string
}

var data, _ = os.ReadFile("07.txt")
var standardDeck = [13]byte{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}
var jokerDeck = [13]byte{'J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A'}
var handExpr = regexp.MustCompile(`([\w\d]{5}) (\d+)`)

func initDeck(cardChars [13]byte) map[byte]int {

	deck := make(map[byte]int, len(cardChars))

	for idx, c := range cardChars {
		deck[c] = idx
	}

	return deck
}

func initHands(handStrs *[]string, deck map[byte]int) []Hand {

	hands := make([]Hand, len(*handStrs)-1)
	for hIdx, handStr := range *handStrs {
		match := handExpr.FindStringSubmatch(handStr)
		if match == nil {
			continue
		}
		hand := &hands[hIdx]
		for cIdx, card := range []byte(match[1]) {
			hand.cards[cIdx] = deck[card]
		}

		hand.str = &match[1]
		hand.bid, _ = strconv.Atoi(match[2])
	}
	return hands
}

func allocWildcard(cardCounts *map[int]int, maxCount *int, wildcard int) {
	wildcardCount := (*cardCounts)[wildcard]
	if wildcardCount == 0 {
		return
	}

	delete(*cardCounts, wildcard)

	newMaxCount := 0
	newMaxCard := 0

	for card, count := range *cardCounts {
		if count > newMaxCount {
			newMaxCount = count
			newMaxCard = card
		}
	}

	(*cardCounts)[newMaxCard] += wildcardCount
	*maxCount = newMaxCount + wildcardCount
}

func (hand *Hand) typeRank(wildcards []int) int {
	maxCount := 0
	cardCounts := map[int]int{}

	for _, ci := range hand.cards {
		cardCounts[ci]++
		if cardCounts[ci] > maxCount {
			maxCount = cardCounts[ci]
		}
	}

	for _, wildcard := range wildcards {
		allocWildcard(&cardCounts, &maxCount, wildcard)
	}

	switch len(cardCounts) {
	case 1: // five of a kind
		return 7
	case 2:
		if maxCount == 4 {
			return 6 // four of a kind
		} else {
			return 5 //full house
		}
	case 3:
		if maxCount == 3 {
			return 4 // three of a kind
		} else {
			return 3 // two pair
		}
	case 4:
		return 2 // one pair
	default:
		return 1 // high card
	}
}
func (hand *Hand) cardsLessThan(otherHand *Hand) bool {
	for i := 0; i < 5; i++ {
		if hand.cards[i] == otherHand.cards[i] {
			continue
		}
		return hand.cards[i] < otherHand.cards[i]
	}
	return false
}

func rankHands(hands []Hand, deck map[byte]int, wildcards ...byte) {
	wildCardVals := make([]int, len(wildcards))

	for i, wildcard := range wildcards {
		wildCardVals[i] = deck[wildcard]
	}

	sort.Slice(hands, func(i, j int) bool {
		iRank, jRank := hands[i].typeRank(wildCardVals), hands[j].typeRank(wildCardVals)
		if iRank == jRank {
			return hands[i].cardsLessThan(&hands[j])
		} else {
			return iRank < jRank
		}
	})
}

func rankedProduct(hands []Hand) (total int) {
	for idx, hand := range hands {
		total += hand.bid * (idx + 1)
	}
	return
}

func PartOne() {
	handStrs := strings.Split(string(data), "\n")
	deck := initDeck(standardDeck)
	hands := initHands(&handStrs, deck)
	rankHands(hands, deck)
	println("Part One", rankedProduct(hands))
}
func PartTwo() {
	handStrs := strings.Split(string(data), "\n")
	deck := initDeck(jokerDeck)
	hands := initHands(&handStrs, deck)
	rankHands(hands, deck, 'J')
	println("Part Two", rankedProduct(hands))
}

func main() {
	PartOne()
	PartTwo()
}
