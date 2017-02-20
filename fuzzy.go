// Package fuzzy provides Sublime Text-like fuzzy text matching on strings
package fuzzy

import (
	"strings"
	"unicode/utf8"
)

// Match takes a blob (the string to be matched) and a query (the string to match with)
// and returns either a true or false value specifying whether the blob fuzzy matches
// the query.
func Match(blob, query string) bool {
	q := len(query)
	if b := len(blob); q == 0 {
		return true
	} else if q > b {
		return false
	}
	i := 0
	matches := 0
	for _, char := range blob {
		if i >= q {
			break
		}
		r, width := utf8.DecodeRuneInString(query[i:])
		if char == r {
			matches += width
			i += width
		}
	}
	return matches == q
}

// MatchFold is the same as Match except it is case-insensitive
func MatchFold(blob, query string) bool {
	return Match(strings.ToLower(blob), strings.ToLower(query))
}
