package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day11() {
	file, err := os.OpenFile("./input/input11.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	stonesMaps := make([]map[int]int, 76)
	for i := range stonesMaps {
		stonesMaps[i] = make(map[int]int) // Initialize each map
	}
	for sc.Scan() {
		line := sc.Text()

		numberStrings := strings.Fields(line)

		for _, numberString := range numberStrings {
			number, err := strconv.Atoi(numberString)
			if err != nil {
				fmt.Printf("Error converting String to Number: %v\n", err)
			}
			stonesMaps[0][number] = 1
		}
	}

	for i := 0; i < 75; i++ {
		// fmt.Println("Round: ", i)
		for stone, count := range stonesMaps[i] {
			nextStones := NextStones(stone)
			for _, nextStone := range nextStones {
				stonesMaps[i+1][nextStone] = stonesMaps[i+1][nextStone] + count
			}
		}
	}
	var count uint64 = 0
	for _, num := range stonesMaps[75] {
		count += uint64(num)
	}
	fmt.Println(count)
}

func NextStones(stone int) []int {
	// rule 1: If 0 turn to 1
	if stone == 0 {
		return []int{1}
	}

	// rule 2: If even number digits, split into two numbers in the middle
	stoneString := strconv.Itoa(stone)
	if len(stoneString)%2 == 0 {
		leftStone, err := strconv.Atoi(stoneString[:len(stoneString)/2])
		if err != nil {
			fmt.Printf("Error converting String to Number: %v\n", err)
		}
		rightStone, err := strconv.Atoi(stoneString[len(stoneString)/2:])
		if err != nil {
			fmt.Printf("Error converting String to Number: %v\n", err)
		}
		return []int{leftStone, rightStone}
	}

	// rule 3: else multiply by 2024
	return []int{stone * 2024}
}

/*

0
1
2024
20 24
2 0 2 4
4048 1 4048 8096
40 48 2024 40 48 80 96
4 0 4 8 20 24 4 0 4 8 8 0 9 6
8096 1 8096 16192 2 0 2 4 8096 1 8096 16192 16192 1 18216 12144
80 96 2024 80 96 32772608 4048 1 4048 8096 80 69 2024 80 69 32772608 32772608 2024 36869184 24579456
8 0 9 6 20 24 8 0 9 6 3277 2608 40 48 2024 40 48 80 96 8 0 6 9 20 24 8 0 6 9 3277 2608 3277 2608 20 24 3686 9184 2457 9456

*/
