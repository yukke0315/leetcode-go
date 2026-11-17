package solution

// 242. Valid Anagram (Easy)
// https://leetcode.com/problems/valid-anagram/
//
// 計算量: O(n) time / O(n) space
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	q := make(map[rune]int, 26)

	for _, c := range s {
		q[c]++
	}

	for _, c := range t {
		q[c]--
	}

	for _, n := range q {
		if n != 0 {
			return false
		}
	}
	return true
}
