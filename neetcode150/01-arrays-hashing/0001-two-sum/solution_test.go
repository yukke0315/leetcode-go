package solution

import (
	"slices"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"例1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"例2", []int{3, 2, 4}, 6, []int{1, 2}},
		{"同じ値2つ", []int{3, 3}, 6, []int{0, 1}},
		{"負数とゼロ", []int{-3, 4, 3, 90}, 0, []int{0, 2}},
		{"答えに絡まない重複", []int{5, 1, 9, 1, 6}, 15, []int{2, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)
			// 返す順番は問わないので、並べ替えてから比べる
			sorted := slices.Clone(got)
			slices.Sort(sorted)
			if !slices.Equal(sorted, tt.want) {
				t.Errorf("twoSum(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
