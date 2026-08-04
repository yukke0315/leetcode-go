package solution

// 49. Group Anagrams (Medium)
// https://leetcode.com/problems/group-anagrams/
//
// 計算量: O(n*klogk) time / O(n*k) space
import "slices"

// func groupAnagrams(strs []string) [][]string {
// 	groups := make(map[string][]string)

// 	for _, s := range strs {
// 		b := []byte(s)
// 		slices.Sort(b)
// 		key := string(b)
// 		groups[key] = append(groups[key], s)
// 	}
// 	result := make([][]string, 0, len(groups))
// 	for _, g := range groups {
// 		result = append(result, g)
// 	}
// 	return result
// }
func groupAnagrams(strs []string) [][]string {
	arr := make(map[string][]string)

	// keyとするために要素の中身をアルファベット順に
	for _, v := range strs {
		b := []byte(v)   // stringのままでは要素の中身を並び替えられない
		slices.Sort(b)
		key := string(b)
		arr[key] = append(arr[key], v) 
	}
	result := make([][]string, 0, len(arr))   // 結果入れ
	for _, w := range arr {
		result = append(result, w)
	}
	return result
}