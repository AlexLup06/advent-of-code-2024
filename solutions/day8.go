package solutions

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"alexlupatsiy.com/aoc24/helpers"
)

// type Line strcut {

// }

func Day8() {
	file, err := os.OpenFile("./input/input8.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Printf("Error opnening file: %v\n", err)
	}
	defer file.Close()

	antennaMap := [][]string{}
	antinodeMap := [][]string{}

	rd := bufio.NewReader(file)

	row := []string{}
	for {
		r, _, err := rd.ReadRune()
		if err != nil {
			break
		}

		if r == 10 {
			antennaMap = append(antennaMap, row)
			antinodeMap = append(antinodeMap, make([]string, len(row)))

			row = []string{}
			continue
		}
		row = append(row, string(r))
	}

	frequencyToLocations := map[string][]helpers.Vector{}

	for y, row := range antennaMap {
		for x, v := range row {
			if v != "." {
				locations := frequencyToLocations[v]
				locations = append(locations, helpers.Vector{X: float64(x), Y: float64(y)})
				frequencyToLocations[v] = locations
			}
		}
	}

	maxY := len(antennaMap[0]) - 1
	maxX := len(antennaMap) - 1

	corner1 := helpers.Vector{X: 0, Y: 0}                         // top-left
	corner2 := helpers.Vector{X: float64(maxX), Y: 0}             // top-right
	corner3 := helpers.Vector{X: float64(maxY), Y: float64(maxY)} // bottom-right
	corner4 := helpers.Vector{X: 0, Y: float64(maxY)}             // bottom-left

	border1 := helpers.Line{A: corner1, B: corner2, AB: corner1.Build(corner2)} // top
	border2 := helpers.Line{A: corner2, B: corner3, AB: corner2.Build(corner3)} // right
	border3 := helpers.Line{A: corner3, B: corner4, AB: corner3.Build(corner4)} // bottom
	border4 := helpers.Line{A: corner4, B: corner1, AB: corner4.Build(corner1)} // left

	for _, locations := range frequencyToLocations {
		for _, v1 := range locations {
			for _, v2 := range locations {
				if v1 == v2 {
					break
				}
				// Part 1
				// vector := v1.Build(v2)

				// antinode1 := v2.Add(vector)
				// antinode2 := v1.Add(helpers.Vector{X: -vector.X, Y: -vector.Y})

				// if isInsideMap(maxX, maxY, antinode1) {
				// 	antinodeMap[antinode1.Y][antinode1.X] = "#"
				// }
				// if isInsideMap(maxX, maxY, antinode2) {
				// 	antinodeMap[antinode2.Y][antinode2.X] = "#"
				// }

				// Part 2
				line := helpers.Line{A: v1, B: v2, AB: v1.Build(v2)}
				t1, u1 := line.CalcIntersection(border1)
				t2, u2 := line.CalcIntersection(border2)
				t3, u3 := line.CalcIntersection(border3)
				t4, u4 := line.CalcIntersection(border4)

				if u1 >= 0 && u1 <= 1 {
					// fmt.Println("intersection with top: ", u1)
					for i := 0; i <= int(math.Floor(math.Abs(t1))); i++ {
						antinode := line.GetVector(float64(helpers.Sgn(t1)) * float64(i))
						antinodeMap[int(antinode.Y)][int(antinode.X)] = "#"
					}
				}
				if u2 >= 0 && u2 <= 1 {
					// fmt.Println("intersection with right: ", u2)
					for i := 0; i <= int(math.Floor(math.Abs(t2))); i++ {
						antinode := line.GetVector(float64(helpers.Sgn(t2)) * float64(i))
						antinodeMap[int(antinode.Y)][int(antinode.X)] = "#"
					}
				}
				if u3 >= 0 && u3 <= 1 {
					// fmt.Println("intersection with bottom: ", u3)
					for i := 0; i <= int(math.Floor(math.Abs(t3))); i++ {
						antinode := line.GetVector(float64(helpers.Sgn(t3)) * float64(i))
						antinodeMap[int(antinode.Y)][int(antinode.X)] = "#"
					}
				}
				if u4 >= 0 && u4 <= 1 {
					// fmt.Println("intersection with left: ", u4)
					for i := 0; i <= int(math.Floor(math.Abs(t4))); i++ {
						antinode := line.GetVector(float64(helpers.Sgn(t4)) * float64(i))
						antinodeMap[int(antinode.Y)][int(antinode.X)] = "#"
					}
				}
			}
		}
	}

	count := 0
	for _, row := range antinodeMap {
		for _, v := range row {
			if v == "#" {
				count++
				fmt.Printf("# ")
			} else {
				fmt.Printf(". ")
			}
		}
		fmt.Println()
	}

	fmt.Println(count)
}

// func isInsideMap(maxX, maxY int, vector helpers.Vector) bool {
// 	return vector.X >= 0 && vector.X <= maxX && vector.Y >= 0 && vector.Y <= maxY
// }
