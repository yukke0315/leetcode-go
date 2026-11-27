package solution

// 1. Two Sum (Easy)
// https://leetcode.com/problems/two-sum/
//
// 計算量: O(n) time / O(n) space
func twoSum(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))

	for i, n := range nums {
		gap := target - n
		if j, ok := seen[gap]; ok {
			return []int{i, j}
		}

		seen[n] = i
	}

	return nil
}
