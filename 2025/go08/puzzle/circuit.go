package puzzle

import (
	"fmt"
	"slices"
)

type Circuit struct {
	Boxes  []int
	Number string
}

func (c *Circuit) Count() int {
	return len(c.Boxes)
}

func (c *Circuit) containsPoint(index int) bool {
	for _, j := range c.Boxes {
		if index == j {
			return true
		}
	}
	return false
}

func (c *Circuit) showBoxes() {
	fmt.Printf(": %s =", c.Number)
	for _, b := range c.Boxes {
		fmt.Printf(" %d", b)
	}
}

func MakeCircuitList(
	sortedJunction []Junction,
	maxConnections int,
	numberOfBoxes int,
	beVerbose bool) ([]*Circuit, *Junction) {

	circuits := make([]*Circuit, 0)

	findCircuit := func(pointIndex int) (int, *Circuit) {
		for i, c := range circuits {
			if c.containsPoint(pointIndex) {
				return i, c
			}
		}
		return -1, nil
	}

	conections := 0
	counter := 0

	for i, junction := range sortedJunction {

		_, circuitA := findCircuit(junction.PointA)
		indexB, circuitB := findCircuit(junction.PointB)

		if beVerbose {
			fmt.Printf("junction %3d,%3d (", junction.PointA, junction.PointB)
			if circuitA != nil {
				fmt.Printf("%s", circuitA.Number)
			} else {
				fmt.Printf("none")
			}
			if circuitB != nil {
				fmt.Printf(",%s", circuitB.Number)
			} else {
				fmt.Printf(",none")
			}
			fmt.Printf("): ")
		}

		if circuitA != nil && circuitB != nil {
			if circuitA == circuitB {
				// same circuit, nothing to do
				if beVerbose {
					fmt.Printf("both boxes connected to same circuit")
				}
			} else {
				// join circuit A and B
				circuitA.Boxes = append(
					circuitA.Boxes, circuitB.Boxes...)
				circuits = append(circuits[:indexB], circuits[indexB+1:]...)
				conections++
				if beVerbose {
					fmt.Printf("join circuit:")
					circuitA.showBoxes()
				}
			}
		} else if circuitA != nil {
			// append PointB to circuitA
			circuitA.Boxes = append(circuitA.Boxes, junction.PointB)
			conections++
			if beVerbose {
				fmt.Printf("appended to left circuit")
				circuitA.showBoxes()
			}

		} else if circuitB != nil {
			// append PointA to circB
			circuitB.Boxes = append(circuitB.Boxes, junction.PointA)
			conections++
			if beVerbose {
				fmt.Printf("appended to right circuit")
				circuitB.showBoxes()
			}
		} else {
			// create a new circuit
			counter++
			newCircuit := &Circuit{
				Boxes:  []int{junction.PointA, junction.PointB},
				Number: fmt.Sprintf("C%03d", counter),
			}
			circuits = append(circuits, newCircuit)
			conections++
			if beVerbose {
				fmt.Printf("new circuit")
				newCircuit.showBoxes()
			}
		}
		if beVerbose {
			fmt.Printf("\n")
		}

		if maxConnections > 0 {
			// task 1: stop at 1000 connections
			if i+1 == maxConnections {
				fmt.Printf("found %d circuits for %d connections\n", len(circuits), maxConnections)
				break
			}
		} else {
			// task 2: stop when alle boxes are connected in a single junction
			if len(circuits) == 1 && len(circuits[0].Boxes) == numberOfBoxes {
				fmt.Printf("last junction: %d and %d\n",
					junction.PointA, junction.PointB)
				return nil, &junction
			}
		}
	}

	slices.SortFunc(circuits, func(a, b *Circuit) int {
		return len(b.Boxes) - len(a.Boxes)
	})

	if beVerbose {
		for i, c := range circuits {
			fmt.Printf("%s (%d):", c.Number, c.Count())
			for _, b := range c.Boxes {
				fmt.Printf(" %d", b)
			}
			fmt.Printf("\n")
			if i > maxConnections+5 {
				break
			}
		}
	}
	return circuits, nil
}
