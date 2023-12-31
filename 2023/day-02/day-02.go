package main

import (
	b "aoc/utils"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type programArgs struct {
	should_break bool // True for part1 and False for part 2
}

type draw struct {
	red   int
	blue  int
	green int
}

type game struct {
	draws   []draw
	game_id int
}

func part1(programOptions programArgs) {

	// input_file := "./data/02/example.txt"
	input_file := "./data/02/day-02.txt"
	fileLines := b.ReadFileAsArray(input_file)

	MAX_RED := 12
	MAX_BLUE := 14
	MAX_GREEN := 13
	sum := 0

	part_2_sum := 0
	for _, calibWord := range fileLines {
		log.Println(calibWord)
		game := splitDraws(calibWord)
		log.Println(game)
		is_valid_game := true
		max_greens := 0
		max_blues := 0
		max_reds := 0
		for _, currentDraw := range game.draws {
			is_valid_game = currentDraw.green <= MAX_GREEN &&
				currentDraw.red <= MAX_RED &&
				currentDraw.blue <= MAX_BLUE

			if currentDraw.green > 0 && currentDraw.green > max_greens {
				max_greens = currentDraw.green
			}
			if currentDraw.blue > 0 && currentDraw.blue > max_blues {
				max_blues = currentDraw.blue
			}
			if currentDraw.red > 0 && currentDraw.red > max_reds {
				max_reds = currentDraw.red
			}
			if programOptions.should_break {
				if !is_valid_game {
					break
				}
			}
		}
		if is_valid_game {
			log.Printf("We have %d as valid game\n", game.game_id)
			sum = sum + game.game_id
		}
		power_for_game := (max_greens * max_blues * max_reds)

		log.Printf("Power: %d; Red: %d, Green: %d, Blue: %d\n", power_for_game, max_reds, max_greens, max_blues)
		part_2_sum = part_2_sum + power_for_game
	}

	fmt.Printf("Part 1: %d\n", sum)
	fmt.Printf("Part 2: %d", part_2_sum)

}

func splitDraws(game_desc string) game {

	game_draw_regexp := regexp.MustCompile("(?P<Count>\\d*) (?P<Color>blue|red|green)")
	game_index_list := strings.Split(game_desc, ":")
	game_id_regexp := regexp.MustCompile("Game (?P<Id>\\d*)")
	game_id_parts := game_id_regexp.FindStringSubmatch(game_index_list[0])
	game_id, _ := strconv.Atoi(game_id_parts[1])
	games := strings.Split(game_index_list[1], ";")

	var gameDrawsCollection []draw

	for _, gameDraws := range games {
		gameDrawsParts := strings.Split(gameDraws, ",")
		var currentDraw = draw{red: 0, blue: 0, green: 0}
		for _, drawString := range gameDrawsParts {
			gameColors := game_draw_regexp.FindStringSubmatch(drawString)

			count, _ := strconv.Atoi(gameColors[1])

			switch gameColors[2] {
			case "blue":
				currentDraw.blue = count
			case "red":
				currentDraw.red = count
			case "green":
				currentDraw.green = count
			}
		}
		gameDrawsCollection = append(gameDrawsCollection, currentDraw)
	}
	return game{draws: gameDrawsCollection, game_id: game_id}
}

func main() {
	if os.Args[1] == "part2" {
		fmt.Println("running day02/part 02!!!")
		part1(programArgs{should_break: false})
	} else {
		fmt.Println("running day02/part 01!!!")
		part1(programArgs{should_break: true})
	}
}
