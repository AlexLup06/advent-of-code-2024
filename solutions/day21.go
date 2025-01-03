package solutions

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"slices"

// 	"alexlupatsiy.com/aoc24/helpers"
// )

// type NumericKeyboard struct {
// 	currentNumKeyHover string
// }

// type DirectionalKeypadInterface interface {
// 	getLengthShortestSeq(code string) int
// }

// type DirectionalKeypad struct {
// 	numKey                NumericKeyboard
// 	currentDirectionHover string
// }

// func newDirectionalKeypad() DirectionalKeypadInterface {
// 	return &DirectionalKeypad{numKey: NumericKeyboard{"A"}, currentDirectionHover: "A"}
// }

// func (dk *DirectionalKeypad) moveTo(nextNumKey string) int {
// 	movements := dk.getMovements(nextNumKey)
// 	currentDirection := dk.currentDirectionHover
// 	currentNumKey := dk.numKey.currentNumKeyHover
// 	divX := movements.X
// 	divY := movements.Y

// 	if slices.Contains([]string{"1", "4", "7"}, currentNumKey) && slices.Contains([]string{"0", "A"}, nextNumKey) {
// 		return dk.move(">", "v", nextNumKey, divX, divY)
// 	}

// 	if slices.Contains([]string{"0", "A"}, currentNumKey) && slices.Contains([]string{"1", "4", "7"}, nextNumKey) {
// 		return dk.move("^", "<", nextNumKey, divX, divY)
// 	}

// 	if divX == 0 {
// 		movementsCost := 0
// 		direction := ""
// 		if divY > 0 {
// 			direction = "v"
// 		} else {
// 			direction = "^"
// 		}
// 		movementsCost += dk.moveNumKeyRobotCost(direction, true)
// 		movementsCost += helpers.AbsInt(divY) - 1          // we only need tp press A to further go in that direction
// 		movementsCost += dk.moveNumKeyRobotCost("A", true) // we need to go to A
// 		dk.numKey.currentNumKeyHover = nextNumKey
// 		return movementsCost
// 	}

// 	if divY == 0 {
// 		movementsCost := 0
// 		direction := ""
// 		if divX > 0 {
// 			direction = ">"
// 		} else {
// 			direction = "<"
// 		}
// 		movementsCost += dk.moveNumKeyRobotCost(direction, true)
// 		movementsCost += helpers.AbsInt(divX) - 1          // we only need tp press A to further go in that direction
// 		movementsCost += dk.moveNumKeyRobotCost("A", true) // we need to go to A
// 		dk.numKey.currentNumKeyHover = nextNumKey
// 		return movementsCost
// 	}

// 	switch {
// 	case divX > 0 && divY > 0:
// 		firstRightCost := dk.move(">", "v", nextNumKey, divX, divY)
// 		dk.currentDirectionHover = currentDirection
// 		dk.numKey.currentNumKeyHover = currentNumKey

// 		firstDownCost := dk.move("v", ">", nextNumKey, divX, divY)
// 		dk.currentDirectionHover = currentDirection
// 		dk.numKey.currentNumKeyHover = currentNumKey

// 		if firstRightCost < firstDownCost {
// 			return dk.move(">", "v", nextNumKey, divX, divY)
// 		} else {
// 			return dk.move("v", ">", nextNumKey, divX, divY)
// 		}
// 	case divX < 0 && divY > 0:
// 		firstLeftCost := dk.move("<", "v", nextNumKey, divX, divY)
// 		dk.currentDirectionHover = currentDirection
// 		dk.numKey.currentNumKeyHover = currentNumKey

// 		firstDownCost := dk.move("v", "<", nextNumKey, divX, divY)
// 		dk.currentDirectionHover = currentDirection
// 		dk.numKey.currentNumKeyHover = currentNumKey

// 		if firstLeftCost < firstDownCost {
// 			return dk.move("<", "^", nextNumKey, divX, divY)
// 		} else {
// 			return dk.move("^", "<", nextNumKey, divX, divY)
// 		}
// 	case divX > 0 && divY < 0:
// 		firstRightCost := dk.move(">", "^", nextNumKey, divX, divY)
// 		dk.currentDirectionHover = currentDirection
// 		dk.numKey.currentNumKeyHover = currentNumKey

// 		firstUpCost := dk.move("^", ">", nextNumKey, divX, divY)
// 		dk.currentDirectionHover = currentDirection
// 		dk.numKey.currentNumKeyHover = currentNumKey

// 		if firstRightCost < firstUpCost {
// 			return dk.move(">", "^", nextNumKey, divX, divY)
// 		} else {
// 			return dk.move("^", ">", nextNumKey, divX, divY)
// 		}
// 	case divX < 0 && divY < 0:
// 		firstLeftCost := dk.move("<", "^", nextNumKey, divX, divY)
// 		dk.currentDirectionHover = currentDirection
// 		dk.numKey.currentNumKeyHover = currentNumKey

// 		firstUpCost := dk.move("^", "<", nextNumKey, divX, divY)
// 		dk.currentDirectionHover = currentDirection
// 		dk.numKey.currentNumKeyHover = currentNumKey

// 		if firstLeftCost < firstUpCost {
// 			return dk.move("<", "^", nextNumKey, divX, divY)
// 		} else {
// 			return dk.move("^", "<", nextNumKey, divX, divY)
// 		}
// 	}
// 	dk.numKey.currentNumKeyHover = nextNumKey
// 	return -1

// }

// func (dk *DirectionalKeypad) move(firstDirection, secondDirection, nextNumKey string, divX, divY int) int {
// 	movementsCost := 0

// 	movementsCost += dk.moveNumKeyRobotCost(firstDirection, true)
// 	movementsCost += dk.moveNumKeyRobotCost(secondDirection, true)

// 	if divX != 0 {
// 		movementsCost += helpers.AbsInt(divX) - 1 // we only need tp press A to further go in that direction
// 	}
// 	if divY != 0 {
// 		movementsCost += helpers.AbsInt(divY) - 1 // we only need tp press A to further go in that direction
// 	}
// 	movementsCost += dk.moveNumKeyRobotCost("A", true) // we need to go to A
// 	dk.numKey.currentNumKeyHover = nextNumKey
// 	return movementsCost
// }

// func (dk *DirectionalKeypad) getMovements(nextPositionString string) helpers.Coordinate {
// 	currentPosition := dk.stringToCoordinates(dk.numKey.currentNumKeyHover)
// 	nextPosition := dk.stringToCoordinates(nextPositionString)
// 	return helpers.Coordinate{X: nextPosition.X - currentPosition.X, Y: nextPosition.Y - currentPosition.Y}
// }

// func (dk *DirectionalKeypad) stringToCoordinates(positionString string) helpers.Coordinate {
// 	switch positionString {
// 	case "1":
// 		return helpers.Coordinate{X: 0, Y: 2}
// 	case "2":
// 		return helpers.Coordinate{X: 1, Y: 2}
// 	case "3":
// 		return helpers.Coordinate{X: 2, Y: 2}
// 	case "4":
// 		return helpers.Coordinate{X: 0, Y: 1}
// 	case "5":
// 		return helpers.Coordinate{X: 1, Y: 1}
// 	case "6":
// 		return helpers.Coordinate{X: 2, Y: 1}
// 	case "7":
// 		return helpers.Coordinate{X: 0, Y: 0}
// 	case "8":
// 		return helpers.Coordinate{X: 1, Y: 0}
// 	case "9":
// 		return helpers.Coordinate{X: 2, Y: 0}
// 	case "0":
// 		return helpers.Coordinate{X: 1, Y: 3}
// 	case "A":
// 		return helpers.Coordinate{X: 2, Y: 3}
// 	}
// 	return helpers.Coordinate{X: -1, Y: -1}
// }

// func (dk *DirectionalKeypad) getLengthShortestSeq(code string) int {
// 	length := 0
// 	for _, char := range code {
// 		length += dk.moveTo(string(char))
// 	}
// 	return length
// }

// func Day21() {
// 	file, err := os.OpenFile("./input/input21.txt", os.O_RDONLY, os.ModePerm)
// 	if err != nil {
// 		fmt.Printf("Errro opening file: %v\n", err)
// 	}
// 	defer file.Close()

// 	codeStrings := []string{}
// 	sc := bufio.NewScanner(file)
// 	for sc.Scan() {
// 		line := sc.Text()
// 		codeStrings = append(codeStrings, line)
// 	}

// 	keypad := newDirectionalKeypad()
// 	sum := 0
// 	for _, codeString := range codeStrings {
// 		shortestLength := keypad.getLengthShortestSeq(codeString)

// 		fmt.Println("shortest length: ", shortestLength, ", for code: ", codeString)
// 		// fmt.Println()

// 		numericPart := helpers.STOI(codeString[:3])
// 		complexity := shortestLength * numericPart
// 		sum += complexity
// 	}
// 	fmt.Println(sum)
// }
