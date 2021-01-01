package main

import "strings"

func Words(s string) string {
	length := len(strings.Fields(s))
	switch {
	case length == 0:
		return "wordless"
	case length == 1:
		return "one word"
	case length < 4:
		return "a few words"
	case length < 8:
		return "many words"
	default:
		return "too many words"
	}
}
