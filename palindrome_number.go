package main

// https://leetcode.com/problems/palindrome-number/

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	tmp := x
	reversedNum := 0
	for tmp > 0 {
		reversedNum = reversedNum*10 + tmp%10
		tmp /= 10
	}
	return reversedNum == x
}
