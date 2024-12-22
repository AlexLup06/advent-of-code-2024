package solutions

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"alexlupatsiy.com/aoc24/helpers"
)

type MemoryTile struct {
	value       string
	blockedTime int
	visited     bool
}

type MemoryPath struct {
	length          int
	currentPosition helpers.Coordinate
	passedTime      int
}

const MAP_SIZE = 71

func Day18() {
	file, err := os.OpenFile("./input/input18.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
	defer file.Close()

	memoryMap := [][]MemoryTile{}
	horizontalBorder := helpers.InitSlice(MAP_SIZE+2, MemoryTile{"#", -1, false})
	for y := 0; y < MAP_SIZE+2; y++ {
		if y == 0 || y == MAP_SIZE+1 {
			memoryMap = append(memoryMap, horizontalBorder)
			continue
		}

		row := []MemoryTile{}
		for x := 0; x < MAP_SIZE+2; x++ {
			if x == 0 {
				row = append(row, MemoryTile{"#", -1, false})
				continue
			}

			if x == MAP_SIZE+1 {
				row = append(row, MemoryTile{"#", -1, false})
				memoryMap = append(memoryMap, row)
				continue
			}

			row = append(row, MemoryTile{".", -1, false})
		}
	}

	sc := bufio.NewScanner(file)
	fallingBytes := []helpers.Coordinate{}
	for sc.Scan() {
		line := sc.Text()
		regexNumber := regexp.MustCompile(`\d+`)
		numberStrings := regexNumber.FindAllString(line, -1)
		numbers := [2]int{}
		for i, numberString := range numberStrings {
			numbers[i] = helpers.STOI(numberString)
		}
		coordinate := helpers.Coordinate{X: numbers[0], Y: numbers[1]}
		fallingBytes = append(fallingBytes, coordinate)
	}

	for j := 0; j < len(fallingBytes); j++ {
		memoryMap[fallingBytes[j].Y+1][fallingBytes[j].X+1].blockedTime = j
		memoryMap[fallingBytes[j].Y+1][fallingBytes[j].X+1].value = "#"

		starterPath := MemoryPath{0, helpers.Coordinate{X: 1, Y: 1}, 0}
		possiblePaths := []MemoryPath{starterPath}

		foundShortestPath := false
		for len(possiblePaths) > 0 && !foundShortestPath {
			newPaths := []MemoryPath{}
			for _, possiblePath := range possiblePaths {
				currentPath := possiblePath

				if currentPath.currentPosition.X == MAP_SIZE && currentPath.currentPosition.Y == MAP_SIZE {
					foundShortestPath = true
					break
				}

				newPossibleTiles := []helpers.Coordinate{
					{X: currentPath.currentPosition.X + 1, Y: currentPath.currentPosition.Y},
					{X: currentPath.currentPosition.X, Y: currentPath.currentPosition.Y - 1},
					{X: currentPath.currentPosition.X - 1, Y: currentPath.currentPosition.Y},
					{X: currentPath.currentPosition.X, Y: currentPath.currentPosition.Y + 1},
				}

				newTiles := []helpers.Coordinate{}
				for _, newPossibleTile := range newPossibleTiles {
					canGo := moveInMemory(newPossibleTile, &memoryMap)

					if canGo {
						newTiles = append(newTiles, newPossibleTile)
					}
				}

				for _, newTile := range newTiles {

					newPath := MemoryPath{}
					newPath.length = currentPath.length + 1
					newPath.currentPosition.X = newTile.X
					newPath.currentPosition.Y = newTile.Y

					newPaths = append(newPaths, newPath)
				}
			}
			possiblePaths = helpers.HardCopy(newPaths)
		}

		if !foundShortestPath {
			fmt.Println("Not able to find shortest path after: ", j, ", which is falling byte: ", fallingBytes[j])
			break
		}

		for y := 0; y < MAP_SIZE; y++ {
			for x := 0; x < MAP_SIZE; x++ {
				memoryMap[y+1][x+1].visited = false
			}
		}
	}
}

func moveInMemory(newPossibleTile helpers.Coordinate, memoryMapPointer *[][]MemoryTile) bool {
	memoryMap := *memoryMapPointer
	newX := newPossibleTile.X
	newY := newPossibleTile.Y

	if memoryMap[newY][newX].visited || memoryMap[newY][newX].value == "#" {
		return false
	}

	memoryMap[newPossibleTile.Y][newPossibleTile.X].visited = true
	return true
}
