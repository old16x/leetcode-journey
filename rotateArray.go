// https://leetcode.com/problems/rotate-array/
package main

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n

	tmp := append(nums[n-k:], nums[:n-k]...)
	copy(nums, tmp)
}
