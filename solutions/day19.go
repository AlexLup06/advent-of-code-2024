package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"alexlupatsiy.com/aoc24/helpers"
)

type PossiblePattern struct {
	remaining   string
	currentNode int
}

type CachePattern struct {
	value string
	count int
}

func Day19() {
	file, err := os.OpenFile("./input/input19.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	patterns := []string{}
	designs := []string{}
	scanPatterns := true
	for sc.Scan() {
		line := sc.Text()

		if line == "" {
			continue
		}

		if scanPatterns {
			patterns = strings.Split(line, ", ")
			scanPatterns = false
		} else {
			designs = append(designs, line)
		}
	}

	longestPattern := 0
	for _, pattern := range patterns {
		if len(pattern) > longestPattern {
			longestPattern = len(pattern)
		}
	}

	// actually count
	solution := []CachePattern{}
	for i, design := range designs {
		fmt.Println("check pattern: ", i)
		cache := []CachePattern{}

		for j := 0; j < len(design); j++ {
			count := 0
			startPattern := PossiblePattern{design[len(design)-1-j:], 0}
			possiblePaths := helpers.Stack[PossiblePattern]()
			possiblePaths.Push(startPattern)

			for {
				if possiblePaths.Length() == 0 {
					break
				}
				currentPath := possiblePaths.Pop()

				foundPattern := false
				for _, cacheObject := range cache {
					if cacheObject.value == currentPath.remaining {
						count += cacheObject.count
						foundPattern = true
						break
					}
				}
				if foundPattern {
					continue
				}

				for _, pattern := range patterns {
					if len(pattern) > len(currentPath.remaining) {
						continue
					}

					if pattern == currentPath.remaining[0:len(pattern)] {
						newPossiblePath := PossiblePattern{}
						newPossiblePath.remaining = currentPath.remaining[len(pattern):]
						possiblePaths.Push(newPossiblePath)
						foundPattern = true
					}
				}

				if !foundPattern && len(currentPath.remaining) == 0 {
					count++
				}

				if !foundPattern && len(currentPath.remaining) > 0 {
					continue
				}
			}
			cache = append(cache, CachePattern{design[len(design)-1-j:], count})

			if j == len(design)-1 {
				solution = append(solution, cache[len(design)-1])
			}
		}
	}

	count := 0
	for _, v := range solution {
		count += v.count
	}
	fmt.Println(count)
}
