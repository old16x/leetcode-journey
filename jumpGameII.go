// https://leetcode.com/problems/jump-game-ii/
package main

func jump(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	endOfRange := 0
	farthest := 0
	minJumps := 0

	for i := 0; i < n-1; i++ {
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}
		if i == endOfRange {
			minJumps++
			endOfRange = farthest
		}
	}
	return minJumps
}
