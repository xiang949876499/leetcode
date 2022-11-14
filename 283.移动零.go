

/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {
	i := 0
	count := 0
	for k, v := range nums {
		if v == 0 {
			count++
		} else {
			nums[i] = nums[k]
			i++
		}
	}
	nums = nums[:i]
	for i := 0; i < count; i++ {
		nums = append(nums, 0)
	}

}

// @lc code=end
