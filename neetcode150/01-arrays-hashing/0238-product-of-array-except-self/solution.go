package solution

// 238. Product Of Array Except Self (Medium)
// https://leetcode.com/problems/product-of-array-except-self/
//
// 計算量: O(n) time / O(n) space
func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))
	left := make([]int, len(nums))
	right := make([]int, len(nums))
	left[0] = 1
	right[len(nums)-1] = 1

	for i := 1; i < len(nums); i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	for i := range nums {
		answer[i] = left[i] * right[i]
	}

	return answer
}
