package solutions

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Day7() {
	file, err := os.OpenFile("./input/input7.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	calibrations := [][]int{}
	for sc.Scan() {
		line := sc.Text()

		numberStrings := strings.Fields(line)
		calibration := []int{}
		for i, v := range numberStrings {
			if i == 0 {
				callibrationNumberString := strings.Replace(v, ":", "", -1)
				calibrationNumber, err := strconv.Atoi(callibrationNumberString)
				if err != nil {
					fmt.Printf("Error converting calibration number: %v\n", err)
				}
				calibration = append(calibration, calibrationNumber)
				continue
			}
			calibrationNumber, err := strconv.Atoi(v)
			if err != nil {
				fmt.Printf("Error converting calibration number: %v\n", err)
			}
			calibration = append(calibration, calibrationNumber)

		}
		calibrations = append(calibrations, calibration)
	}

	sum1 := 0
	sum2 := 0

	for _, calibration := range calibrations {
		calibNum := calibration[0]
		matched1 := false
		matched2 := false

		// Part 1
		for operations := 0; operations < int(math.Pow(2, float64(len(calibration)-1))); operations++ {
			calibrationResult := calibration[1]
			for i, v := range calibration[2:] {

				// 0 is + and 1 is *
				operation := (operations >> i) & 1
				if operation == 0 {
					calibrationResult += v
				} else {
					calibrationResult *= v
				}
			}

			if calibrationResult == calibNum {
				matched1 = true
				break
			}
		}

		// Part 2
		operationsCount := len(calibration) - 1 - 1
		for operations := 0; operations < int(math.Pow(3, float64(len(calibration)-1))); operations++ {
			calibrationResult := calibration[1]
			operationBase3 := toBase3(operations)
			operationBase3String := strconv.Itoa(operationBase3)

			var sb strings.Builder
			// fill rest with zeros
			for i := 0; i < operationsCount-len(operationBase3String); i++ {
				sb.WriteString("0")
			}
			sb.WriteString(operationBase3String)

			fullOperationBase3 := sb.String()
			for i, v := range calibration[2:] {

				// 0 is + and 1 is * and 2 is ||
				operationString := fullOperationBase3
				currentOpertion, err := strconv.Atoi(operationString[i : i+1])
				if err != nil {
					fmt.Printf("Error converting operationBase3 to 0,1,2: %v\n", err)
				}
				switch currentOpertion {
				case 0:
					calibrationResult += v
				case 1:
					calibrationResult *= v
				case 2:
					var sb strings.Builder
					sb.WriteString(strconv.Itoa(calibrationResult))
					sb.WriteString(strconv.Itoa(v))
					calibrationResult, err = strconv.Atoi(sb.String())

					if err != nil {
						fmt.Printf("Error concatinating: %v\n", err)
					}
				}
			}

			if calibrationResult == calibNum {
				matched2 = true
				break
			}
		}

		if matched1 {
			sum1 += calibNum
		}
		if matched2 {
			sum2 += calibNum
		}
	}
	fmt.Println(sum1)
	fmt.Println(sum2)
}

// 10 == 101
func toBase3(x int) int {
	if x == 0 {
		return 0
	}
	digits := math.Floor(logBase3(float64(x)))
	base3number := 0
	for digits >= 0 { // x=10 digits=2 ; x=1 digits=1 ; x=1 digits=0
		y1 := int(math.Pow(3, digits))                  // y1=3^2=9 ; y1=3^1=3 ; y1=3^0=1
		y2 := int(math.Floor(float64(x) / float64(y1))) // y2=10/9=1,...=1 ;  y2=1/3=0.3333=0 ; y2= 1/1=1
		base3number += int(math.Pow(10, digits)) * y2
		x -= y1 * y2
		digits--
	}
	return base3number
}

func logBase3(x float64) float64 {
	return math.Log(x) / math.Log(3)
}
