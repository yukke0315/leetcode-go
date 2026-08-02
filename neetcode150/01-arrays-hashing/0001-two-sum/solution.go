package solution

// 1. Two Sum (Easy)
// https://leetcode.com/problems/two-sum/
//
// 計算量: O(?) time / O(?) space
func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)
	for i, v := range nums {
		gap := target - v
		if j, ok := seen[gap]; ok {
			return []int{i, j}
		}
		seen[v] = i
	}
	return nil
}