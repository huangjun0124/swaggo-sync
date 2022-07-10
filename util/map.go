package util

import "strings"

func MapKeysToSlice(m map[string]bool) []string {
	s := make([]string, 0, len(m))
	for k := range m {
		if len(strings.TrimSpace(k)) > 0 {
			s = append(s, k)
		}
	}
	return s
}
