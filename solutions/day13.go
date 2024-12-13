package solutions

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"

	"alexlupatsiy.com/aoc24/helpers"
)

type Button struct {
	dx, dy int
}

type ClawMachine struct {
	buttonA, buttonB Button
	prizeX, prizeY   int
}

const costPushA = 3
const costPushB = 1

func Day13() {
	file, err := os.OpenFile("./input/input13.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Printf("Error opening file:%v\n", err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	findNumbers := regexp.MustCompile(`\d+`)

	i := 0
	clawMachines := []ClawMachine{}
	buttonA := Button{}
	buttonB := Button{}
	for sc.Scan() {
		line := sc.Text()

		numberStrings := findNumbers.FindAllString(line, -1)

		switch i % 4 {
		case 0:
			// Button A
			dx := helpers.STOI(numberStrings[0])
			dy := helpers.STOI(numberStrings[1])

			buttonA = Button{dx, dy}
		case 1:
			// Button B
			dx := helpers.STOI(numberStrings[0])
			dy := helpers.STOI(numberStrings[1])

			buttonB = Button{dx, dy}
		case 2:
			// Claw Maschine
			// prizeX := helpers.STOI(numberStrings[0])
			// prizeY := helpers.STOI(numberStrings[1])
			prizeX := helpers.STOI(numberStrings[0]) + 10000000000000
			prizeY := helpers.STOI(numberStrings[1]) + 10000000000000
			clawMachine := ClawMachine{buttonA, buttonB, prizeX, prizeY}
			clawMachines = append(clawMachines, clawMachine)
		}
		i++
	}

	totalCost := 0
	for _, clawMachine := range clawMachines {
		A := clawMachine.buttonA.dx
		B := clawMachine.buttonB.dx
		C := clawMachine.prizeX
		D := clawMachine.buttonA.dy
		E := clawMachine.buttonB.dy
		F := clawMachine.prizeY

		denum := (A*E - B*D)

		if denum == 0 {
			// the lines are parallel => no solution
			continue
		}

		if float64(C/B) == float64(F/E) && float64(A/B) == float64(D/E) {
			// the lines are the same line => infinetly many solutions => pick the cheapest one
			continue
		}

		// fmt.Println("Exactly one solution")
		// the lines meet at one point => exactly one solution
		pressA := float64(C*E-B*F) / float64(denum) // x
		pressB := float64(A*F-C*D) / float64(denum) // y

		if math.Mod(pressA, 1) != 0 || math.Mod(pressB, 1) != 0 {
			// pressA or PressB not integers
			continue
		}

		// Part 1
		// if pressA <= 100 && pressA >= 0 && pressB <= 100 && pressB >= 0 {

		// Part 2
		if pressA >= 0 && pressB >= 0 {
			totalCost += int(pressA)*costPushA + int(pressB)*costPushB
		}
	}
	fmt.Println(totalCost)
}

/*

A*x+B*y=C
D*x+E*y=F

y=-A/B * x + C/B
y=-D/E * x + F/E

(A b)=(A B | C)
      (D E | F)

x=det(A1)/det(A)=(C*E-B*F)/(A*E-B*D)
y=det(A2)/det(A)=(A*F-C*D)/(A*E-B*D)

*/
