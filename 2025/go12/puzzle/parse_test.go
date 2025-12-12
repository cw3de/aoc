package puzzle_test

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegexp(t *testing.T) {
	// "0:"
	reIndex := regexp.MustCompile(`^(\d):$`)
	maIndex := reIndex.FindStringSubmatch("3:")
	assert.NotNil(t, maIndex)
	assert.Equal(t, 2, len(maIndex))
	assert.Equal(t, "3", maIndex[1])

	// "###"
	reShape := regexp.MustCompile(`^[#\.]{3}$`)
	maShape := reShape.FindStringSubmatch("##.")
	assert.NotNil(t, maShape)

	// "12x34: 12 34 56 78 90 12"
	reRegion := regexp.MustCompile(`^(\d+)x(\d+): (\d+) (\d+) (\d+) (\d+) (\d+) (\d+)`)
	maRegion := reRegion.FindStringSubmatch("12x23: 34 45 56 67 78 89")
	assert.NotNil(t, maRegion)
	assert.Equal(t, 9, len(maRegion))
	assert.Equal(t, "12", maRegion[1])
	assert.Equal(t, "23", maRegion[2])
	assert.Equal(t, "34", maRegion[3])
	assert.Equal(t, "45", maRegion[4])
	assert.Equal(t, "56", maRegion[5])
	assert.Equal(t, "67", maRegion[6])
	assert.Equal(t, "78", maRegion[7])
	assert.Equal(t, "89", maRegion[8])
}
