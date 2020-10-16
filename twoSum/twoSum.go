package twoSum

func TwoSum(nums []int, target int) []int {
	outcome := make([]int, 2)
	var shouldBreak bool = false
	for i, firstVal := range nums {
		for j,secondVal := range nums[i + 1:] {
			if firstVal + secondVal == target {
				outcome[0] = i
				outcome[1] = i + j + 1

				shouldBreak = true
				break
			}
		}
		if shouldBreak == true {
			break
		}
	}

	return outcome
}