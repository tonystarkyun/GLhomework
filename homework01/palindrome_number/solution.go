package palindromenumber

// IsPalindrome checks whether x is a palindrome number.
func IsPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reversedHalf := 0
	for x > reversedHalf {
		reversedHalf = reversedHalf*10 + x%10
		x /= 10
	}

	return x == reversedHalf || x == reversedHalf/10
}
