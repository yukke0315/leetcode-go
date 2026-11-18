package solution

// 49. Group Anagrams (Medium)
// https://leetcode.com/problems/group-anagrams/
//
// 計算量: O(n*klogk) time / O(n*k) space
import "slices"

func groupAnagrams(strs []string) [][]string {
	anagram := make(map[string][]string)

	for _, s := range strs {
		a := []byte(s)
		slices.Sort(a)
		key := string(a)
		anagram[key] = append(anagram[key], s)
	}

	groups := [][]string{}

	for _, v := range anagram {
		groups = append(groups, v)
	}

	return groups
}
