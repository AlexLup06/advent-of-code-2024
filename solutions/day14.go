package solutions

import (
	"fmt"
	"os"
)

func Day14() {
	file, err := os.OpenFile("./input/input14.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Printf("Error opnening ifle %v", err)
	}
	defer file.Close()
}
