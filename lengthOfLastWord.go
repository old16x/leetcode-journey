// https://leetcode.com/problems/length-of-last-word/
package main

func lengthOfLastWord(s string) int {
	l := len(s) - 1
	count := 0
	for l >= 0 && s[l] == ' ' {
		l--
	}
	for l >= 0 && s[l] != ' ' {
		l--
		count++
	}
	return count
}
