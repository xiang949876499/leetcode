/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	// if len(s) != len(t) {
	// 	return false
	// }
	// sMap := make(map[rune]int)
	// tMap := make(map[rune]int)
	// for _, v := range s {
	// 	sMap[v]++
	// }
	// for _, v := range t {
	// 	tMap[v]++
	// }
	// if len(sMap) != len(tMap) {
	// 	return false
	// }
	// for k, v := range sMap {
	// 	if v != tMap[k] {
	// 		return false
	// 	}
	// }
	// return true

	if len(s) != len(t) {
		return false
	}
	cnt := map[rune]int{}
	for _, ch := range s {
		cnt[ch]++
	}
	for _, ch := range t {
		cnt[ch]--
		if cnt[ch] < 0 {
			return false
		}
	}
	return true
}

// @lc code=end

