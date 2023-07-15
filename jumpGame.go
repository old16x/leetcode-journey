// https://leetcode.com/problems/jump-game/
package main

func canJump(nums []int) bool {
	l := len(nums)
	if l == 0 {
		return false
	}
	goal := l - 1
	for i := l - 1; i >= 0; i-- {
		if i+nums[i] >= goal {
			goal = i
		}
	}
	if goal == 0 {
		return true
	}
	return false
}
