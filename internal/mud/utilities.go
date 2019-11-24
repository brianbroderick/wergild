package mud

import (
	"math/rand"
	"regexp"
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

// format to a certain number of characters for telnet happiness.
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

func ToSlug(str string) string {
	reAn := regexp.MustCompile(`[^a-z0-9]`)
	reSp := regexp.MustCompile(`\s+`)

	str = strings.ToLower(str)
	str = reAn.ReplaceAllString(str, " ")
	str = strings.TrimSpace(str)
	str = reSp.ReplaceAllString(str, "_")

	return str
}
