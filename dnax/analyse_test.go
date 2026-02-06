package main

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnalyse(t *testing.T) {
	file, err := os.ReadFile("testseq.txt")
	if err != nil {
		return
	}
	mymap, err2 := analyseDNASeq(file)
	if err2 != nil {
		return
	}

	gcPercentage := 0.5
	baseCount := 8
	assert.Equal(t, strconv.Itoa(baseCount), mymap["Base Count"], "Base count should be 8")

	assert.Equal(t, strconv.FormatFloat(gcPercentage, 'E', -1, 64),
		mymap["GC_percentage"], "GC percentage should be 0.5")

	assert.Equal(t, "CGCGTATA", mymap["Complementary Sequence"],
		"Complement sequence should be CGCGTATA")

}
