package digiutils

import (
	"math"
	"regexp"
	"sort"
	"strings"
)

// AddBinary takes two parameters a and b and returns
// their sum as a binary string.
func AddBinary(a, b string) string {
	n, m := len(a), len(b)
	carry := 0
	res := []byte{}
	for i, j := n-1, m-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		x := 0
		if i >= 0 {
			x = int(a[i] - '0')
		}
		y := 0
		if j >= 0 {
			y = int(b[j] - '0')
		}
		sum := x + y + carry
		res = append([]byte{byte(sum%2) + '0'}, res...)
		carry = sum / 2
	}
	if carry > 0 {
		res = append([]byte{'1'}, res...)
	}
	return string(res)
}

// IsPalindrome checks whether s is a palindrome or not
// after converting all uppercase letters into lowercase
// letters and removing all non-alphanumeric characters.
func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(s, "")
	s = strings.ReplaceAll(s, " ", "")

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}

// RomanToInt convert roman numeral to int, it returns
// 0 if s is not a valid roman numeral.
func RomanToInt(s string) (res int) {
	roman := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	i := 0
	for i < len(s) {
		if _, ok := roman[s[i]]; !ok {
			return 0
		}
		if i != len(s)-1 && roman[s[i]] < roman[s[i+1]] {
			res += roman[s[i+1]] - roman[s[i]]
			if i+2 > len(s)-1 {
				return
			}
			i += 2
			continue
		}

		res += roman[s[i]]
		i++
	}
	return
}

// ContainsDuplicate finds an unduplicated number,
// if there is no duplication it will return false.
func ContainsDuplicate(nums []int) bool {
	sort.Ints(nums)

	for i := 0; i <= len(nums)-2; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}

	return false
}


// IsPowerOfTwo returns true if n is a power of two.
// Otherwise, return false.
func IsPowerOfTwo(n int) bool {
	square := 0
	res := 0.0

	for int(res) <= n {
		res = math.Pow(2.0, float64(square))
		if int(res) == n {
			return true
		}
		square++
	}

	return false
}
