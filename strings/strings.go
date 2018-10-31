package strings

import "strings"

func Uppercase(stringToUppercase string) string {
	uppercased := strings.ToTitle(stringToUppercase)
	return uppercased
}