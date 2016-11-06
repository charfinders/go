package runefinder

import (
	"fmt"
	"strconv"
	"strings"
)

func Parse(ucd_line string) (rune, string) {
	parts := strings.Split(ucd_line, ";")
	if len(parts) != 15 {
		panic("Wrong number of fields in UCD line")
	}
	char, err := strconv.ParseInt(parts[0], 16, 32)
	if err != nil {
		fmt.Printf("Parse: error parsing codepoint on line %v", ucd_line)
		panic(err)
	}
	return rune(char), parts[1]
}
