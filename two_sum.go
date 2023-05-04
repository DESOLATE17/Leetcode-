package main

import "fmt"

// https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int { //O(n^2 / 2)
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			if v+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	pairs := make(map[int]int, len(nums))

	for i, v := range nums {
		if ind, ok := pairs[target-v]; ok {
			return []int{ind, i}
		} else {
			pairs[v] = i
		}
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum2([]int{2, 7, 11, 15}, 9))
}
