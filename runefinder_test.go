package runefinder

import "testing"

func TestSimpleParse(t *testing.T) {
	line := "0041;LATIN CAPITAL LETTER A;Lu;0;L;;;;;N;;;;0061;"
	char, name := Parse(line)
	the_rune := 'A'
	if char != the_rune {
		t.Error("Expected", the_rune, "got", char)
	}
	the_name := "LATIN CAPITAL LETTER A"
	if name != the_name {
		t.Error("Expected", the_name, "got", name)
	}
}
