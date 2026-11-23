package solution

// 242. Valid Anagram (Easy)
// https://leetcode.com/problems/valid-anagram/
//
// 計算量: O(n) time / O(n) space
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	cnt := make(map[rune]int, 26)

	for _, c := range s {
		cnt[c]++
	}

	for _, c := range t {
		cnt[c]--

		if cnt[c] < 0 {
			return false
		}
	}
	return true
}
