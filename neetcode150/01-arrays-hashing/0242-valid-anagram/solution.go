package solution

// 242. Valid Anagram (Easy)
// https://leetcode.com/problems/valid-anagram/
//
// 計算量: O(n) time / O(n) space
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := make(map[rune]int)
	for _, n := range s {
		count[n]++
	}

	for _, n := range t {
		count[n]--
	}

	for _, v := range count {
		if v != 0 {
			return false
		}
	}
	return true
}