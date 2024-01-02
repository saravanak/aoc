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

var HandTypes = []string{
	"five-of-a-kind",
	"four-of-a-kind",
	"full-house",
	"three-of-a-kind",
	"two-pair",
	"one-pair",
	"high-card",
}
var LetterOrders = []string{
	"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2",
}

//TODO: Implement sorting based on the above orders

func (s *CardsList) Evaluate() {
	var cards = make([]Card, 0)

	for _, card := range s.Cards {
		cards = append(cards, *card)
		card.Categorize()
		log.Printf("%s: %s", *card.CardName, card.catetgory)
	}

	/**
	*   Sort the cards and then rank them
	 */
}

func (c *Card) Categorize() {
	c.frequencyMap = make(map[rune]int)

	for _, character := range *c.CardName {
		log.Printf("%c", character)
		c.frequencyMap[character] += 1
	}

	var counts = make([]int, 0)
	for _, v := range c.frequencyMap {
		counts = append(counts, v)
	}
	slices.SortFunc(counts, (func(a int, b int) int {
		if a > b {
			return -1
		} else if a > b {

			return 1
		}
		return 0
	}))
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

func (s *CardsList) EvaluatePart2() {

}
