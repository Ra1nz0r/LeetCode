package group_anagrams

import (
	"sort"
	"strings"
)

func GroupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string, len(strs))

	for _, str := range strs {
		chars := strings.Split(str, "")
		sort.Strings(chars)

		key := strings.Join(chars, "")
		groups[key] = append(groups[key], str)
	}

	result := make([][]string, 0, len(groups))

	for _, group := range groups {
		result = append(result, group)
	}

	return result
}
