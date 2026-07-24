package solution

import "testing"

// Go のテーブル駆動テスト。以降の問題もこの形をコピーして使う。
func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"重複あり", []int{1, 2, 3, 1}, true},
		{"重複なし", []int{1, 2, 3, 4}, false},
		{"複数種類が重複", []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
		{"要素1つ", []int{7}, false},
		{"空", []int{}, false},
		{"負数を含む", []int{-1, 0, 1, -1}, true},
	}

	for _, tt := range tests {
		// t.Run でサブテストにすると、どのケースが落ちたか名前で分かる
		t.Run(tt.name, func(t *testing.T) {
			got := containsDuplicate(tt.nums)
			if got != tt.want {
				t.Errorf("containsDuplicate(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

// ベンチマークは任意。go test -bench=. で走る。
func BenchmarkContainsDuplicate(b *testing.B) {
	nums := make([]int, 10000)
	for i := range nums {
		nums[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		containsDuplicate(nums)
	}
}
