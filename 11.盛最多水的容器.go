package main

/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	max := 0
	len := len(height)
	for i, j := 0, len-1; i < j; {
		num := 0
		if height[i] > height[j] {
			num = height[j]
			j--
		} else {
			num = height[i]
			i++
		}

		count := num * (j - i + 1)
		if max < count {
			max = count
		}

	}
	return max
}

// func main() {
// 	test := []int{1, 1}

// 	num := maxArea(test)
// 	fmt.Printf("%d", num)
// }

// @lc code=end
