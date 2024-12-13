package solutions

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day3() {
	file, err := os.OpenFile("./input/input3.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
	defer file.Close()

	rd := bufio.NewReader(file)
	sum := 0
	valid := true
	for {
		r, _, err := rd.ReadLine()
		if err != nil {
			break
		}
		// regexpMult := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
		// mults := regexpMult.FindAllString(string(r), -1)

		regexpInstructions := regexp.MustCompile(`(mul\([0-9]{1,3},[0-9]{1,3}\))|(do\(\))|(don't\(\))`)
		instructions := regexpInstructions.FindAllString(string(r), -1)
		validMults := []string{}
		for _, instr := range instructions {
			if instr == "do()" {
				valid = true
				continue
			}
			if instr == "don't()" {
				valid = false
				continue
			}
			if valid {
				validMults = append(validMults, instr)
			}
		}

		// for _, mult := range mults {
		for _, mult := range validMults {
			regexpNotNum := regexp.MustCompile(`[^\d]`)
			numberStringsWithWhitespaces := regexpNotNum.ReplaceAll([]byte(mult), []byte(" "))
			numberStrings := strings.Fields(string(numberStringsWithWhitespaces))

			val1, err1 := strconv.Atoi(numberStrings[0])
			val2, err2 := strconv.Atoi(numberStrings[1])

			if err1 != nil || err2 != nil {
				fmt.Printf("Error converting a number: %v, %v\n", err1, err2)
			}

			sum += val1 * val2
		}
	}
	fmt.Println(sum)
}
