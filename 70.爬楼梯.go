/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	if n < 2 {
		return 1
	}
	if n == 2 {
		return 2
	}
	j, k = 1, 2
	for i := 3; i < n; i++ {
		j, k = k, j+k
	}
}

// @lc code=end

