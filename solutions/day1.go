package solutions

import (
	"bufio"
	"fmt"
	"os"

	// "slices"
	"strconv"
	"strings"
	// "alexlupatsiy.com/aoc24/helpers"
)

func Day1() {
	file, err := os.OpenFile("./input/input1.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Printf("Error opnening file: %v", err)
		return
	}
	defer file.Close()
	sc := bufio.NewScanner(file)

	// dist := 0
	simScore := 0
	var left []int
	var right []int

	for sc.Scan() {
		line := sc.Text()

		values := strings.Fields(line)
		first, err1 := strconv.Atoi(values[0])
		second, err2 := strconv.Atoi(values[1])
		if err1 != nil || err2 != nil {
			fmt.Printf("Error reading number: %v, %v\n", err1, err2)
		}
		left = append(left, first)
		right = append(right, second)

	}

	// Part 1
	// slices.Sort(left)
	// slices.Sort(right)
	// for i := 0; i < len(left); i++ {
	// 	dist = dist + helpers.AbsInt(right[i]-left[i])
	// }

	// Part2
	simMap := map[int]int{}

	for _, v := range right {
		elem := simMap[v]
		simMap[v] = elem + 1
	}

	for _, v := range left {
		elem := simMap[v]
		simScore += v * elem
	}

	if err := sc.Err(); err != nil {
		fmt.Printf("scan file error: %v", err)
		return
	}

	// fmt.Println(dist)
	fmt.Println(simScore)
}
