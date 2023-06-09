package main

import "fmt"

// https://leetcode.com/problems/roman-to-integer/

func romanToInt(s string) int {
	res := 0
	romanInt := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i := 0; i < len(s)-1; i++ {
		if romanInt[rune(s[i])] < romanInt[rune(s[i+1])] {
			res -= romanInt[rune(s[i])]
		} else {
			res += romanInt[rune(s[i])]
		}
	}
	return res + romanInt[rune(s[len(s)-1])]

}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
