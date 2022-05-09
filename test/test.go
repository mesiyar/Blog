package main

import "fmt"

func sortArrayByParity(nums []int) []int {
	l := len(nums)
	left, right := 0, l-1
	for left < right {
		for left < right && nums[left]%2 == 0 {
			left++
		}
		for left < right && nums[right]%2 == 1 {
			right--
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
	return nums
}

func reverse(x int) int {
	ret := 0
	for x != 0 {
		n := x % 10
		x /= 10
		ret = n + ret*10
	}
	return ret
}

func main() {
	//arr := []int{1, 0, 3}
	//fmt.Println(sortArrayByParity(arr))
	//x := 0
	//r := 0
	//for x != 0 {
	//	n := x % 10
	//	x /= 10
	//	r = n + r*10
	//}
	//fmt.Println(r)
	//fmt.Println(intToRoman(1994))
	// a := "abc"
	// b := "c"
	// fmt.Println(strStr(a, b))
	s := "IDID"
	fmt.Println(diStringMatch(s))
}

func diStringMatch(s string) []int {
	l := len(s)
	perm := make([]int, l+1)
	low, high := 0, l
	for i, ch := range s {
		if ch == 'I' {
			perm[i] = low
			low++
		} else {
			perm[i] = high
			high--
		}
	}
	perm[l] = low
	return perm
}

func strStr(haystack string, needle string) int {
	ln := len(needle)
	lh := len(haystack)
	if haystack == needle {
		return 0
	}
	for i := 0; i <= (lh - ln); i++ {
		s := haystack[i : ln+i]
		if s == needle {
			return i
		}
	}
	return -1
}

func intToRoman(num int) string {
	str := ""
	for num > 0 {
		if num >= 1000 {
			str += "M"
			num -= 1000
		} else if num >= 900 {
			str += "CM"
			num -= 900
		} else if num >= 500 {
			str += "D"
			num -= 500
		} else if num >= 400 {
			str += "CD"
			num -= 400
		} else if num >= 100 {
			str += "C"
			num -= 100
		} else if num >= 90 {
			str += "XC"
			num -= 90
		} else if num >= 50 {
			str += "L"
			num -= 50
		} else if num >= 40 {
			str += "XL"
			num -= 40
		} else if num >= 10 {
			str += "X"
			num -= 10
		} else if num >= 9 {
			str += "IX"
			num -= 9
		} else if num >= 5 {
			str += "V"
			num -= 5
		} else if num >= 4 {
			str += "IV"
			num -= 4
		} else {
			str += "I"
			num--
		}
	}
	return str
}
