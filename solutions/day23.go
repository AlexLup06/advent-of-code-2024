package solutions

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Graph struct {
	nodes []Node
}

type Node struct {
	id       string
	adjacent []string
}

func (g *Graph) addNode(id string) {
	for _, node := range g.nodes {
		if node.id == id {
			return
		}
	}
	g.nodes = append(g.nodes, Node{id, []string{}})
}

func (g *Graph) addEdge(from, to string) {
	nodes := g.nodes
	for _, node := range nodes {
		if node.id == from {
			g.nodes[g.getNodeIndex(from)].adjacent = append(g.nodes[g.getNodeIndex(from)].adjacent, to)
		}
		if node.id == to {
			g.nodes[g.getNodeIndex(to)].adjacent = append(g.nodes[g.getNodeIndex(to)].adjacent, from)
		}
	}
}

func (g *Graph) getNodeIndex(id string) int {
	for i, node := range g.nodes {
		if node.id == id {
			return i
		}
	}
	return -1
}

func (g *Graph) getNode(id string) Node {
	for _, node := range g.nodes {
		if node.id == id {
			return node
		}
	}
	return Node{}
}

func Day23() {
	file, err := os.OpenFile("./input/input23.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	graph := Graph{}

	for sc.Scan() {
		line := sc.Text()
		nodes := strings.Split(line, "-")
		graph.addNode(nodes[0])
		graph.addNode(nodes[1])
		graph.addEdge(nodes[0], nodes[1])
	}

	// Part 1
	// alreadyChecked := []string{}
	// count := 0
	// for _, node := range graph.nodes {
	// 	adjacentsIds := node.adjacent

	// 	for i := 0; i < len(adjacentsIds); i++ {
	// 		for j := i + 1; j < len(adjacentsIds); j++ {
	// 			nodeId1 := adjacentsIds[i]
	// 			nodeId2 := adjacentsIds[j]

	// 			node1 := graph.getNode(nodeId1)
	// 			node2 := graph.getNode(nodeId2)

	// 			if slices.Contains(node1.adjacent, nodeId2) && slices.Contains(node2.adjacent, nodeId1) && !slices.Contains(alreadyChecked, nodeId1) && !slices.Contains(alreadyChecked, nodeId2) {
	// 				for i, char := range string(node.id + nodeId1 + nodeId2) {
	// 					if string(char) == "t" && i%2 == 0 {
	// 						count++
	// 						break
	// 					}
	// 				}
	// 			}
	// 		}
	// 	}
	// 	alreadyChecked = append(alreadyChecked, node.id)
	// }

	// sort Nodes after the number of adjacent nodes
	sorted := []Node{}

	for _, node := range graph.nodes {
		if len(sorted) == 0 {
			sorted = append(sorted, node)
			continue
		}

		for i, current := range sorted {
			if len(current.adjacent) < len(node.adjacent) {
				sorted = slices.Insert(sorted, i, node)
				break
			}
		}

		sorted = append(sorted, node)
	}
}
