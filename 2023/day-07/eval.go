package main

import (
	b "aoc/utils"
	// "fmt"
	"log"
	"slices"
	"strconv"

	// "strconv"
	"strings"

	"github.com/alecthomas/participle/v2"
)

var MAX_ITERS = 10000

func Parse(contents string) *CardsList {
	var basicParser = participle.MustBuild[CardsList](
		participle.Lexer(basicLexer),
		participle.UseLookahead(2),
	)

	ast, err := basicParser.ParseString("", contents)
	if err != nil {
		log.Println("Parse String error", err)
		return nil
	}
	return ast
}

var HandTypes = map[string]int{
	"five-of-a-kind":  7,
	"four-of-a-kind":  6,
	"full-house":      5,
	"three-of-a-kind": 4,
	"two-pair":        3,
	"one-pair":        2,
	"high-card":       1,
}
var LetterOrders = map[rune]int{
	'A': 13,
	'K': 12, 'Q': 11, 'J': 10, 'T': 9, '9': 8, '8': 7, '7': 6, '6': 5, '5': 4, '4': 3, '3': 2, '2': 1,
}

var LetterOrdersPart2 = map[rune]int{
	'A': 13,
	'K': 12, 'Q': 11, 'T': 9, '9': 8, '8': 7, '7': 6, '6': 5, '5': 4, '4': 3, '3': 2, '2': 1, 'J': 0,
}

func (s *CardsList) Evaluate() {
	var cards = make([]Card, 0)

	for _, card := range s.Cards {
		card.Categorize(false)
		cards = append(cards, *card)
		log.Printf("%s: %s\n", *card.CardName, card.catetgory)
	}

	/**
	*   Sort the cards and then rank them
	 */

	sortCards(cards, LetterOrders)

	sum := 0
	for i := len(cards) - 1; i >= 0; i-- {
		sum += *cards[i].Bid * (len(cards) - i)
		log.Printf("%s: %d * %d\n", *cards[i].CardName, *cards[i].Bid, len(cards)-i)
	}

	log.Printf("Solution: %d", sum)

}

func (s *CardsList) EvaluatePart2() {
	var cards = make([]Card, 0)

	for _, card := range s.Cards {
		card.Categorize(true)
		cards = append(cards, *card)
	}

	/**
	*   Sort the cards and then rank them
	 */

	sortCards(cards, LetterOrdersPart2)
	for _, card := range cards {
		log.Printf("%s: %s\n", *card.CardName, card.catetgory)
	}

	sum := 0
	for i := len(cards) - 1; i >= 0; i-- {
		sum += *cards[i].Bid * (len(cards) - i)
		log.Printf("%s: %d * %d\n", *cards[i].CardName, *cards[i].Bid, len(cards)-i)
	}

	log.Printf("Solution: %d", sum)

}

func sortCards(cards []Card, letterOrders map[rune]int) {
	log.Printf("Before sort: %+v", b.Map(cards, (func(a Card) string { return *a.CardName })))
	slices.SortFunc(cards, (func(a Card, b Card) int {
		var bName = *b.CardName
		// log.Printf("A:%+v B: %+v", a.catetgory, b.catetgory)
		if a.catetgory == b.catetgory {
			for index, aRune := range *a.CardName {
				var (
					rankA = letterOrders[aRune]
					rankB = letterOrders[rune(bName[index])]
				)
				if rankA > rankB {
					return -1
				}
				if rankA < rankB {
					return 1
				}
				//No return 0 here since we want to leave the sorting to the next alphabet down to the fifth
			}
		}

		var (
			aHand = HandTypes[a.catetgory]
			bHand = HandTypes[b.catetgory]
		)

		if aHand > bHand {
			return -1
		} else if aHand < bHand {
			return 1
		}
		return 0

	}))
	log.Printf("After Sort: %+v", b.Map(cards, (func(a Card) string { return *a.CardName })))
}

func (c *Card) Categorize(withJoker bool) {
	c.frequencyMap = make(map[rune]int)

	for _, character := range *c.CardName {
		c.frequencyMap[character] += 1
	}

	var counts = make([]int, 0)
	for _, v := range c.frequencyMap {
		counts = append(counts, v)
	}

	slices.SortFunc(counts, b.IntComparer)

	var countsAsStrings = strings.Join(b.Map(counts, (func(a int) string {
		return strconv.Itoa(a)
	})), ",")

	switch countsAsStrings {
	case "5":
		c.catetgory = "five-of-a-kind"
	case "4,1":
		c.catetgory = "four-of-a-kind"
	case "3,2":
		c.catetgory = "full-house"
	case "3,1,1":
		c.catetgory = "three-of-a-kind"
	case "2,2,1":
		c.catetgory = "two-pair"
	case "2,1,1,1":
		c.catetgory = "one-pair"
	case "1,1,1,1,1":
		c.catetgory = "high-card"
	}

}
