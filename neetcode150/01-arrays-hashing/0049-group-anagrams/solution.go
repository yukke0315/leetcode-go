package solution

// 49. Group Anagrams (Medium)
// https://leetcode.com/problems/group-anagrams/
//
// 計算量: O(n*klogk) time / O(n*k) space
import "slices"

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)

	for _, s := range strs {
		b := []byte(s)
		slices.Sort(b)
		key := string(b)
		anagrams[key] = append(anagrams[key], s)
	}

	results := make([][]string, 0, len(anagrams))

	for _, v := range anagrams {
		results = append(results, v)
	}

	return results
}
