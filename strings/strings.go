package strings

import "strings"

func Uppercase(stringToUppercase string, trimWhiteSpace bool) string {
	uppercased := stringToUppercase
	if trimWhiteSpace {
		uppercased = strings.Trim(uppercased, " ")
	}

	return strings.ToTitle(uppercased)
}