package main

// https://leetcode.com/problems/contains-duplicate/submissions/944386393/

func containsDuplicate(nums []int) bool {
	numsMap := make(map[int]int)
	for _, v := range nums {
		if _, ok := numsMap[v]; ok {
			return true
		}
		numsMap[v] = 1
	}
	return false
}
