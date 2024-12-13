package solutions

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type Position struct {
	x, y, orientation int
}

func Day6() {
	file, err := os.OpenFile("./input/input6.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Printf("Error opening File: %v\n", err)
	}
	defer file.Close()

	rd := bufio.NewReader(file)

	obstMap := [][]string{}
	row := []string{}
	position := Position{-1, -1, 3}
	start := Position{-1, -1, -1}
	for {
		r, _, err := rd.ReadRune()
		if err != nil {
			break
		}
		if r == 10 {
			obstMap = append(obstMap, row)
			row = []string{}
			continue
		}

		if string(r) == "^" {
			position.x = len(row)
			position.y = len(obstMap)
			start = Position{len(row), len(obstMap), 3}
			// guard looks up first
			row = append(row, "X")
			continue
		}

		row = append(row, string(r))
	}

	// Part 1
	for {
		newX, newY := move(position)
		if newX < 0 || newX >= len(obstMap[0]) || newY < 0 || newY >= len(obstMap) {
			break
		}
		if obstMap[newY][newX] == "#" {
			position.orientation = (position.orientation + 1) % 4
			continue
		}
		position.x = newX
		position.y = newY
		obstMap[position.y][position.x] = "X"
	}

	for _, row := range obstMap {
		fmt.Println(row)
	}
	fmt.Println()

	//
	// count1 := 0
	possibleObstaclesPlaces := []Position{}
	for y, row := range obstMap {
		for x, v := range row {
			if v == "X" {
				left := Position{x - 1, y, -1}
				if x-1 >= 0 && !slices.Contains(possibleObstaclesPlaces, left) && obstMap[left.y][left.x] != "#" {
					possibleObstaclesPlaces = append(possibleObstaclesPlaces, left)
				}
				bottom := Position{x, y + 1, -1}
				if y+1 < len(obstMap) && !slices.Contains(possibleObstaclesPlaces, bottom) && obstMap[bottom.y][bottom.x] != "#" {
					possibleObstaclesPlaces = append(possibleObstaclesPlaces, bottom)
				}
				right := Position{x + 1, y, -1}
				if x+1 < len(obstMap[0]) && !slices.Contains(possibleObstaclesPlaces, right) && obstMap[right.y][right.x] != "#" {
					possibleObstaclesPlaces = append(possibleObstaclesPlaces, right)
				}
				top := Position{x, y - 1, -1}
				if y-1 >= 0 && !slices.Contains(possibleObstaclesPlaces, top) && obstMap[top.y][top.x] != "#" {
					possibleObstaclesPlaces = append(possibleObstaclesPlaces, top)
				}
				curPos := Position{x, y, -1}
				if !slices.Contains(possibleObstaclesPlaces, curPos) && curPos.y != start.y && curPos.x != start.x && obstMap[curPos.y][curPos.x] != "#" {
					possibleObstaclesPlaces = append(possibleObstaclesPlaces, curPos)
				}
			}
		}
	}
	for y, row := range obstMap {
		for x, v := range row {
			if v == "X" {
				if x != start.x && y != start.y {
					// count1++
					obstMap[y][x] = "."
				}
			}
		}
	}
	// fmt.Println(count1)

	// Part 2

	maxX := len(obstMap[0])
	maxY := len(obstMap)

	count2 := 0
	possibleObstacles := []Position{}
	for _, obstaclePos := range possibleObstaclesPlaces {
		x := obstaclePos.x
		y := obstaclePos.y
		obstMap[y][x] = "#"

		pathTraveled := []Position{}

		for {
			// the new Position of the guard after one step
			pathTraveled = append(pathTraveled, position)
			newX, newY := move(position)

			// the guard left the map => obstacle at (x,y) is no good
			if !isInside(newX, newY, maxX, maxY) {
				break
			}

			newMapSpot := obstMap[newY][newX]

			// the guard turns when he hits an obstacle
			if newMapSpot == "#" {
				position.orientation = (position.orientation + 1) % 4
				continue
			}

			isLoop := false
			for i, field := range pathTraveled {
				if i == len(pathTraveled)-1 {
					break
				}
				if field == position && pathTraveled[i+1].x == newX && pathTraveled[i+1].y == newY {
					isLoop = true
					break
				}
			}

			if isLoop {
				possibleNewObstacle := Position{x, y, -1}
				if !slices.Contains(possibleObstacles, possibleNewObstacle) {
					count2++
					possibleObstacles = append(possibleObstacles, possibleNewObstacle)
				}
				break
			}

			position.x = newX
			position.y = newY
			obstMap[position.y][position.x] = "X"
		}
		// reset map
		for i, row := range obstMap {
			for j, v := range row {
				if v == "X" {
					obstMap[i][j] = "."
				}
			}
		}
		obstMap[y][x] = "."
		obstMap[start.y][start.x] = "X"
		position = start
	}
	fmt.Println(count2)
}

func move(position Position) (int, int) {
	switch position.orientation {
	// right
	case 0:
		return position.x + 1, position.y
	// bottom
	case 1:
		return position.x, position.y + 1
	// left
	case 2:
		return position.x - 1, position.y
	// top
	case 3:
		return position.x, position.y - 1
	}
	return -1, -1
}

func isInside(x, y, maxX, maxY int) bool {
	return x >= 0 && x < maxX && y >= 0 && y < maxY
}
