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

func HardCopy[T any](orig []T) []T {
	return append(make([]T, 0, len(orig)), orig...)
}

type stack[T any] struct {
	Push   func(T)
	Pop    func() T
	Length func() int
}

func Stack[T any]() stack[T] {
	slice := make([]T, 0)
	return stack[T]{
		Push: func(i T) {
			slice = append(slice, i)
		},
		Pop: func() T {
			res := slice[len(slice)-1]
			slice = slice[:len(slice)-1]
			return res
		},
		Length: func() int {
			return len(slice)
		},
	}
}
