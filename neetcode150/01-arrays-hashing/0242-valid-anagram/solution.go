package solution

// 242. Valid Anagram (Easy)
// https://leetcode.com/problems/valid-anagram/
//
// 計算量: O(n) time / O(n) space
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	ana := make(map[rune]int)
	for _, v := range s {
		ana[v]++
	}

	for _, v := range t {
		ana[v]--
	}

	for _, v := range ana {
		if v != 0 {
			return false
		}
	}
	return true
}
