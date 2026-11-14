package solution

// 217. Contains Duplicate (Easy)
// https://leetcode.com/problems/contains-duplicate/
//
// 計算量: O(n) time / O(n) space
func containsDuplicate(nums []int) bool {
	q := make(map[int]int)

	for _, v := range nums {
		q[v] += 1
	}

	for _, cnt := range q {
		if cnt >= 2 {
			return true
		}
	}

	return false
}
