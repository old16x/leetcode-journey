// https://leetcode.com/problems/two-sum/
package main

func twoSum(nums []int, target int) []int {
	tmp := make(map[int]int)

	for index, value := range nums {
		diff := target - value
		if idx, ok := tmp[diff]; ok {
			return []int{idx, index}
		}
		tmp[value] = index
	}
	return []int{}
}
