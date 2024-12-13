package solutions

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Day5() {
	file, err := os.OpenFile("./input/input5.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	pageRules := map[int][]int{}
	updates := [][]int{}

	scanPageRules := true
	for sc.Scan() {
		line := sc.Text()

		if line == "" {
			scanPageRules = false
			continue
		}

		if scanPageRules {
			ruleNumberStrings := strings.Split(line, "|")
			leftSide, err := strconv.Atoi(ruleNumberStrings[0])
			if err != nil {
				fmt.Printf("Error Converting string to int: %v\n", err)
			}
			rightSide, err := strconv.Atoi(ruleNumberStrings[1])
			if err != nil {
				fmt.Printf("Error Converting string to int: %v\n", err)
			}

			values := pageRules[leftSide]
			values = append(values, rightSide)
			pageRules[leftSide] = values
		}

		if !scanPageRules {
			updateNumberStrings := strings.Split(line, ",")
			update := []int{}
			for _, v := range updateNumberStrings {
				updateNumber, err := strconv.Atoi(v)
				if err != nil {
					fmt.Printf("Error Converting string to int: %v\n", err)
				}
				update = append(update, updateNumber)
			}
			updates = append(updates, update)
		}
	}

	sum := 0
	invalidUpdates := [][]int{}
	for _, update := range updates {
		slices.Reverse(update)
		updateValid := true
		// fmt.Printf("Checking Update: %v\n", update)
		for i, v := range update {
			rulesForV := pageRules[v]

			for j := i + 1; j < len(update); j++ {
				checker := update[j]
				// fmt.Printf("Checking if %v is braking rule of %v\n", checker, v)
				if slices.Contains(rulesForV, checker) {
					// fmt.Printf("%v broke rule of %v\n", checker, v)
					updateValid = false
					break
				}
			}

			if !updateValid {
				invalidUpdates = append(invalidUpdates, update)
				break
			}
		}

		if updateValid {
			// fmt.Printf("No Rule Broken. Middle number is %v.\n", update[len(update)/2])
			sum += update[len(update)/2]
		}
		// fmt.Println("")
	}

	sum2 := 0
	for _, invalidUpdate := range invalidUpdates {
		for _, invalidNumber := range invalidUpdate {
			countIsOnRightSide := 0

			for _, checker := range invalidUpdate {
				if checker == invalidNumber {
					continue
				}

				if slices.Contains(pageRules[checker], invalidNumber) {
					countIsOnRightSide++
				}
			}

			if countIsOnRightSide == len(invalidUpdate)/2 {
				sum2 += invalidNumber
				// fmt.Println(invalidNumber)
			}

		}
	}

	fmt.Println(sum)
	fmt.Println(sum2)
}
