package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func main() {
	fmt.Println("Day 1!")

	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	result := make([]int, 0)
	temp := 0

	for _, s := range strings.Split(string(input), "\n") {
		if s == "" {
			result = append(result, temp)
			temp = 0
			continue
		}

		value, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}

		temp += value
	}

	sort.Ints(result)

	sumOf3 := sum([]int{result[len(result)-1], result[len(result)-2], result[len(result)-3]})

	fmt.Println(result[len(result)-1], sumOf3)
}
