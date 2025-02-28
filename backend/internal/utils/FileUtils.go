package utils

import "strings"

func SanitizeFolderName(name string) string {
	name = strings.Map(func(r rune) rune {
		if r == '/' || r == '\\' || r == ':' || r == '*' || r == '?' || r == '"' || r == '<' || r == '>' || r == '|' {
			return -1
		}
		return r
	}, name)

	name = strings.TrimSpace(name)

	if name == "" {
		return "New Folder"
	}

	return name
}
