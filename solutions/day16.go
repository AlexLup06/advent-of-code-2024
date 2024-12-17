package solutions

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"slices"

	"alexlupatsiy.com/aoc24/helpers"
)

type MazeTile struct {
	x, y        int
	value       string
	lowestScore int
	pathIds     []int
	isBest      bool
}

type PathPart struct {
	x, y, orientation int
}

type Path struct {
	id         int
	points     int
	PathTiles  []PathPart
	goStraight bool
}

func Day16() {
	file, err := os.OpenFile("./input/input16.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
	defer file.Close()

	maze := [][]MazeTile{}
	row := []MazeTile{}
	start := helpers.Coordinate{}
	end := helpers.Coordinate{}
	id := rand.Int()
	y := 0
	x := 0
	rd := bufio.NewReader(file)
	for {
		r, _, err := rd.ReadRune()
		if err != nil {
			break
		}

		if r == 10 {
			maze = append(maze, row)
			row = []MazeTile{}
			y++
			x = 0
			continue
		}

		if string(r) == "S" {
			start.X = x
			start.Y = y
			row = append(row, MazeTile{x, y, string(r), 0, []int{id}, false})
			x++
			continue
		}

		if string(r) == "E" {
			end.X = x
			end.Y = y
		}

		row = append(row, MazeTile{x, y, string(r), -2, []int{}, false})
		x++
	}

	allPaths := map[int]Path{id: {id, 0, []PathPart{{start.X, start.Y, 0}}, false}}

	/*
		Idea: practiacally do Dijkstra here. I just try to implement it from my head. Don't want to google it
		=> Start at start tile. This is at the beginning the only possible Path
	*/
	keepMoving := true
	for keepMoving {
		allPathsReachedEnd := true
		for _, path := range allPaths {
			reachedEnd := moveInMaze(path, &maze, &allPaths)
			allPathsReachedEnd = allPathsReachedEnd && reachedEnd
		}
		keepMoving = !allPathsReachedEnd
	}

	pathIdsThatMadeIt := maze[end.Y][end.X].pathIds

	lowestScore := int(math.Inf(1))
	for _, pathIdThatMadeIt := range pathIdsThatMadeIt {
		if allPaths[pathIdThatMadeIt].points < lowestScore {
			lowestScore = allPaths[pathIdThatMadeIt].points
		}
	}
	bestPathIds := []int{}
	for _, pathIdThatMadeIt := range pathIdsThatMadeIt {
		if allPaths[pathIdThatMadeIt].points == lowestScore {
			bestPathIds = append(bestPathIds, pathIdThatMadeIt)
		}
	}

	for _, bestPathId := range bestPathIds {
		bestPaths := allPaths[bestPathId].PathTiles

		for _, bestPathPart := range bestPaths {
			maze[bestPathPart.y][bestPathPart.x].isBest = true
		}
	}

	count := 0
	for _, row := range maze {
		for _, v := range row {
			if v.isBest {
				count++
			}
		}
	}
	fmt.Println(count)
	fmt.Println(lowestScore)

}

/*
 */
func moveInMaze(path Path, mazePointer *[][]MazeTile, allPathsPointer *map[int]Path) bool {
	maze := (*mazePointer)
	allPaths := (*allPathsPointer)
	currentPathTile := allPaths[path.id].PathTiles[len(allPaths[path.id].PathTiles)-1]

	if maze[currentPathTile.y][currentPathTile.x].value == "E" {
		return true
	}

	possibleNextPositions := []Position{
		{X: currentPathTile.x + 1, Y: currentPathTile.y, Orientation: 0},
		{X: currentPathTile.x, Y: currentPathTile.y - 1, Orientation: 1},
		{X: currentPathTile.x - 1, Y: currentPathTile.y, Orientation: 2},
		{X: currentPathTile.x, Y: currentPathTile.y + 1, Orientation: 3},
	}

	turn := []bool{}
	newDirections := []Position{}
	allNewPoints := []int{}

	for _, possibleNextPosition := range possibleNextPositions {
		canMove, shouldTurn, newPoints := checkDirection(possibleNextPosition, path, mazePointer, allPathsPointer)

		if canMove {
			turn = append(turn, shouldTurn)
			newDirections = append(newDirections, possibleNextPosition)
			allNewPoints = append(allNewPoints, newPoints)
		}
	}

	if len(newDirections) == 0 {

		return true
	}

	// the current path should keep going and only if we are at an intersection should we create a new path
	for i, newDirection := range newDirections {
		newPathParts := []PathPart{}
		newPathParts = append(newPathParts, path.PathTiles...)

		newPathPart := PathPart{newDirection.X, newDirection.Y, newDirection.Orientation}
		if allNewPoints[i] == 1 {
			newPathPart.orientation = currentPathTile.orientation

		}
		if allNewPoints[i] == 1000 {
			newPathPart.x = currentPathTile.x
			newPathPart.y = currentPathTile.y
		}
		newPathParts = append(newPathParts, newPathPart)

		newPoints := path.points + allNewPoints[i]
		nextShouldGoStraight := !path.goStraight && turn[i]
		// one Path keeps going
		if i == 0 {
			allPaths[path.id] = Path{path.id, newPoints, newPathParts, nextShouldGoStraight}
			if !slices.Contains(maze[newPathPart.y][newPathPart.x].pathIds, path.id) {
				maze[newPathPart.y][newPathPart.x].pathIds = append(maze[newPathPart.y][newPathPart.x].pathIds, path.id)
			}
			continue
		}

		// new path from intersection
		newId := rand.Int()
		allPaths[newId] = Path{newId, newPoints, newPathParts, nextShouldGoStraight}
		maze[newPathPart.y][newPathPart.x].pathIds = append(maze[newPathPart.y][newPathPart.x].pathIds, newId)
	}

	return false
}

/*
	shouldGoStriaght         did turn       =>     new shouldGoStriaght
			0					0							0
			0					1							1
			1					0							0
			1					1							-
*/

/*
don't go:
- into a wall
- don't turn if we should go straight
- onto a tile where we have already been
- onto a tile where some other path was already faster

return:
  - bool: use this direction
  - bool: turn or not turn
  - int: points to add through the action
*/
func checkDirection(nextPosition Position, path Path, mazePointer *[][]MazeTile, allPathsPointer *map[int]Path) (bool, bool, int) {
	shouldGoStraight := path.goStraight
	newX := nextPosition.X
	newY := nextPosition.Y
	newO := nextPosition.Orientation
	currentPathPart := path.PathTiles[len(path.PathTiles)-1]

	maze := *mazePointer
	allPaths := *allPathsPointer

	// don't turn if we should go straight
	if shouldGoStraight && newO != currentPathPart.orientation {
		return false, false, -1
	}

	// into a wall
	if maze[newY][newX].value == "#" {
		return false, false, -1
	}

	// onto a tile where we have already been
	for _, pathPart := range path.PathTiles {
		if newX == pathPart.x && newY == pathPart.y {
			return false, false, -1
		}
	}

	// onto a tile where some other path was already faster
	for _, id := range maze[newY][newX].pathIds {
		if path.points >= allPaths[id].points {
			return false, false, -1
		}
	}

	pointsToAdd := 0
	if math.Abs(float64(newO-currentPathPart.orientation)) == 2 {
		pointsToAdd = 2000
	} else if newO != currentPathPart.orientation {
		pointsToAdd = 1000
	}

	// 180 degress means going back which is not allowed
	if pointsToAdd == 2000 {
		return false, false, -1
	}

	// we are turning
	if pointsToAdd == 1000 {
		return true, true, pointsToAdd
	}

	return true, false, 1
}
