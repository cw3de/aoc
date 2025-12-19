package puzzle

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// type Lights []bool

type Button struct {
	Numbers []int  // index into lights or joltage
	Pattern []bool // true, if button cotains index
}

type Machine struct {
	Lights  []bool
	Joltage []int
	Buttons []*Button
}

func (b *Button) String() string {
	var sb strings.Builder
	sb.WriteRune('(')
	for _, p := range b.Pattern {
		if p {
			sb.WriteString("1")
		} else {
			sb.WriteString(".")
		}
	}
	sb.WriteRune(')')
	return sb.String()
}

func Show(m *Machine) {
	fmt.Printf("-------------------------\n")
	fmt.Printf("[")
	for _, light := range m.Lights {
		if light {
			fmt.Printf("#")
		} else {
			fmt.Printf(".")
		}
	}
	fmt.Printf("] : %v\n", m.Joltage)
	for _, but := range m.Buttons {
		fmt.Printf(" ")
		for _, ison := range but.Pattern {
			if ison {
				fmt.Printf("1")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf(" : %v\n", but.Numbers)
	}
}

func Load(filename string) []*Machine {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return Parse(string(data))
}

func Parse(text string) []*Machine {
	lines := strings.Split(text, "\n")
	list := make([]*Machine, 0)
	for _, line := range lines {
		if len(line) > 0 {
			list = append(list, ParseLine(line))
		}
	}
	return list
}

func ParseLine(line string) *Machine {
	fields := strings.Split(line, " ")
	if len(fields) < 3 {
		panic("need a minimum of 3 fields")
	}
	lastField := len(fields) - 1
	lights := ParseLights(fields[0])
	joltage := ParseListOfInt(fields[lastField], '{', '}')
	size := len(lights)
	if len(joltage) != size {
		panic("len(lights) must match len(joltage)")
	}
	buttons := make([]*Button, 0)

	for f := 1; f < lastField; f++ {
		numbers := ParseListOfInt(fields[f], '(', ')')
		pattern := make([]bool, 0)
		for b := 0; b < size; b++ {
			pattern = append(pattern, false)
		}
		for _, b := range numbers {
			pattern[b] = true
		}

		buttons = append(buttons, &Button{
			Numbers: numbers,
			Pattern: pattern,
		})
	}

	return &Machine{
		Lights:  lights,
		Joltage: joltage,
		Buttons: buttons,
	}
}

// a string like "[.##.]"
func ParseLights(word string) []bool {
	lights := make([]bool, 0)
	last := len(word) - 1
	for i, c := range word {
		switch i {
		case 0:
			if c != '[' {
				panic("'[' expected")
			}
		case last:
			if c != ']' {
				panic("']' expected")
			}
		default:
			switch c {
			case '.':
				lights = append(lights, false)
			case '#':
				lights = append(lights, true)
			default:
				panic("'.' or '#' expected")
			}
		}
	}
	return lights
}

// parse a string like "(1,2,3)" or "{1,2,3}"
func ParseListOfInt(word string, start, end byte) []int {
	values := make([]int, 0)
	last := len(word) - 1
	if word[0] != start || word[last] != end {
		panic("start/end charcater missing")
	}

	rawWord := word[1:last]
	numFields := strings.Split(rawWord, ",")
	for _, numField := range numFields {
		value, err := strconv.Atoi(numField)
		if err != nil {
			panic(err)
		}
		values = append(values, value)
	}
	return values
}
