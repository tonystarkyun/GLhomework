package longestcommonprefix

// LongestCommonPrefix returns the longest common prefix among all strings.
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for i := 1; i < len(strs) && prefix != ""; i++ {
		current := strs[i]

		for len(prefix) > len(current) {
			prefix = prefix[:len(prefix)-1]
		}

		for j := 0; j < len(prefix); j++ {
			if prefix[j] != current[j] {
				prefix = prefix[:j]
				break
			}
		}
	}

	return prefix
}
