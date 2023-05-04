package main

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/

//Input: nums = [0,0,1,1,1,2,2,3,3,4]
//Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
//Explanation: Your function should return k = 5, with the first five elements of nums being 0, 1, 2, 3,
//and 4 respectively.
//It does not matter what you leave beyond the returned k (hence they are underscores).

func removeDuplicates(nums []int) int {
	place := 1
	cnt := 0
	for cur := 1; cur < len(nums); cur++ {
		if nums[cur-1] != nums[cur] {
			if cnt != 0 || cur == len(nums)-1 {
				nums[place] = nums[cur]
				cnt = 1
			}
			place++
		} else {
			cnt++
		}
	}
	return place
}

func removeDuplicates2(nums []int) int {
	i := 1
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[j-1] {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
