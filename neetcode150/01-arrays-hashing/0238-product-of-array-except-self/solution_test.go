package solution

import (
	"slices"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"例1", []int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{"例2", []int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
		{"長さ2", []int{2, 3}, []int{3, 2}},
		{"負の数だけ", []int{-2, -3, -4}, []int{12, 8, 6}},
		{"0 が先頭", []int{0, 4, 5}, []int{20, 0, 0}},
		{"0 が末尾", []int{4, 5, 0}, []int{0, 0, 20}},
		{"0 が2つ", []int{0, 2, 0, 3}, []int{0, 0, 0, 0}},
		{"同じ値の繰り返し", []int{2, 2, 2, 2}, []int{8, 8, 8, 8}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := productExceptSelf(slices.Clone(tt.nums))
			if !slices.Equal(got, tt.want) {
				t.Errorf("productExceptSelf(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
