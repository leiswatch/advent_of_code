package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// a - rock, b - paper, c - scissors
// x - rock, y - paper, z - scissor
// rock - 1, paper - 2, scissor - 3
// win - 6, draw - 3, lose - 0

// part 2
// x - lose, y - draw, z - win

func main() {
	outcomes := "AXAYAZBXBYBZCXCYCZ"
	points_part_1 := "483159726"
	points_part_2 := "348159267"
	fmt.Println("Day 2!")

	input, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	defer input.Close()

	scanner := bufio.NewScanner(input)

	result_part_1 := 0
	result_part_2 := 0

	for scanner.Scan() {
		round := strings.Replace(scanner.Text(), " ", "", 1)

		temp, err := strconv.Atoi(string(points_part_1[strings.Index(outcomes, round)/2]))
		temp_2, err2 := strconv.Atoi(string(points_part_2[strings.Index(outcomes, round)/2]))

		if err != nil  || err2 != nil {
			panic(err)
		}
		result_part_1 += temp
		result_part_2 += temp_2
	}

	fmt.Println("part 1:", result_part_1)
	fmt.Println("part 2:", result_part_2)
}
