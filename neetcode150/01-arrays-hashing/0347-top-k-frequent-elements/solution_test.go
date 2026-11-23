package solution

import (
	"slices"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int // 昇順で書く（答えの順番は問わないため、got も並べ替えて比べる）
	}{
		{"例1", []int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{"例2", []int{1}, 1, []int{1}},
		{"例3", []int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}, 2, []int{1, 2}},
		{"負の数を含む", []int{-1, -1, 2}, 1, []int{-1}},
		{"出現順と回数の順が違う", []int{3, 0, 1, 0}, 1, []int{0}},
		{"多いものが後ろにある", []int{1, 2, 2, 3, 3, 3}, 2, []int{2, 3}},
		{"k が種類数と同じ", []int{4, 4, 5, 6, 6, 6}, 3, []int{4, 5, 6}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequent(slices.Clone(tt.nums), tt.k)
			sorted := slices.Clone(got)
			slices.Sort(sorted)
			if !slices.Equal(sorted, tt.want) {
				t.Errorf("topKFrequent(%v, %d) = %v, want %v（順不同）", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
