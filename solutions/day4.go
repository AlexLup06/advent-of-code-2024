package solutions

import (
	"bufio"
	"fmt"
	"os"
)

func Day4() {
	file, err := os.OpenFile("./input/input4.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
	defer file.Close()

	rd := bufio.NewReader(file)

	characters := [][]string{}
	row := []string{}

	addX := []string{".", ".", "."}
	row = append(row, addX...)
	for {
		r, _, err := rd.ReadRune()

		if err != nil {
			break
		}
		if r == 10 {
			// make matrix bigger to not handle lots of cases
			row = append(row, addX...)
			characters = append(characters, row)
			row = []string{}
			row = append(row, addX...)
		} else {
			row = append(row, string(r))
		}
	}

	// make matrix bigger to not handle lots of cases
	lenX := len(characters[0])
	addY := [][]string{make([]string, lenX), make([]string, lenX), make([]string, lenX)}
	charMatrix := [][]string{}
	charMatrix = append(charMatrix, addY...)
	charMatrix = append(charMatrix, characters...)
	charMatrix = append(charMatrix, addY...)

	// lenY := len(charMatrix)
	sum1 := 0
	sum2 := 0
	for y, row := range charMatrix {
		for x, v := range row {
			if v == "X" {
				// left
				sum1 += checkXMAS(charMatrix[y][x-1], charMatrix[y][x-2], charMatrix[y][x-3])
				// up left
				sum1 += checkXMAS(charMatrix[y-1][x-1], charMatrix[y-2][x-2], charMatrix[y-3][x-3])
				// up
				sum1 += checkXMAS(charMatrix[y-1][x], charMatrix[y-2][x], charMatrix[y-3][x])
				// up right
				sum1 += checkXMAS(charMatrix[y-1][x+1], charMatrix[y-2][x+2], charMatrix[y-3][x+3])
				// right
				sum1 += checkXMAS(charMatrix[y][x+1], charMatrix[y][x+2], charMatrix[y][x+3])
				// down right
				sum1 += checkXMAS(charMatrix[y+1][x+1], charMatrix[y+2][x+2], charMatrix[y+3][x+3])
				// down
				sum1 += checkXMAS(charMatrix[y+1][x], charMatrix[y+2][x], charMatrix[y+3][x])
				// down left
				sum1 += checkXMAS(charMatrix[y+1][x-1], charMatrix[y+2][x-2], charMatrix[y+3][x-3])
			}

			if v == "A" {
				// M . S
				// . A .
				// M . S
				sum2 += checkX_MAS(charMatrix[y-1][x-1], charMatrix[y+1][x+1], charMatrix[y+1][x-1], charMatrix[y-1][x+1])
				// M . M
				// . A .
				// S . S
				sum2 += checkX_MAS(charMatrix[y-1][x-1], charMatrix[y+1][x+1], charMatrix[y-1][x+1], charMatrix[y+1][x-1])
				// S . M
				// . A .
				// S . M
				sum2 += checkX_MAS(charMatrix[y-1][x+1], charMatrix[y+1][x-1], charMatrix[y+1][x+1], charMatrix[y-1][x-1])
				// S . S
				// . A .
				// M . M
				sum2 += checkX_MAS(charMatrix[y+1][x+1], charMatrix[y-1][x-1], charMatrix[y+1][x-1], charMatrix[y-1][x+1])
			}

		}
	}
	fmt.Println(sum1)
	fmt.Println(sum2)
}

func checkXMAS(m, a, s string) int {
	if m == "M" && a == "A" && s == "S" {
		return 1
	}
	return 0
}

func checkX_MAS(m1, s1, m2, s2 string) int {
	if m1 == "M" && s1 == "S" && m2 == "M" && s2 == "S" {
		return 1
	}
	return 0
}
