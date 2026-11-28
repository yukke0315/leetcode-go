package solution

// 217. Contains Duplicate (Easy)
// https://leetcode.com/problems/contains-duplicate/
//
// 計算量: O(n) time / O(n) space
func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool, len(nums))

	for _, n := range nums {
		if seen[n] {
			return true
		}

		seen[n] = true
	}

	return false
}
