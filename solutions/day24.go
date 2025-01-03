package solutions

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	"alexlupatsiy.com/aoc24/helpers"
)

type Gate struct {
	id       string
	gateType string
	input1Id string
	input2Id string
	outputId string
}

type Wire struct {
	id    string
	value int
}

type Circuit struct {
	gates []Gate
	wires []Wire
}

func (c *Circuit) addWire(id string, value int) {
	for _, wire := range c.wires {
		if wire.id == id {
			return
		}
	}
	c.wires = append(c.wires, Wire{id, value})
}

func (c *Circuit) addGate(gateType, input1Id, input2Id, outputId string) {
	id := gateType + input1Id + input2Id + outputId
	for _, gate := range c.gates {
		if gate.id == id {
			return
		}
	}
	c.gates = append(c.gates, Gate{id, gateType, input1Id, input2Id, outputId})
	c.addWire(input1Id, -1)
	c.addWire(input2Id, -1)
	c.addWire(outputId, -1)
}

func (c *Circuit) runGate(gate Gate) bool {
	inputWire1 := c.getWire(gate.input1Id)
	inputWire2 := c.getWire(gate.input2Id)

	if inputWire1.value == -1 || inputWire2.value == -1 {
		return false
	}

	c.writeOutput(inputWire1.value, inputWire2.value, gate.outputId, gate.gateType)
	return true
}

func (c *Circuit) writeOutput(input1, input2 int, outputId, gateType string) {
	outputIndex := c.getWireIndex(outputId)
	switch gateType {
	case "XOR":
		c.wires[outputIndex].value = input1 ^ input2
	case "OR":
		c.wires[outputIndex].value = input1 | input2
	case "AND":
		c.wires[outputIndex].value = input1 & input2
	}
}

func (c *Circuit) getWire(id string) Wire {
	for _, wire := range c.wires {
		if wire.id == id {
			return wire
		}
	}
	return Wire{}
}

func (c *Circuit) getWireIndex(id string) int {
	for i, wire := range c.wires {
		if wire.id == id {
			return i
		}
	}
	return -1
}

func Day24() {
	file, err := os.OpenFile("./input/input24.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	circuit := Circuit{}

	scanningStarters := true
	for sc.Scan() {
		line := sc.Text()

		if line == "" {
			scanningStarters = false
			continue
		}

		if scanningStarters {
			elements := strings.Split(line, ":")
			circuit.addWire(elements[0], helpers.STOI(strings.Fields(elements[1])[0]))
		}

		if !scanningStarters {
			elements := strings.Fields(line)
			circuit.addGate(elements[1], elements[0], elements[2], elements[4])
		}
	}

	stillGoing := true
	for stillGoing {
		keepGoing := false
		for _, gate := range circuit.gates {
			isDone := circuit.runGate(gate)
			keepGoing = keepGoing || !isDone
		}
		stillGoing = keepGoing
	}

	output := helpers.InitSlice(46, -1)
	for _, wire := range circuit.wires {
		if string(wire.id[0]) == "z" {
			number := helpers.STOI(string(wire.id[1]) + string(wire.id[2]))
			output[number] = wire.value
		}
	}

	xVal := helpers.InitSlice(45, -1)
	yVal := helpers.InitSlice(45, -1)
	for _, wire := range circuit.wires {
		if string(wire.id[0]) == "x" {
			number := helpers.STOI(string(wire.id[1]) + string(wire.id[2]))
			xVal[number] = wire.value
		}
		if string(wire.id[0]) == "y" {
			number := helpers.STOI(string(wire.id[1]) + string(wire.id[2]))
			yVal[number] = wire.value
		}
	}

	fmt.Print(binaryToInt(output) == binaryToInt(xVal)+binaryToInt(yVal), "  =>  ")
	fmt.Println(binaryToInt(xVal), " + ", binaryToInt(yVal), " = ", binaryToInt(output))
}

func binaryToInt(binary []int) int {
	result := 0
	for i, bit := range binary {
		if bit == 1 {
			result += int(math.Pow(2, float64(i)))
		}
	}
	return result
}
