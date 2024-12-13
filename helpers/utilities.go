package helpers

import (
	"fmt"
	"strconv"
)

func InitSlice(length int, value int) []int {
	slice := make([]int, length)
	for i := range slice {
		slice[i] = value
	}
	return slice
}

func LenInt(i int) int {
	return len(strconv.Itoa(i))
}

func STOI(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("Error converting string to number:%v\n", err)
	}
	return x
}
