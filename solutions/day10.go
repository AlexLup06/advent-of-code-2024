package solutions

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"

	"alexlupatsiy.com/aoc24/helpers"
)

type TrailPosition struct {
	x, y, value int
}

func Day10() {
	file, err := os.OpenFile("./input/input10.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
	}
	defer file.Close()

	starters := []helpers.Coordinate{}
	topoMap := [][]int{}

	rd := bufio.NewReader(file)
	row := []int{}
	row = append(row, -1)
	for {
		r, _, err := rd.ReadRune()

		if err != nil {
			break
		}

		if r == 10 {
			row = append(row, -1)
			topoMap = append(topoMap, row)
			row = []int{}
			row = append(row, -1)
			continue
		}

		number, err := strconv.Atoi(string(r))
		if err != nil {
			fmt.Printf("Error converting number: %v", err)
		}
		if string(r) == "0" {
			starters = append(starters, helpers.Coordinate{X: len(row), Y: len(topoMap) + 1})
		}
		row = append(row, number)
	}
	topoMap = append(topoMap, helpers.InitSlice(len(topoMap[0]), -1))
	topoMap = slices.Insert(topoMap, 0, helpers.InitSlice(len(topoMap[0]), -1))

	sum := 0
	for _, starter := range starters {
		score := 0
		currentTrailPosition := TrailPosition{starter.X, starter.Y, 0}
		possibleTrail := []TrailPosition{currentTrailPosition}
		// Part 1
		// checkedTopPositions := []TrailPosition{}

		for len(possibleTrail) > 0 {
			trailPosition := possibleTrail[0]
			possibleTrail = slices.Delete(possibleTrail, 0, 1)
			x := trailPosition.x
			y := trailPosition.y

			// check top,left,bottom,right
			nextTrails := []helpers.Coordinate{{X: x, Y: y - 1}, {X: x - 1, Y: y}, {X: x, Y: y + 1}, {X: x + 1, Y: y}}
			for _, nextTrail := range nextTrails {
				nextTrailPosition, ok := checkTrail(trailPosition.value, topoMap[nextTrail.Y][nextTrail.X], nextTrail.X, nextTrail.Y)

				// Part 1
				// if ok && nextTrailPosition.value == 9 && !slices.Contains(checkedTopPositions, nextTrailPosition) {

				// Part 2 -- Just the if Condition
				if ok && nextTrailPosition.value == 9 {
					// Part 1
					// checkedTopPositions = append(checkedTopPositions, nextTrailPosition)
					score++
				}
				if ok && nextTrailPosition.value != 9 {
					possibleTrail = append(possibleTrail, nextTrailPosition)
				}
			}
		}
		sum += score
	}
	fmt.Println(sum)
}

func checkTrail(trailPositionValue, nextValue, nextX, nextY int) (TrailPosition, bool) {
	if trailPositionValue+1 == nextValue {
		return TrailPosition{nextX, nextY, nextValue}, true
	}
	return TrailPosition{-1, -1, -1}, false
}
