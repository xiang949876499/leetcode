/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 轮转数组
 */

// @lc code=start
func rotate(nums []int, k int) {
	arr := nums[:k]
	nums = nums[k:]
	for i := 0; i < k; i++ {
		nums = append(nums, arr[i])
	}
}

// @lc code=end

