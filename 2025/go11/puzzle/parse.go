package puzzle

import (
	"fmt"
	"os"
	"strings"
)

type Rack struct {
	Input   int
	Outputs []int
}

type Puzzle struct {
	Name2Num map[string]int
	Num2Name map[int]string
	Racks    map[int]Rack
}

func Load(filename string) *Puzzle {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return Parse(string(data))
}

func Parse(text string) *Puzzle {
	p := &Puzzle{
		Name2Num: make(map[string]int),
		Num2Name: make(map[int]string),
		Racks:    make(map[int]Rack, 0),
	}

	lines := strings.Split(text, "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		fields := strings.Split(line, " ")
		if len(fields) < 2 {
			panic("at least two fields expected")
		}

		var r Rack

		for i, name := range fields {
			if i == 0 {
				inputName := name[:len(name)-1]
				r.Input = p.FindDevice(inputName)
			} else {
				r.Outputs = append(r.Outputs, p.FindDevice(name))
			}
		}
		p.Racks[r.Input] = r
	}

	fmt.Printf("parsed %d racks with %d devices\n", len(p.Racks), len(p.Name2Num))

	return p
}

func (p *Puzzle) MustFind(name string) int {
	num, found := p.Name2Num[name]
	if found {
		return num
	}
	panic(fmt.Errorf("rack '%s' not found", name))
}

func (p *Puzzle) FindDevice(name string) int {

	num, found := p.Name2Num[name]
	if found {
		return num
	}
	num = len(p.Name2Num)
	p.Name2Num[name] = num
	p.Num2Name[num] = name
	// fmt.Printf("%s -> %d\n", name, num)
	return num
}

func (p *Puzzle) Name(num int) string {
	return p.Num2Name[num]
}

func (p *Puzzle) Show() {
	fmt.Printf("--------------------\n")
	for i := 0; i < len(p.Racks); i++ {
		fmt.Printf("%s(%d):", p.Name(i), i)

		for _, o := range p.Racks[i].Outputs {
			fmt.Printf(" %s(%d)", p.Name(o), o)
		}
		fmt.Printf("\n")
	}
}
