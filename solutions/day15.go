package solutions

import (
	"bufio"
	"fmt"
	"os"
	"slices"

	"alexlupatsiy.com/aoc24/helpers"
)

func Day15() {
	file, err := os.OpenFile("./input/input15.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Printf("Error opeening file: %v\n", err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	warehouseMap := [][]string{}
	position := helpers.Coordinate{}
	movements := []string{}
	parsingMap := true
	y := 0
	for sc.Scan() {
		x := 0
		line := sc.Text()
		if line == "" {
			parsingMap = false
			continue
		}

		if parsingMap {
			row := []string{}
			for _, v := range line {
				// Part 1
				// if string(v) == "@" {
				// 	position.X = x
				// 	position.Y = y
				// }
				// row = append(row, string(v))
				// x++
				switch string(v) {
				case "#":
					row = append(row, []string{"#", "#"}...)
				case ".":
					row = append(row, []string{".", "."}...)
				case "O":
					row = append(row, []string{"[", "]"}...)
				case "@":
					position.X = x
					position.Y = y
					row = append(row, []string{"@", "."}...)
				}
				x += 2
			}
			warehouseMap = append(warehouseMap, row)
		} else {
			for _, v := range line {
				movements = append(movements, string(v))
			}
		}
		y++
	}

	// For Part 2 left and right the behaviour does not change
	for _, movement := range movements {
		switch movement {
		case ">":
			moveInWarehouse(1, 0, &warehouseMap, &position)
		case "^":
			moveInWarehouse2(-1, &warehouseMap, &position)
		case "<":
			moveInWarehouse(-1, 0, &warehouseMap, &position)
		case "v":
			moveInWarehouse2(1, &warehouseMap, &position)
		}
	}

	sum := 0
	for y, row := range warehouseMap {
		for x, v := range row {
			if v == "[" {
				sum += 100*y + x
			}
		}
	}
	fmt.Println(sum)
}

func moveInWarehouse(dx, dy int, warehouse *[][]string, p *helpers.Coordinate) {
	// Part 1
	// movingPosX := p.X
	// movingPosY := p.Y

	// for {
	// 	if (*warehouse)[movingPosY+dy][movingPosX+dx] == "." {
	// 		(*warehouse)[movingPosY+dy][movingPosX+dx] = (*warehouse)[movingPosY][movingPosX]
	// 		(*warehouse)[p.Y+dy][p.X+dx] = "@"
	// 		(*warehouse)[p.Y][p.X] = "."
	// 		p.X += dx
	// 		p.Y += dy
	// 		break
	// 	}
	// 	if (*warehouse)[movingPosY+dy][movingPosX+dx] == "#" {
	// 		break
	// 	}
	// 	movingPosX += dx
	// 	movingPosY += dy
	// }

	objectsToMove := []helpers.Coordinate{{X: p.X, Y: p.Y}}
	ableToMove := true
	for {
		lastObject := objectsToMove[len(objectsToMove)-1]

		if (*warehouse)[lastObject.Y+dy][lastObject.X+dx] == "#" {
			ableToMove = false
			break
		}

		if (*warehouse)[lastObject.Y+dy][lastObject.X+dx] == "." {
			break
		}

		objectsToMove = append(objectsToMove, helpers.Coordinate{X: lastObject.X + dx, Y: lastObject.Y + dy})
	}

	if ableToMove {
		for i := len(objectsToMove) - 1; i >= 0; i-- {
			movingObject := objectsToMove[i]
			(*warehouse)[movingObject.Y+dy][movingObject.X+dx] = (*warehouse)[movingObject.Y][movingObject.X]
		}
		(*warehouse)[p.Y][p.X] = "."
		(*warehouse)[p.Y+dy][p.X+dx] = "@"
		p.X += dx
		p.Y += dy
	}
}

/*
Part 2

# Moving up and down is no different

Idea: Start with robot and add an object if its touching it. Add this object to a new row tough and be sure to add both parts. Then next iteration checl whether
those object are thouching any other objects int he moving direction and so on.

If on of the last row is touching a border, nothing is able to move.

If all of the last rows objects don't touch anything in the moving direction then move them all
*/
func moveInWarehouse2(dy int, warehouse *[][]string, p *helpers.Coordinate) {

	// first row to move contains just the robot, second row the boxes the box the robot drectly touches,
	// the third row the first row of boxes touches and so on
	movingRows := [][]helpers.Coordinate{{{X: p.X, Y: p.Y}}}
	for {
		isAbleToMove := true
		nextMovingRow := []helpers.Coordinate{}
		everythingAbleToMove := true

		// iterate over all things in the last to move row. All rows before that touch no border and are free to move
		for _, movingPos := range movingRows[len(movingRows)-1] {
			// one of the moving objects hits a border
			if (*warehouse)[movingPos.Y+dy][movingPos.X] == "#" {
				isAbleToMove = false
				break
			}

			if (*warehouse)[movingPos.Y+dy][movingPos.X] == "[" {
				// add the object that this moving object is touching
				if !slices.Contains(nextMovingRow, helpers.Coordinate{X: movingPos.X, Y: movingPos.Y + dy}) {
					nextMovingRow = append(nextMovingRow, helpers.Coordinate{X: movingPos.X, Y: movingPos.Y + dy})
				}

				// add the right part of the box if it has not already been added
				if !slices.Contains(nextMovingRow, helpers.Coordinate{X: movingPos.X + 1, Y: movingPos.Y + dy}) {
					nextMovingRow = append(nextMovingRow, helpers.Coordinate{X: movingPos.X + 1, Y: movingPos.Y + dy})
				}
				everythingAbleToMove = false
			}

			if (*warehouse)[movingPos.Y+dy][movingPos.X] == "]" {
				// add the object that this moving object is touching
				if !slices.Contains(nextMovingRow, helpers.Coordinate{X: movingPos.X, Y: movingPos.Y + dy}) {
					nextMovingRow = append(nextMovingRow, helpers.Coordinate{X: movingPos.X, Y: movingPos.Y + dy})
				}

				// add the left part of the box if it has not already been added
				if !slices.Contains(nextMovingRow, helpers.Coordinate{X: movingPos.X - 1, Y: movingPos.Y + dy}) {
					nextMovingRow = append(nextMovingRow, helpers.Coordinate{X: movingPos.X - 1, Y: movingPos.Y + dy})
				}
				everythingAbleToMove = false
			}
		}

		// robot not able to move, becuase something touches a border
		if !isAbleToMove {
			break
		}

		movingRows = append(movingRows, nextMovingRow)
		// last row does not touch anything and we are able to move everything: push last row to next, secondToLast to last and so on
		if everythingAbleToMove {
			for i := len(movingRows) - 1; i >= 0; i-- {
				for _, movingObject := range movingRows[i] {
					(*warehouse)[movingObject.Y+dy][movingObject.X] = (*warehouse)[movingObject.Y][movingObject.X]
					(*warehouse)[movingObject.Y][movingObject.X] = "."
				}
			}
			(*warehouse)[p.Y][p.X] = "."
			(*warehouse)[p.Y+dy][p.X] = "@"
			p.Y += dy
			break
		}
	}
}
