package solution

// 49. Group Anagrams (Medium)
// https://leetcode.com/problems/group-anagrams/
//
// 計算量: O(n*klogk) time / O(n*k) space
import "slices"

func groupAnagrams(strs []string) [][]string {
    groups := make(map[string][]string)

	for _, s := range strs {
		b := []byte(s)
		slices.Sort(b)
		key := string(b)
		groups[key] = append(groups[key], s)
	}
	result := make([][]string, 0, len(groups))
	for _, g := range groups {
		result = append(result, g)
	}
	return result
}
