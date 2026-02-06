package main

import (
	"fmt"
	"io"
	"strings"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
)

var useColor = false

func analyseDNASeq(seq []byte) (results map[string]string, err error) {
	bases := map[byte]bool{'A': true, 'T': true, 'C': true, 'G': true}
	strSeq := strings.Replace(string(seq), " ", "", -1)

	countA := 0
	countT := 0
	countG := 0
	countC := 0

	results = make(map[string]string)

	for _, base := range seq {

		if !bases[base] || base == ' ' {
			return nil, fmt.Errorf("invalid base, %c found in sequence", base)
		}

		switch base {
		case 'A':
			countA++
		case 'C':
			countC++
		case 'T':
			countT++
		case 'G':
			countG++
		}
	}

	results["Base Count"] = fmt.Sprintf(
		"from %d --> %dA, %dT, %dG, %dC",
		len(seq), countA, countT, countG, countC,
	)

	countGC := countC + countG
	gcPercentage := (float64(countGC) / float64(len(seq))) * 100

	results["GC_percentage"] = fmt.Sprintf("%.2f%%", gcPercentage)

	results["Complementary Sequence"] = string(computeComplementSeq(strSeq))
	return
}

func computeComplementSeq(seq string) string {
	complementByteSeq := map[byte]byte{
		'A': 'T',
		'T': 'A',
		'G': 'C',
		'C': 'G',
	}

	result := make([]byte, len(seq))

	for i := 0; i < len(seq); i++ {
		result[len(seq)-i-1] = complementByteSeq[seq[i]]
	}

	return string(result)
}

func display(results map[string]string, w io.Writer) error {
	precedence := []string{
		"Base Count",
		"GC_percentage",
		"Complementary Sequence",
	}

	fmt.Fprintf(w, "=== DNA Sequence Results ===")
	fmt.Fprintln(w)

	for _, key := range precedence {
		value, exists := results[key]
		displayKey := strings.ReplaceAll(key, "_", " ")
		if !exists {
			continue
		}

		displayKey = strings.ReplaceAll(displayKey, "_", " ")
		if key == "Base Count" {
			value = colorize(value)
		} else if key == "Complementary Sequence" {
			value = colorize(value)
		}
		if _, err := fmt.Fprintf(w, " %s: %s\n", displayKey, value); err != nil {
			return err
		}
	}
	return nil
}

func colorize(s string) string {

	if !useColor {
		return s
	}
	var sb strings.Builder
	for _, base := range s {
		switch base {
		case 'A':
			sb.WriteString(colorGreen + "A" + colorReset)
		case 'T':
			sb.WriteString(colorYellow + "T" + colorReset)
		case 'G':
			sb.WriteString(colorRed + "G" + colorReset)
		case 'C':
			sb.WriteString(colorBlue + "C" + colorReset)
		default:
			sb.WriteRune(base)
		}
	}
	return sb.String()
}

func enableColor() {
	useColor = true
}

func disableColor() {
	useColor = false
}
