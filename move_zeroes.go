package main

// https://leetcode.com/problems/move-zeroes/submissions/944548098/
//Input: nums = [0,1,0,3,12]
//
//  *                 *
// [0,1,0,3,12]    [0,1,0,3,12]     [0,1,0,3,12]
//  *               *                  *
//Output: [1,3,12,0,0]

func moveZeroes(nums []int) {
	zero := 0
	for i, v := range nums {
		if v != 0 {
			nums[i], nums[zero] = nums[zero], nums[i]
			zero++
		}
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
}
