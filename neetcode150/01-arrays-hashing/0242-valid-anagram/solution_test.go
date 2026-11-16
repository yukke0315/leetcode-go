package solution

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		s, t string
		want bool
	}{
		{"例1", "anagram", "nagaram", true},
		{"例2", "rat", "car", false},
		{"長さが違う", "ab", "abc", false},
		{"文字の種類は同じで回数が違う", "aab", "abb", false},
		{"1文字", "a", "a", true},
		{"同じ文字列", "listen", "listen", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isAnagram(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("isAnagram(%q, %q) = %v, want %v", tt.s, tt.t, got, tt.want)
			}
		})
	}
}
