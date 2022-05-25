package main

import "fmt"

//283. 移动零
func moveZeroes1(nums []int) {
	var n int
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[n] = nums[i]
			n++
		}

	}
	for {
		if n < len(nums) {
			nums[n] = 0
			n++
		} else {
			break
		}

	}
}
func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes1(nums)
	fmt.Println(nums)
}
