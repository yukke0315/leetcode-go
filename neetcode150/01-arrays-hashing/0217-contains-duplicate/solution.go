package solution

// 217. Contains Duplicate (Easy)
// https://leetcode.com/problems/contains-duplicate/
//
// 計算量: O(?) time / O(?) space
func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)
	
	for _, v := range nums {
		if seen[v] {
			return true
		}
		seen[v] = true
	}
	return false
}
