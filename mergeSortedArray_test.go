package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	merge(nums1, m, nums2, n)

	expected := []int{1, 2, 2, 3, 5, 6}
	if !reflect.DeepEqual(nums1, expected) {
		t.Errorf("Merge failed. Expected %v, but got %v", expected, nums1)
	}
}
