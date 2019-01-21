package twofer

import "strings"

func ShareWith(name string) string {
	defaultText := "One for you, one for me."
	if name == "" {
		return defaultText
	}
	return strings.Replace(defaultText, "you", name, -1)
}
