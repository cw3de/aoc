package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IdRange struct {
	Start int
	End   int
}

func ReadLines(filename string) []IdRange {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	idRanges := make([]IdRange, 0)

	data = data[:len(data)-1]

	for _, part := range strings.Split(string(data), ",") {
		pair := strings.Split(part, "-")
		if len(pair) != 2 {
			panic("invalid pair")
		}
		start, err := strconv.Atoi(pair[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(pair[1])
		if err != nil {
			panic(err)
		}
		idRanges = append(idRanges, IdRange{
			Start: start,
			End:   end,
		})
	}
	return idRanges
}

func isSillyId(id int) bool {
	txt := strconv.Itoa(id)
	fullLen := len(txt)
	if fullLen%2 != 0 {
		return false
	}
	halfLen := fullLen / 2

	return txt[:halfLen] == txt[halfLen:]
}

func task1(filename string) {
	idRanges := ReadLines(filename)
	result := 0
	for _, idRange := range idRanges {
		for id := idRange.Start; id <= idRange.End; id++ {
			if isSillyId(id) {
				result += id
			}
		}
	}
	fmt.Println(filename, "task 1:", result)
}

func isPattern(txt string, patLen int) bool {
	fullLen := len(txt)

	if 2*patLen > fullLen {
		return false
	}

	if fullLen%patLen != 0 {
		return false
	}

	pattern := txt[:patLen]
	for pos := patLen; pos < fullLen; pos += patLen {
		if txt[pos:pos+patLen] != pattern {
			return false
		}
	}
	return true
}

func isVerySillyId(id int) bool {
	txt := strconv.Itoa(id)
	fullLen := len(txt)
	halfLen := fullLen / 2

	for patLen := 1; patLen <= halfLen; patLen++ {
		if isPattern(txt, patLen) {
			return true
		}
	}

	return false
}

func task2(filename string) {
	idRanges := ReadLines(filename)
	result := 0
	for _, idRange := range idRanges {
		fmt.Printf("%d - %d (%d)\n",
			idRange.Start,
			idRange.End,
			idRange.End-idRange.Start)
		for id := idRange.Start; id <= idRange.End; id++ {
			if isVerySillyId(id) {
				fmt.Printf("   %d\n", id)
				result += id
			}
			//
		}
	}
	fmt.Println(filename, "task 2:", result)
}

func main() {
	// task1("sample.txt")
	task1("input.txt")
	// task2("sample.txt")
	task2("input.txt")
}
