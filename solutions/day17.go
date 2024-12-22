package solutions

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"alexlupatsiy.com/aoc24/helpers"
)

type Computer struct {
	registerA          int
	registerB          int
	registerC          int
	instructionPointer int
	program            []int
	output             []int
}

func (c *Computer) next() bool {
	jumped := false
	// fmt.Println(c.program[c.instructionPointer])
	// fmt.Println("--")
	switch c.program[c.instructionPointer] {
	case 0:
		operand := c.getOperand(true)
		c.registerA = c.registerA >> operand // division x/(2^n)= x >> n
	case 1:
		operand := c.getOperand(false)
		c.registerB = operand ^ c.registerB
	case 2:
		operand := c.getOperand(true)
		c.registerB = operand & 7 // n % 2^i = n & (2^i - 1)
	case 3:
		jumped = c.jnz()
	case 4:
		c.registerB = c.registerB ^ c.registerC
	case 5:
		operand := c.getOperand(true)
		c.output = append(c.output, operand&7)
	case 7:
		operand := c.getOperand(true)
		c.registerC = c.registerA >> operand // division x/(2^n)= x >> n
	}

	if !jumped {
		c.instructionPointer += 2
	}

	// program halts
	if c.instructionPointer >= len(c.program) {
		return false
	}

	// program keeps going
	return true
}

func (c *Computer) jnz() bool {
	operand := c.getOperand(false)
	regAValue := c.registerA
	if regAValue != 0 {
		c.instructionPointer = operand
		return true
	}
	return false
}

func (c *Computer) getOperand(isCombo bool) int {
	if isCombo {
		switch c.program[c.instructionPointer+1] {
		case 0:
			return 0
		case 1:
			return 1
		case 2:
			return 2
		case 3:
			return 3
		case 4:
			return c.registerA
		case 5:
			return c.registerB
		case 6:
			return c.registerC
		}
	}

	return c.program[c.instructionPointer+1]
}

type PossiblePath struct {
	currentK int
	regA     []int
}

func Day17() {
	file, err := os.OpenFile("./input/input17.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	registerValues := []int{}
	program := []int{}
	readingRegisters := true
	for sc.Scan() {
		line := sc.Text()

		if line == "" {
			readingRegisters = false
			continue
		}

		regNumbers := regexp.MustCompile(`\d+`)
		if readingRegisters {
			stringNumber := regNumbers.FindAllString(line, -1)
			number := helpers.STOI(stringNumber[0])
			registerValues = append(registerValues, number)
		} else {
			stringNumbers := regNumbers.FindAllString(line, -1)

			for _, stringNumber := range stringNumbers {
				number := helpers.STOI(stringNumber)
				program = append(program, number)
			}
		}
	}

	if err := sc.Err(); err != nil {
		fmt.Printf("Error scanning file: %v\n", err)
	}

	computer := Computer{
		registerA:          registerValues[0],
		registerB:          registerValues[1],
		registerC:          registerValues[2],
		instructionPointer: 0,
		program:            program,
		output:             []int{},
	}
	for computer.next() {
	}

	var sb strings.Builder

	for i, output := range computer.output {
		if i == len(computer.output)-1 {
			sb.WriteString(strconv.Itoa(output))
			continue
		}
		sb.WriteString(strconv.Itoa(output))
		sb.WriteString(",")
	}

	fmt.Println(sb.String())
	fmt.Println("______")

	output := []int{2, 4, 1, 3, 7, 5, 4, 7, 0, 3, 1, 5, 5, 5, 3, 0}
	startRegA := helpers.InitSlice(48, -1)
	startRegA[0] = 1
	startRegA[1] = 1
	startRegA[2] = 0

	// TODO: implement backtracking
	possibilities := helpers.Stack[PossiblePath]()
	possibilities.Push(PossiblePath{len(output) - 2, startRegA})

	finishedPossiblities := []PossiblePath{}

	for possibilities.Length() > 0 {
		possibility := possibilities.Pop()
		k := possibility.currentK
		regA := helpers.HardCopy(possibility.regA)

		if k < 0 {
			continue
		}

		for y := 0; y < 8; y++ {
			switch y {
			case 0:
				regA[len(regA)-3*k-3] = 0
				regA[len(regA)-3*k-2] = 0
				regA[len(regA)-3*k-1] = 0
			case 1:
				regA[len(regA)-3*k-3] = 0
				regA[len(regA)-3*k-2] = 0
				regA[len(regA)-3*k-1] = 1
			case 2:
				regA[len(regA)-3*k-3] = 0
				regA[len(regA)-3*k-2] = 1
				regA[len(regA)-3*k-1] = 0
			case 3:
				regA[len(regA)-3*k-3] = 0
				regA[len(regA)-3*k-2] = 1
				regA[len(regA)-3*k-1] = 1
			case 4:
				regA[len(regA)-3*k-3] = 1
				regA[len(regA)-3*k-2] = 0
				regA[len(regA)-3*k-1] = 0
			case 5:
				regA[len(regA)-3*k-3] = 1
				regA[len(regA)-3*k-2] = 0
				regA[len(regA)-3*k-1] = 1
			case 6:
				regA[len(regA)-3*k-3] = 1
				regA[len(regA)-3*k-2] = 1
				regA[len(regA)-3*k-1] = 0
			case 7:
				regA[len(regA)-3*k-3] = 1
				regA[len(regA)-3*k-2] = 1
				regA[len(regA)-3*k-1] = 1
			}

			// b = A[A.len-3*(k+1),A.len-3*k-1] ^ 011
			// b = A[A.len-3*(k+1),A.len-3*k-1] ^ 3
			b := 0
			for i := 0; i <= 2; i++ {
				if regA[len(regA)-3*k-i-1] == 1 {
					b += int(math.Pow(2, float64(i)))
				}
			}
			b = b ^ 3

			// x := A[A.len-3*(k+1)-b,A.len-3*k-b)
			x := 0
			for i := b; i <= b+2; i++ {
				if len(regA)-3*k-i-1 < 0 {
					continue
				}

				if regA[len(regA)-3*k-i-1] == 1 {
					x += int(math.Pow(2, float64(i-b)))
				}
			}

			// new Possibility => try it
			if x^b^5 == output[k] {
				newPossibility := PossiblePath{k - 1, helpers.HardCopy(regA)}
				possibilities.Push(newPossibility)

				if k == 0 {
					finishedPossiblities = append(finishedPossiblities, newPossibility)
				}
			}
		}

	}

	for _, finishedPossbility := range finishedPossiblities {
		fmt.Println(finishedPossbility.regA)
	}
}

/*
	A=100
	2 => bst	B=A&111			3bits
	1 => bxl	B=B^011			3bits
	7 => cdv	C=A>>B			xbits
	4 => bxc	B=B^C			xbits
	0 => adv	A=A>>3
	1 => bxl	B=B^101			xbits
	5 => out	output: B&111	3bits
	3 => jnz

	- last three bits don't matter
	- outputs as many digits as ceil(countBits(A)/3)

	output_k = {
		b = A[A.len-3*(k+1),A.len-3*k-1] ^ 011
		return A[A.len-3*(k+1),A.len-3*k-1] ^ A[A.len-3*(k+1)-1-b,A.len-3*k-1-b] ^ 110
	}

	100 => 010

	output:= 2,4,1,3,7,5,4,7,0,3,1,5,5,5,3,0
*/
