package main

import (
	"fmt"
	"sort"
)

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

func rotate(nums []int, k int) {
	arr := nums[k:]
	nums = nums[:k]
	for i := k; i < len(nums); i++ {
		arr = append(arr, nums[i])
	}
	fmt.Printf("", nums, arr)
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	stack := make([]rune, 0)
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, v := range s {
		switch v {
		case '(', '[', '{':
			stack = append(stack, v)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[v] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	if len(stack) == 0 {
		return true
	}

	return false
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := make(map[rune]int)
	tMap := make(map[rune]int)
	for _, v := range s {
		sMap[v]++
	}
	for _, v := range t {
		tMap[v]++
	}
	if len(sMap) != len(tMap) {
		return false
	}
	for k, v := range sMap {
		if v != tMap[k] {
			return false
		}
	}
	return true
}

func myPow(x float64, n int) float64 {
	if n == 1 {
		return x
	} else if n == 0 {
		return 1
	} else if n < 0 {
		x = 1 / x
		n *= -1
	}

	res := myPow(x, n/2)

	if n%2 == 0 {
		return res * res
	} else {
		return res * res * x
	}
}

func subsets(nums []int) [][]int {
	arr := [][]int{}
	arr = append(arr, []int{})

	for i := 0; i < len(nums); i++ {
		for _, v := range arr{
			tmp := make([]int, len(v))
			copy(tmp, v)
			tmp = append(tmp, nums[i])
			fmt.Println("test = ",tmp)
			arr = append(arr, tmp)
			fmt.Println("arr = ",arr)
		}
	}
	return arr
}

func main() {
	//num := []int{0,0,0}
	//print(threeSum(num))

	//num := []int{1,2,3}
	//test := plusOne(num)
	//fmt.Printf("", test)

	//num := []int{1,2,3,4,5,6,7}
	//rotate(num,3)
	//fmt.Printf("",num)

	//s := "(("
	//fmt.Printf("", isValid(s))
	nums := []int{0,3,5,7,9}
	fmt.Println("end ", subsets(nums))
}
