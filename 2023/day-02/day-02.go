package main

import (
	// "4d63.com/strrev"
	b "aoc/utils"
	"fmt"
	"regexp"

	// "regexp"
	// "maps"
	"strconv"
	"strings"
)

type draw struct {
	red   int
	blue  int
	green int
}

type game struct {
	draws   []draw
	game_id int
}

func part1() {

	// input_file := "./data/02/example.txt"
	input_file := "./data/02/day-02.txt"
	fileLines := b.ReadFileAsArray(input_file)

	MAX_RED := 12
	MAX_BLUE := 14
	MAX_GREEN := 13
	sum := 0

	part_2_sum := 0
	for _, calibWord := range fileLines {
		fmt.Println(calibWord)
		game := splitDraws(calibWord)
		fmt.Println(game)
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
			// if !is_valid_game {
			// 	break
			// }
		}
		if is_valid_game {
			fmt.Printf("We have %d as valid game\n", game.game_id)
			sum = sum + game.game_id
		}
		power_for_game := (max_greens * max_blues * max_reds)

		fmt.Printf("Power: %d; Red: %d, Green: %d, Blue: %d\n", power_for_game, max_reds, max_greens, max_blues)
		part_2_sum = part_2_sum + power_for_game
	}

	fmt.Println(sum)
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
			// fmt.Printf("%#v\n", gameColors)
			// fmt.Printf("%#v\n", game_draw_regexp.SubexpNames())

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
	part1()
}
