package twoSum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	//Constraints:
	//	2 <= nums.length <= 105
	//	-109 <= nums[i] <= 109
	//	-109 <= target <= 109
	//	Only one valid answer exists.

	t.Run("Array has 2 elements", func(t *testing.T) {
		arr := []int{1,2}
		target := 3
		expected := []int{0,1}
		response := TwoSum(arr, target)

		if !reflect.DeepEqual(expected, response) {
			t.Errorf("Expected %v, but got %v\n", expected, response)
		}
	})
	t.Run("Array has more than 2 elements", func(t *testing.T) {
		arr := []int{1,2,4,12,15}
		target := 19
		expected := []int{2,4}
		response := TwoSum(arr, target)

		if !reflect.DeepEqual(expected, response) {
			t.Errorf("Expected %v, but got %v\n", expected, response)
		}
	})
	t.Run("The target is a negative number", func(t *testing.T) {
		arr := []int{0,2,4,-12,-15}
		target := -10
		expected := []int{1,3}
		response := TwoSum(arr, target)

		if !reflect.DeepEqual(expected, response) {
			t.Errorf("Expected %v, but got %v\n", expected, response)
		}
	})
}
