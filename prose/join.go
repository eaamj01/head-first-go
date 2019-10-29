package prose

import "strings"

// JoinWithCommas Accept a slice of strings and return string joined with commas
func JoinWithCommas(phrases []string) string {
	switch len(phrases) {
	case 0:
		return ""
	case 1:
		return phrases[0]

	case 2:
		return phrases[0] + " and " + phrases[1]

	default:

	}
	// join every phrase except the last one with commas
	result := strings.Join(phrases[:len(phrases)-1], ", ")
	result += ", and "
	// add the last phrase
	result += phrases[len(phrases)-1]
	return result
}
