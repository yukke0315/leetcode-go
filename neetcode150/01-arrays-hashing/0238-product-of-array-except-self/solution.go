package solution

// 238. Product Of Array Except Self (Medium)
// https://leetcode.com/problems/product-of-array-except-self/
//
// 計算量: O(n) time / O(n) space
func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))
	bfr := make([]int, len(nums))
	bfr[0] = 1
	aft := make([]int, len(nums))
	aft[len(nums)-1] = 1

	for i := 1; i < len(nums); i++ {
		bfr[i] = bfr[i-1] * nums[i-1]
	}

	for i := len(nums) - 2; i >= 0; i-- {
		aft[i] = aft[i+1] * nums[i+1]
	}

	for i := range nums {
		answer[i] = bfr[i] * aft[i]
	}

	return answer
}
