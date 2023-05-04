package main

import "fmt"

// https://leetcode.com/problems/longest-common-prefix/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	res := strs[0]
	for _, v := range strs {
		if len(res) > len(v) {
			res = string([]rune(res)[:len(v)])
		}
		for j, ch := range []rune(v) {
			if len(res) > j {
				if rune(res[j]) != ch {
					res = string([]rune(res)[:j])
					break
				}
			} else {
				break
			}
		}
	}
	return res
}

func main() {
	s := []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(s))
	s = []string{"ab", "a"}
	fmt.Println(longestCommonPrefix(s))
}
