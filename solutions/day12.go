package solutions

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type GardenPlot struct {
	id       int
	value    string
	sidesOut int
	checked  bool
	x, y     int
}

type Region struct {
	area, perimeter int
}

func Day12() {
	file, err := os.OpenFile("./input/input12.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
	defer file.Close()

	rd := bufio.NewReader(file)

	gardenMap := [][]GardenPlot{}
	regions := map[int]Region{}
	row := []GardenPlot{{-1, ".", -1, false, -1, -1}}
	y := 1
	x := 1
	for {
		r, _, err := rd.ReadRune()

		if err != nil {
			break
		}

		if r == 10 {
			row = append(row, GardenPlot{-1, ".", -1, false, -1, -1})
			gardenMap = append(gardenMap, row)
			row = []GardenPlot{{-1, ".", -1, false, -1, -1}}
			y++
			x = 1
			continue
		}
		row = append(row, GardenPlot{-1, string(r), -1, false, x, y})
		x++
	}
	border := []GardenPlot{}
	for range gardenMap[0] {
		border = append(border, GardenPlot{-1, ".", -1, false, -1, -1})
	}
	gardenMap = append(gardenMap, border)
	gardenMap = slices.Insert(gardenMap, 0, border)

	id := 0
	for y := 1; y < len(gardenMap)-1; y++ {
		for x := 1; x < len(gardenMap[0])-1; x++ {
			if gardenMap[y][x].checked {
				continue
			}

			plotsToCheck := []GardenPlot{gardenMap[y][x]}

			for len(plotsToCheck) > 0 {
				plotToCheck := plotsToCheck[0]
				neighbours := []GardenPlot{
					gardenMap[plotToCheck.y-1][plotToCheck.x],   // top 0
					gardenMap[plotToCheck.y-1][plotToCheck.x-1], // top-left
					gardenMap[plotToCheck.y][plotToCheck.x-1],   // left 1
					gardenMap[plotToCheck.y+1][plotToCheck.x-1], // bottom-left
					gardenMap[plotToCheck.y+1][plotToCheck.x],   // bottom 2
					gardenMap[plotToCheck.y+1][plotToCheck.x+1], // bottom-right
					gardenMap[plotToCheck.y][plotToCheck.x+1],   // right 3
					gardenMap[plotToCheck.y-1][plotToCheck.x+1], // top-right
				}
				updatedPlot, sameRegionNeighbours, discard := checkNeighbours(plotsToCheck[0], neighbours, id)
				plotsToCheck = slices.Delete(plotsToCheck, 0, 1)
				if discard {
					continue
				}

				for _, sameRegionNeighbour := range sameRegionNeighbours {
					if !slices.Contains(plotsToCheck, sameRegionNeighbour) {
						plotsToCheck = append(plotsToCheck, sameRegionNeighbour)
					}
				}

				gardenMap[plotToCheck.y][plotToCheck.x] = updatedPlot
				// update or create region
				region := regions[updatedPlot.id]
				region.area++
				region.perimeter += updatedPlot.sidesOut
				regions[updatedPlot.id] = region
			}
			id++
		}
	}

	price := 0
	for _, region := range regions {
		price += region.area * region.perimeter
	}
	fmt.Println(regions)
	fmt.Println(price)

}

/*
Idea of Part 2: count corners for each area

. | A	 or  A A .
B + -	   	 A + -		=> this mean we are an indent "left" corner or an outdent "left" corner of each edge
B B .	 	 . | B
*/
func checkNeighbours(current GardenPlot, neighbours []GardenPlot, id int) (GardenPlot, []GardenPlot, bool) {
	if current.value == "." || current.checked {
		return current, neighbours, true
	}

	sidesOut := 0
	sameRegionNeighbours := []GardenPlot{}

	// need to check each neighbour one by one
	for i, neighbour := range neighbours {
		if i%2 != 0 {
			continue
		}
		// Part 1
		if current.value != neighbour.value {
			sidesOut++
		}

		// Part 2
		/*
			1. Iteration:
			2. Iteration:
			3. Iteration:
			4. Iteration:
		*/
		// isIndentCorner := current.value == neighbours[i].value && current.value != neighbours[i+1].value && current.value == neighbours[(i+2)%8].value
		// isOutdentCorner := current.value != neighbours[i].value && current.value != neighbours[(i+2)%8].value
		// if isIndentCorner || isOutdentCorner {
		// 	sidesOut++
		// }

		if current.value == neighbour.value && !neighbour.checked {
			sameRegionNeighbours = append(sameRegionNeighbours, neighbour)
		}
	}

	current.id = id
	current.sidesOut = sidesOut
	current.checked = true
	return current, sameRegionNeighbours, false
}
