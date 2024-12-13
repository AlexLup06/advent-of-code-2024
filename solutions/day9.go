package solutions

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

type DiskSpace struct {
	id, quantity int
}

func Day9() {
	file, err := os.OpenFile("./input/input9.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}

	defer file.Close()

	rd := bufio.NewReader(file)
	diskMap := []DiskSpace{}
	i := 0
	for {
		r, _, err := rd.ReadRune()

		if err != nil || r == 10 {
			break
		}

		quantity, err := strconv.Atoi(string(r))
		if err != nil {
			fmt.Printf("Error converting number: %v\n", err)
		}

		switch i % 2 {
		case 0:
			diskMap = append(diskMap, DiskSpace{id: i / 2, quantity: quantity})
		case 1:
			diskMap = append(diskMap, DiskSpace{id: -1, quantity: quantity})
		}
		i++
	}

	//Part 1
	// virtualDiskMap := []DiskSpace{}
	// pointerFinalMemory := len(diskMap) - 1
	// for i, v := range diskMap {
	// 	if i > pointerFinalMemory {
	// 		break
	// 	}
	// 	if v.id != -1 {
	// 		virtualDiskMap = append(virtualDiskMap, v)
	// 		continue
	// 	}

	// 	freeSpace := v.quantity

	// 	for freeSpace > 0 {
	// 		finalUsedMemory := diskMap[pointerFinalMemory]
	// 		spaceUsed := finalUsedMemory.quantity

	// 		if i > pointerFinalMemory {
	// 			// free space is already all used up => we got contineous block of memory
	// 			break
	// 		}
	// 		if spaceUsed == 0 {
	// 			// file is 0 big => just go to the next one
	// 			pointerFinalMemory -= 2
	// 			continue
	// 		}
	// 		if freeSpace >= spaceUsed {
	// 			// freespace is NOT used up by the finalUsedMemory block => jump to the next usedMemory block
	// 			diskMap[pointerFinalMemory].quantity = 0
	// 			virtualDiskMap = append(virtualDiskMap, DiskSpace{id: finalUsedMemory.id, quantity: spaceUsed})
	// 			freeSpace -= spaceUsed
	// 			pointerFinalMemory -= 2
	// 		} else {
	// 			// all freespace is used up by the finalUsedMemory block => decrement space of finalUsedMemoryBlock
	// 			diskMap[pointerFinalMemory].quantity = spaceUsed - freeSpace
	// 			virtualDiskMap = append(virtualDiskMap, DiskSpace{id: finalUsedMemory.id, quantity: freeSpace})
	// 			freeSpace = 0
	// 		}
	// 	}
	// }

	// Part 2
	pointerMemoryToMove := len(diskMap) - 1
	for pointerMemoryToMove > 1 {
		memoryToMove := diskMap[pointerMemoryToMove]
		if memoryToMove.id == -1 {
			pointerMemoryToMove--
			continue
		}

		for i, v := range diskMap {
			if pointerMemoryToMove <= i {
				break
			}
			if v.id != -1 {
				continue
			}

			freeSpace := v.quantity
			if freeSpace < memoryToMove.quantity {
				// freeSpace is too little => go to next if available
				continue
			}
			// big enough free space found => decrease freespace by memory
			diskMap[i].quantity = freeSpace - memoryToMove.quantity
			// insert new freeSpace which is decreased by the memory block
			diskMap[pointerMemoryToMove].id = -1
			diskMap = slices.Insert(diskMap, i, memoryToMove)
			pointerMemoryToMove++
			break
		}
		pointerMemoryToMove--
	}

	sum := 0
	i = 0
	for _, v := range diskMap {
		if v.id == -1 {
			i += v.quantity
			continue
		}
		for j := i; j < i+v.quantity; j++ {
			sum += j * v.id
		}
		i += v.quantity
	}
	fmt.Println(sum)
}
