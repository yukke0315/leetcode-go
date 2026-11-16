package solution

import (
	"slices"
	"strings"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		{"例1",
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
		{"空文字1つ", []string{""}, [][]string{{""}}},
		{"1文字", []string{"a"}, [][]string{{"a"}}},
		{"同じ文字列の重複", []string{"ab", "ba", "ab"}, [][]string{{"ab", "ab", "ba"}}},
		{"全部別グループ", []string{"abc", "abd", "xyz"}, [][]string{{"abc"}, {"abd"}, {"xyz"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupAnagrams(tt.strs)
			if !slices.EqualFunc(normalize(got), normalize(tt.want), slices.Equal) {
				t.Errorf("groupAnagrams(%q) = %q, want %q", tt.strs, got, tt.want)
			}
		})
	}
}

// グループの順番もグループ内の順番も問わないので、両方そろえてから比べる。
// 元のスライスを書き換えないよう、コピーしてから並べ替える。
func normalize(groups [][]string) [][]string {
	out := make([][]string, 0, len(groups))
	for _, g := range groups {
		c := slices.Clone(g)
		slices.Sort(c)
		out = append(out, c)
	}
	slices.SortFunc(out, func(a, b []string) int {
		return strings.Compare(strings.Join(a, ","), strings.Join(b, ","))
	})
	return out
}
