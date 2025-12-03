package bank

import (
	"strconv"
)

func searchHighestDiget(text string, startPos, endPos int) int {
	bestPos := startPos
	for pos := startPos + 1; pos <= endPos; pos++ {
		if text[pos] > text[bestPos] {
			bestPos = pos
		}
	}
	return bestPos
}

func FindBest(line string, numberOfDigits int) int64 {
	startPos := 0
	for len(line) > numberOfDigits && startPos < numberOfDigits {
		// search the highest digit after startPos
		// at least (numberOfDigits - startPos) digits must remain
		last := len(line) + startPos - numberOfDigits
		nextPos := searchHighestDiget(line, startPos, last)
		// remove digits between startPos and nextPos
		line = line[:startPos] + line[nextPos:]
		startPos++
	}
	// if we didn't find enough digits to remove, truncate at the end
	line = line[:numberOfDigits]

	value, err := strconv.ParseInt(line, 10, 64)
	if err != nil {
		panic(err)
	}
	return value
}
