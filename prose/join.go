package prose

import "strings"

// JoinWithCommas Accept a slice of strings and return string joined with commas
func JoinWithCommas(phrases []string) string {
	// join every phrase except the last one with commas
	result := strings.Join(phrases[:len(phrases)-1], ", ")
	result += ", and "
	// add the last phrase
	result += phrases[len(phrases)-1]
	return result
}
