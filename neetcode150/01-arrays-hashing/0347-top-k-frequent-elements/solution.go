package solution

// 347. Top K Frequent Elements (Medium)
// https://leetcode.com/problems/top-k-frequent-elements/
//
// 計算量: O(?) time / O(?) space
func topKFrequent(nums []int, k int) []int {
	cnt := make(map[int]int)

	for _, n := range nums {
		cnt[n]++
	}

	s := make([][]int, len(nums)+1)

	for n, v := range cnt {
		s[v] = append(s[v], n)
	}

	result := make([]int, 0, k)
	for i := len(nums); i >= 0; i-- {
		for _, c := range s[i] {
			result = append(result, c)

			if len(result) == k {
				return result
			}
		}
	}

	return nil
}
