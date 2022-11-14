/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	digits[len(digits)-1]++
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] >= 10 {
			digits[i] = 0
			if i > 0 {
				digits[i-1]++
			} else {

				digits = make([]int, len(digits)+1)
				digits[0] = 1
			}
		}

	}
	return digits
}

// @lc code=end

