package fuzzy

import "testing"

func TestMatch(t *testing.T) {
	b := "Grumpy wizards make toxic brew for the evil queen and Jack."
	if q := "GryJ."; !Match(b, q) {
		t.Errorf("Match() input `%v` should match query `%v`", b, q)
	}

	// Ensure UTF8 text works
	// Hindi:
	b = "हिन्दी"
	q := "हन्द"
	if !Match(b, q) {
		t.Errorf("Match() input `%v` should match query `%v` in both case sensitive and insensitive mode", b, q)
	}
	// Russian:
	b = "Cука блять"
	if q := "Cабят"; !Match(b, q) {
		t.Errorf("Match() input `%v` should match query `%v` (case sensitive)", b, q)
	}
}
