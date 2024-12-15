package solutions

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"

	"alexlupatsiy.com/aoc24/helpers"
)

type Robot struct {
	x, y, dx, dy int
}

const WIDE = 101
const TALL = 103
const SECONDS = 10000

func Day14() {
	file, err := os.OpenFile("./input/input14.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Printf("Error opnening ifle %v", err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	robots := []Robot{}

	for sc.Scan() {
		line := sc.Text()
		regexNumber := regexp.MustCompile(`-?\d+`)
		numberStrings := regexNumber.FindAllString(line, -1)

		robot := Robot{}
		for i, numberString := range numberStrings {
			number := helpers.STOI(numberString)
			switch i {
			case 0:
				robot.x = number
			case 1:
				robot.y = number
			case 2:
				robot.dx = number
			case 3:
				robot.dy = number
			}
		}
		robots = append(robots, robot)
	}

	// quadrants := []int{0, 0, 0, 0}
	secondsForTree := []int{}
	bathroomMap := [TALL][WIDE]string{}
	for i := 1; i < SECONDS; i++ {
		for _, robot := range robots {
			newX := (robot.x + robot.dx*i) % WIDE
			newY := (robot.y + robot.dy*i) % TALL

			if newX < 0 {
				newX += WIDE
			}

			if newY < 0 {
				newY += TALL
			}

			// Part 1
			// switch {
			// case newX > WIDE/2 && newY < TALL/2: // first Quadrant
			// 	quadrants[0] = quadrants[0] + 1
			// case newX < WIDE/2 && newY < TALL/2: // second Quadrant
			// 	quadrants[1] = quadrants[1] + 1
			// case newX < WIDE/2 && newY > TALL/2: // third Quadrant
			// 	quadrants[2] = quadrants[2] + 1
			// case newX > WIDE/2 && newY > TALL/2: // fourth Quadrant
			// 	quadrants[3] = quadrants[3] + 1
			// }
			bathroomMap[newY][newX] = "X"
		}

		// check distribution in 20x20 fields
		const SQUARESIDE int = 20
		const SQAURES int = 25
		distributions := [SQAURES]int{}
		for k := 0; k < 100; k++ {
			for j := 0; j < 100; j++ {
				if bathroomMap[k][j] == "X" {
					distributions[(k/SQUARESIDE)*int(math.Sqrt(float64(SQAURES)))+(j/SQUARESIDE)]++
				}
			}
		}

		var average float64 = 0
		var sqaureStdev float64 = 0
		var stdev float64 = 0
		for _, distribution := range distributions {
			average += float64(distribution) / float64(SQAURES)
		}
		for _, distribution := range distributions {
			sqaureStdev += math.Pow(float64(distribution-int(average)), 2) / float64(SQAURES)
		}
		stdev = math.Sqrt(sqaureStdev)

		var n float64 = 4
		countOutside_NStdev := 0
		for _, distribution := range distributions {
			if float64(distribution) > average+n*stdev || float64(distribution) < average-n*stdev {
				countOutside_NStdev++
			}
		}
		if countOutside_NStdev >= 1 {
			secondsForTree = append(secondsForTree, i)
		}

		// reset map
		for k := 0; k < TALL; k++ {
			for j := 0; j < WIDE; j++ {
				bathroomMap[k][j] = ""
			}
		}
	}

	// print possible trees
	for _, secondForTree := range secondsForTree {
		for _, robot := range robots {
			newX := (robot.x + robot.dx*secondForTree) % WIDE
			newY := (robot.y + robot.dy*secondForTree) % TALL

			if newX < 0 {
				newX += WIDE
			}
			if newY < 0 {
				newY += TALL
			}
			bathroomMap[newY][newX] = "X"
		}
		printMap(bathroomMap)
		fmt.Println("Iteration: ", secondForTree)
		fmt.Println("______________________________________________________")

		for k := 0; k < TALL; k++ {
			for j := 0; j < WIDE; j++ {
				bathroomMap[k][j] = ""
			}
		}
	}
}

func printMap(m [103][101]string) {
	for _, row := range m {
		for _, v := range row {
			switch v {
			case "X":
				fmt.Print("X")
			default:
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
