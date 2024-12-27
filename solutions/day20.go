package solutions

import (
	"bufio"
	"fmt"
	"os"

	"alexlupatsiy.com/aoc24/helpers"
)

type CPUTile struct {
	value    string
	distance int
	position helpers.Coordinate
}

type CPUpossiblePath struct {
	tile     helpers.Coordinate
	distance int
}

func Day20() {
	file, err := os.OpenFile("./input/input20.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
	defer file.Close()

	rd := bufio.NewReader(file)

	cpuMap := [][]CPUTile{}
	row := []CPUTile{}
	end := helpers.Coordinate{}
	start := helpers.Coordinate{}
	y := 0
	x := 0
	for {
		r, _, err := rd.ReadRune()
		if err != nil {
			break
		}

		if r == 10 {
			cpuMap = append(cpuMap, row)
			row = []CPUTile{}
			x = 0
			y++
			continue
		}

		if string(r) == "E" {
			end.Y = y
			end.X = x
			row = append(row, CPUTile{string(r), 0, helpers.Coordinate{X: x, Y: y}})
			x++
			continue
		}
		if string(r) == "S" {
			start.Y = y
			start.X = x
			row = append(row, CPUTile{string(r), -1, helpers.Coordinate{X: x, Y: y}})
			x++
			continue
		}
		row = append(row, CPUTile{string(r), -1, helpers.Coordinate{X: x, Y: y}})
		x++
	}

	starterPossiblePath := CPUpossiblePath{end, 0}
	possiblePaths := []CPUpossiblePath{starterPossiblePath}
	for len(possiblePaths) > 0 {
		newPossiblePaths := []CPUpossiblePath{}
		for _, possiblePath := range possiblePaths {
			if possiblePath.tile == start {
				continue
			}

			directions := []helpers.Coordinate{
				{X: possiblePath.tile.X + 1, Y: possiblePath.tile.Y},
				{X: possiblePath.tile.X, Y: possiblePath.tile.Y - 1},
				{X: possiblePath.tile.X - 1, Y: possiblePath.tile.Y},
				{X: possiblePath.tile.X, Y: possiblePath.tile.Y + 1},
			}

			possibleDirections := []helpers.Coordinate{}
			for _, direction := range directions {
				canMove := moveInCPU(direction, &cpuMap)

				if canMove {
					possibleDirections = append(possibleDirections, direction)
				}
			}

			for _, possibleDirection := range possibleDirections {
				newDistance := possiblePath.distance + 1
				newPossiblePath := CPUpossiblePath{possibleDirection, newDistance}
				cpuMap[possibleDirection.Y][possibleDirection.X].distance = newDistance
				newPossiblePaths = append(newPossiblePaths, newPossiblePath)
			}
		}
		possiblePaths = newPossiblePaths
	}

	// Part 1
	//
	// count := 0
	// validTiles := []string{".", "E", "S"}
	// for y := 1; y < len(cpuMap); y++ {
	// 	for x := 1; x < len(cpuMap[0])-3; x++ {
	// 		left := cpuMap[y][x]
	// 		middle := cpuMap[y][x+1]
	// 		right := cpuMap[y][x+2]

	// 		if slices.Contains(validTiles, left.value) && middle.value == "#" && slices.Contains(validTiles, right.value) {
	// 			saved := helpers.AbsDiffInt(left.distance, right.distance) - 2
	// 			if saved >= 100 {
	// 				count++
	// 			}
	// 		}
	// 	}
	// }
	// for y := 1; y < len(cpuMap)-3; y++ {
	// 	for x := 1; x < len(cpuMap[0]); x++ {
	// 		left := cpuMap[y][x]
	// 		middle := cpuMap[y+1][x]
	// 		bottom := cpuMap[y+2][x]

	// 		if slices.Contains(validTiles, left.value) && middle.value == "#" && slices.Contains(validTiles, bottom.value) {
	// 			saved := helpers.AbsDiffInt(left.distance, bottom.distance) - 2
	// 			if saved >= 100 {
	// 				count++
	// 			}
	// 		}
	// 	}
	// }
	// fmt.Println(count)

	/* Part 2

	For Each Tile check each radius between 2 and 20 for shortcuts but its not a circle but a diamond. divX+divY={3,...,20}
	Keep a cache for pairs already checked

	*/

	maxX := len(cpuMap[0])
	maxY := len(cpuMap)
	count := 0
	for _, row := range cpuMap {
		for _, cpuTile := range row {

			if cpuTile.value == "#" {
				continue
			}

			for x := -20; x <= 20; x++ {
				for y := 0; y <= 20-helpers.AbsInt(x); y++ {
					start := cpuTile.position
					endDOWN := helpers.Coordinate{X: start.X + x, Y: start.Y - y}
					endUP := helpers.Coordinate{X: start.X + x, Y: start.Y + y}

					if CheckInsideMap(endUP.X, endUP.Y, maxX, maxY) && cpuMap[endUP.Y][endUP.X].value != "#" {
						saved := cpuTile.distance - cpuMap[endUP.Y][endUP.X].distance - (helpers.AbsInt(x) + y)
						if saved >= 100 {
							count++
						}
					}

					if CheckInsideMap(endDOWN.X, endDOWN.Y, maxX, maxY) && cpuMap[endDOWN.Y][endDOWN.X].value != "#" && y != 0 {
						saved := cpuTile.distance - cpuMap[endDOWN.Y][endDOWN.X].distance - (helpers.AbsInt(x) + y)
						if saved >= 100 {
							count++
						}
					}
				}
			}
		}
	}
	fmt.Println(count)
}

func moveInCPU(direction helpers.Coordinate, cpuMapPointer *[][]CPUTile) bool {
	cpuMap := *cpuMapPointer
	newX := direction.X
	newY := direction.Y

	if cpuMap[newY][newX].distance > -1 || cpuMap[newY][newX].value == "#" {
		return false
	}
	return true
}

func CheckInsideMap(x, y, maxX, maxY int) bool {
	if x <= 0 || y <= 0 || x >= maxX-1 || y >= maxY-1 {
		return false
	}
	return true
}
