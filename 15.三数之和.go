/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	ans := make([][]int, 0)

	if len(nums) < 3 {
		return ans
	}
	sort.Ints(nums)

	for k := 0; k < len(nums); k++ {
		if nums[k] > 0 {
			break
		}

		if k > 0 && nums[k] == nums[k-1] {
			continue
		}

		for i, j := k+1, len(nums)-1; i < j; {

			if j < len(nums)-1 && nums[j] == nums[j+1] {
				j--
				continue
			}

			if nums[k]+nums[i]+nums[j] < 0 {
				i++
			} else if nums[k]+nums[i]+nums[j] > 0 {
				j--
			} else {
				ans = append(ans, []int{nums[k], nums[i], nums[j]})
				i++
				j--
			}
		}
	}

	return ans
}

// @lc code=end

