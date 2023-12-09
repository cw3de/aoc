package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"

	"github.com/cw3de/aoc/prime"
)

type Node struct {
	Name  string
	Left  *Node
	Right *Node
}

type Network struct {
	Instructions string
	Nodes        map[string]*Node
}

func (n *Network) FindOrCreateNode(name string) *Node {
	if node, ok := n.Nodes[name]; ok {
		return node
	}
	node := &Node{Name: name}
	n.Nodes[name] = node
	return node
}

func (n *Network) CycleUntilZ(start *Node) int {
	count := 0
	node := start
	for {
		if strings.HasSuffix(node.Name, "Z") {
			return count
		}
		for _, step := range n.Instructions {
			switch step {
			case 'L':
				node = node.Left
			case 'R':
				node = node.Right
			default:
				panic("unknown step")
			}
			count++
		}
	}
}

func LoadNetwork(filename string) *Network {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	network := &Network{
		Instructions: lines[0],
		Nodes:        make(map[string]*Node),
	}
	re := regexp.MustCompile(`^(\w\w\w) = \((\w\w\w), (\w\w\w)\)$`)

	for _, line := range lines[1:] {
		m := re.FindStringSubmatch(line)
		if m != nil {
			node := network.FindOrCreateNode(m[1])
			node.Left = network.FindOrCreateNode(m[2])
			node.Right = network.FindOrCreateNode(m[3])
		}
	}
	return network
}

func ShowNetwork(network *Network) {
	for name, node := range network.Nodes {
		fmt.Printf("%s: %s, %s\n", name, node.Left.Name, node.Right.Name)
	}
}

func task1(filename string) {
	network := LoadNetwork(filename)
	// ShowNetwork(network)

	node, ok := network.Nodes["AAA"]
	if !ok {
		panic("no AAA node")
	}

	count := 0
	for {
		for _, step := range network.Instructions {
			if node.Name == "ZZZ" {
				fmt.Printf("Found ZZZ after %d steps\n", count)
				return
			}

			switch step {
			case 'L':
				node = node.Left
			case 'R':
				node = node.Right
			default:
				panic("unknown step")
			}
			count++
		}
	}
}

func AreAllNodesFinal(nodes []*Node) bool {
	for _, node := range nodes {
		if !strings.HasSuffix(node.Name, "Z") {
			return false
		}
	}
	return true
}

func task2(filename string) {
	network := LoadNetwork(filename)
	// ShowNetwork(network)

	// finde die Zykluslänge für alle Knoten, die mit A enden
	cycleLength := []int{}
	for _, node := range network.Nodes {
		if strings.HasSuffix(node.Name, "A") {
			fmt.Printf("Found %s\n", node.Name)
			cycle := network.CycleUntilZ(node)
			fmt.Printf("Cycle length: %d\n", cycle)
			cycleLength = append(cycleLength, cycle)
		}
	}
	fmt.Println(cycleLength)

	maximum := slices.Max(cycleLength)

	// finde das kleinste gemeinsame Vielfache
	primes := prime.NewPrimes(maximum)
	result := primes.GetGreatestCommonDivisor(cycleLength)
	fmt.Printf("Result: %d\n", result)
}

func main() {
	// task1("sample1.txt")
	// task1("sample2.txt")
	// task1("input.txt")
	task2("sample3.txt")
	task2("input.txt")
}
