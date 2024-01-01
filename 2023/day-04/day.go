package main

/**
* Part 1: 21959
 */
import (
	b "aoc/utils"
	"fmt"
	"github.com/bitfield/script"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var line_number = 0

type cardDraw struct {
	winningCards []int
	ourCards     []int
}

type game struct {
	cardList []cardDraw
}

var currentGame = game{cardList: make([]cardDraw, 0)}

func lineParser(line string) string {

	var winningCards = make([]int, 1)
	var ourCards = make([]int, 1)

	card := strings.Split(line, ":")
	cardNumber, _ := strconv.Atoi(card[0])

	cardDraws := strings.Split(card[1], "|")
	line_parser := regexp.MustCompile("(\\d*)")
	line_number++

	greaterThanZeroPredicate := func(num int) bool { return num > 0 }
	winningCards = b.Filter(b.Map(line_parser.FindAllStringSubmatch(cardDraws[0], -1), (func(cardValue []string) int {
		cardNumber, _ = strconv.Atoi(cardValue[0])
		return cardNumber
	})), (greaterThanZeroPredicate))

	ourCards = b.Filter(b.Map(line_parser.FindAllStringSubmatch(cardDraws[1], -1), (func(cardValue []string) int {
		cardNumber, _ = strconv.Atoi(cardValue[0])
		return cardNumber
	})), greaterThanZeroPredicate)

	sort.Slice(winningCards, func(x int, y int) bool { return winningCards[x] < winningCards[y] })
	sort.Slice(ourCards, func(x int, y int) bool { return ourCards[x] < ourCards[y] })

	currentGame.cardList = append(currentGame.cardList, cardDraw{winningCards, ourCards})

	return line
}

func part01() {

	var sum = 0
	for _, currentCardGame := range currentGame.cardList {
		var score = 0
		for _, ourCard := range currentCardGame.ourCards {
			log.Println("Searching for ", ourCard)
			foundIndex := sort.Search(len(currentCardGame.winningCards), (func(index int) bool {
				if index == (len(currentCardGame.winningCards)) {
					return false
				}
				// log.Println(index, currentCardGame.winningCards[index], ourCard)
				return currentCardGame.winningCards[index] >= ourCard
			}))
			if foundIndex < len(currentCardGame.winningCards) && currentCardGame.winningCards[foundIndex] == ourCard {
				log.Printf("Found %d", ourCard)
				if score == 0 {
					score = 1
				} else {
					score = score * 2
				}
			}
		}
		// log.Println("Score is", score)
		sum += score
	}
	fmt.Printf("Part 01 Soution is %d", sum)
}

func part02() {
	var sum = 0
	var cardCounts = make([]int, len(currentGame.cardList))
	for index := range cardCounts {
		cardCounts[index] = 1
	}
	for cardIndex, currentCardGame := range currentGame.cardList {
		var score = 0
		for _, ourCard := range currentCardGame.ourCards {
			// log.Println("Searching for ", ourCard)
			foundIndex := sort.Search(len(currentCardGame.winningCards), (func(index int) bool {
				if index == (len(currentCardGame.winningCards)) {
					return false
				}
				// log.Println(index, currentCardGame.winningCards[index], ourCard)
				return currentCardGame.winningCards[index] >= ourCard
			}))
			if foundIndex < len(currentCardGame.winningCards) && currentCardGame.winningCards[foundIndex] == ourCard {
				// log.Printf("Found %d", ourCard)
				score += 1
			}
		}
		log.Printf("Found %d for cardIndex %d", score, cardIndex)
		for i := cardIndex + 1; i < cardIndex+score+1; i++ {
			cardCounts[i] += cardCounts[cardIndex]
		}
	}

	for index := range cardCounts {
		sum += cardCounts[index]
	}
	fmt.Printf("Part 02 Soution is %d", sum)
}

func main() {
	fileName := "./data/04/full.txt"
	// fileName := "./data/04/example.txt"
	script.File(fileName).FilterLine(lineParser).Wait()
	if os.Args[1] == "part2" {
		fmt.Println("running day04/part 02!!!")
		part02()
	} else {
		fmt.Println("running day04/part 01!!!")
		part01()
	}

}
