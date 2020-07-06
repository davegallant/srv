package utils

import (
	"regexp"
	"sort"
	"strings"
)

// StripHTMLTags uses regex to strip all html elements
func StripHTMLTags(s string) string {
	const pattern = `(<\/?[a-zA-A]+?[^>]*\/?>)*`
	r := regexp.MustCompile(pattern)
	groups := r.FindAllString(s, -1)
	sort.Slice(groups, func(i, j int) bool {
		return len(groups[i]) > len(groups[j])
	})
	for _, group := range groups {
		if strings.TrimSpace(group) != "" {
			s = strings.ReplaceAll(s, group, "")
		}
	}
	return s
}
