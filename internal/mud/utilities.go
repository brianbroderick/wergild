package mud

import (
	"math/rand"
	"strings"
	"time"
)

/**
 * Stuff that should probably exist in Go (in my opinion)
 */

func containsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func random_int(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func formatToWidth(msg string, lineLen int) string {
	words := strings.Fields(msg)
	var slice []string
	curLine := 0

	for _, word := range words {
		slice = append(slice, word+" ")

		curLine += len(word) + 1
		if curLine >= lineLen {
			slice = append(slice, "\n")
			curLine = 0
		}
	}

	return strings.TrimSpace(strings.Join(slice, ""))
}
