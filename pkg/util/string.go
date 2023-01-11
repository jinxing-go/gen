package util

import (
	"strings"
)

func Includes[T comparable](value T, searchValue []T) bool {
	for _, v := range searchValue {
		if v == value {
			return true
		}
	}

	return false
}

// Studly user_id to UserId
func Studly(name string) string {
	newName := strings.Builder{}
	upChar := true
	for _, chr := range name {
		if upChar {
			if chr >= 'a' && chr <= 'z' {
				chr -= 'a' - 'A'
			}

			newName.WriteRune(chr)
			upChar = false
			continue
		}

		if chr == '_' || chr == ' ' {
			upChar = true
			continue
		}

		newName.WriteRune(chr)
	}

	return newName.String()
}
