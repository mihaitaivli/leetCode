package twoSum2SortedArray

import (
"reflect"
"testing"
)

func TestTwoSum(t *testing.T) {
	//Constraints:
	//	2 <= nums.length <= 3 * 104
	//	-1000 <= nums[i] <= 1000
	//	nums is sorted in increasing order.
	//	-1000 <= target <= 1000


	t.Run("Array has 2 elements", func(t *testing.T) {
		arr := []int{1,2}
		target := 3
		expected := []int{0,1}
		response := TwoSumII(arr, target)

		if !reflect.DeepEqual(expected, response) {
			t.Errorf("Expected %v, but got %v\n", expected, response)
		}
	})
	t.Run("Array has more than 2 elements", func(t *testing.T) {
		arr := []int{1,2,4,12,15}
		target := 19
		expected := []int{2,4}
		response := TwoSumII(arr, target)

		if !reflect.DeepEqual(expected, response) {
			t.Errorf("Expected %v, but got %v\n", expected, response)
		}
	})
	t.Run("The target is a negative number", func(t *testing.T) {
		arr := []int{-15,-12,0,1,2}
		target := -13
		expected := []int{0,4}
		response := TwoSumII(arr, target)

		if !reflect.DeepEqual(expected, response) {
			t.Errorf("Expected %v, but got %v\n", expected, response)
		}
	})
}

