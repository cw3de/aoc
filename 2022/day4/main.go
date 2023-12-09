package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ReadLines(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data), "\n")
}

type Range struct {
	from int
	to   int
}

func (r Range) contains(o Range) bool {
	return r.from >= o.from && r.to <= o.to
}

func (r Range) overlaps(o Range) bool {
	// 123......
	// ......789
	if r.to < o.from || r.from > o.to {
		return false
	}
	return true
}

type Sections struct {
	e1 Range
	e2 Range
}

func parseSections(line string) Sections {
	re := regexp.MustCompile(`^(\d+)-(\d+),(\d+)-(\d+)$`)
	m := re.FindStringSubmatch(line)
	if m == nil {
		panic("no match in: " + line)
	}
	var s Sections
	var err error
	s.e1.from, err = strconv.Atoi(m[1])
	if err != nil {
		panic(err)
	}
	s.e1.to, err = strconv.Atoi(m[2])
	if err != nil {
		panic(err)
	}
	s.e2.from, err = strconv.Atoi(m[3])
	if err != nil {
		panic(err)
	}
	s.e2.to, err = strconv.Atoi(m[4])
	if err != nil {
		panic(err)
	}
	return s
}

func task1(filename string) {
	lines := ReadLines(filename)
	result := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		s := parseSections(line)
		if s.e1.contains(s.e2) || s.e2.contains(s.e1) {
			result++
		}
	}
	fmt.Println(filename, "task 1:", result)
}

func task2(filename string) {
	lines := ReadLines(filename)
	result := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		s := parseSections(line)
		if s.e1.overlaps(s.e2) {
			result++
		}
	}
	fmt.Println(filename, "task 2:", result)
}

func main() {
	task1("sample.txt")
	task1("input.txt")
	task2("sample.txt")
	task2("input.txt")
}
