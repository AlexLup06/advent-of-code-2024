package solutions

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"alexlupatsiy.com/aoc24/helpers"
)

// Part 1
// func Day2() {
// 	file, err := os.OpenFile("./input/input2.txt", os.O_RDONLY, os.ModePerm)
// 	if err != nil {
// 		fmt.Printf("Error opening file: %v\n", err)
// 	}
// 	defer file.Close()

// 	sc := bufio.NewScanner(file)

// 	safeCount := 0
// 	for sc.Scan() {
// 		line := sc.Text()

// 		stringValues := strings.Fields(line)
// 		intValues := []int{}
// 		for _, v := range stringValues {
// 			intV, err := strconv.Atoi(v)
// 			if err != nil {
// 				fmt.Printf("Error converting string to number: %v\n", err)
// 			}
// 			intValues = append(intValues, intV)
// 		}

// 		if isReportSafe(intValues) {
// 			safeCount++
// 		}
// 	}

//	if err := sc.Err(); err != nil {
// 		fmt.Printf("scan file error: %v", err)
// 	return
// }

// 	fmt.Println(safeCount)
// }

// Part 2
func Day2() {
	file, err := os.OpenFile("./input/input2.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	safeCount := 0
	for sc.Scan() {
		line := sc.Text()

		stringValues := strings.Fields(line)
		intValues := []int{}
		for _, v := range stringValues {
			intV, err := strconv.Atoi(v)
			if err != nil {
				fmt.Printf("Error converting string to number: %v\n", err)
			}
			intValues = append(intValues, intV)
		}

		var isValid = false
		for j := -1; j < len(intValues); j++ {
			if j == -1 {
				isValid = isReportSafe(intValues)
			} else {
				intValuesCopy := make([]int, len(intValues))
				copy(intValuesCopy, intValues)
				isValid = isReportSafe(slices.Delete(intValuesCopy, j, j+1))
			}
			if isValid {
				safeCount++
				break
			}
		}
	}
	if err := sc.Err(); err != nil {
		fmt.Printf("scan file error: %v", err)
		return
	}

	fmt.Println(safeCount)
}

func isReportSafe(values []int) bool {
	reportValid := true
	ascCount := 0
	for i := 0; i < len(values)-1 && reportValid; i++ {
		diff := helpers.AbsDiffInt(values[i], values[i+1])
		if diff < 1 || diff > 3 {
			reportValid = false
		}

		if values[i]-values[i+1] < 0 {
			ascCount++
		}
	}

	if !(ascCount == 0 || ascCount == len(values)-1) {
		reportValid = false
	}

	return reportValid
}
