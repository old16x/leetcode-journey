package main

import (
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3
	expectedResult := 2
	expectedNums := []int{2, 2}

	result := removeElement(nums, val)
	if result != expectedResult {
		t.Errorf("removeElement(nums, %d) = %d; want %d", val, result, expectedResult)
	}

	if !reflect.DeepEqual(nums[:result], expectedNums) {
		t.Errorf("nums = %v; want %v", nums[:result], expectedNums)
	}
}
