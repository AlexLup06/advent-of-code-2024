package solutions

import (
	"bufio"
	"fmt"
	"os"

	"alexlupatsiy.com/aoc24/helpers"
)

func Day22() {
	file, err := os.OpenFile("./input/input22.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	secretNumbers := []int{}

	for sc.Scan() {
		line := sc.Text()
		secretNumber := helpers.STOI(line)
		secretNumbers = append(secretNumbers, secretNumber)
	}
	prices := helpers.InitSlice(len(secretNumbers), []int{})
	priceChanges := helpers.InitSlice(len(secretNumbers), []int{})

	for i := 0; i < 2000; i++ {
		for j, secretNumber := range secretNumbers {
			newSectretNumber := nextSecretNumber(secretNumber)

			pricet0 := secretNumber % 10
			pricet1 := newSectretNumber % 10

			priceChange := pricet1 - pricet0

			if i == 0 {
				prices[j] = append(prices[j], pricet0)
				priceChanges[j] = append(priceChanges[j], 100)
			}
			prices[j] = append(prices[j], pricet1)
			priceChanges[j] = append(priceChanges[j], priceChange)

			secretNumbers[j] = newSectretNumber
		}
	}

	possibleSequences := [][4]int{}
	for a := -9; a <= 9; a++ {
		for b := -9; b <= 9; b++ {
			for c := -9; c <= 9; c++ {
				for d := -9; d <= 9; d++ {
					if a+b < -9 || b+c < -9 || c+d < -9 || a+b > 9 || b+c > 9 || c+d > 9 || a+b+c+d > 9 || a+b+c+d < -9 || a+b+c > 9 || a+b+c < -9 {
						continue
					}
					possibleSequence := [4]int{a, b, c, d}
					possibleSequences = append(possibleSequences, possibleSequence)
				}
			}
		}
	}

	mostBananas := 0

	fmt.Println(len(possibleSequences))

	/*
		optimization: for each monkey go check all sequnces in one go and add to a map the number of bananas for that sequence
		skip sequnce of price Changes where:
			a+b < -9 || b+c < -9 || c+d < -9 || a+b > 9 || b+c > 9 || c+d > 9 || a+b+c+d > 9 || a+b+c+d < -9 || a+b+c > 9 || a+b+c < -9
		so no looping over all possible sequences
	*/
	for k, possibleSequence := range possibleSequences {
		fmt.Println(k)
		bananas := 0
		for j := 0; j < len(secretNumbers); j++ {
			priceChangesV_j := priceChanges[j]
			for i := 1; i < len(priceChanges[0])-3; i++ {
				if priceChangesV_j[i] == possibleSequence[0] && priceChangesV_j[i+1] == possibleSequence[1] && priceChangesV_j[i+2] == possibleSequence[2] && priceChangesV_j[i+3] == possibleSequence[3] {
					bananas += prices[j][i+3]
					break
				}
			}
			if bananas+9*(len(secretNumbers)-j) < mostBananas {
				break
			}
		}

		if bananas > mostBananas {
			mostBananas = bananas
		}
	}

	fmt.Println(mostBananas)

}

func nextSecretNumber(secretNumber int) int {
	newSecretNumber := secretNumber

	first := (secretNumber * 64)
	newSecretNumber = (newSecretNumber ^ first) % 16777216

	second := (newSecretNumber / 32)
	newSecretNumber = (newSecretNumber ^ second) % 16777216

	third := (newSecretNumber * 2048)
	newSecretNumber = (newSecretNumber ^ third) % 16777216

	return newSecretNumber
}
