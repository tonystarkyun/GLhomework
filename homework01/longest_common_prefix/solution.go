package longestcommonprefix

// LongestCommonPrefix returns the longest common prefix among all strings.
func LongestCommonPrefix(strs []string) string {
	// 如果输入切片为空，最长公共前缀只能是空字符串。
	if len(strs) == 0 {
		return ""
	}

	// 先把第一个字符串当作候选公共前缀。
	prefix := strs[0]
	// 从第二个字符串开始，依次和当前候选前缀做比较。
	// 一旦 prefix 变成空串，就可以提前结束循环。
	for i := 1; i < len(strs) && prefix != ""; i++ {
		// 当前要拿来比较的字符串。
		current := strs[i]

		// 如果候选前缀比当前字符串还长，先裁剪到不超过当前字符串长度。
		// 因为公共前缀不可能长于任意一个参与比较的字符串。
		for len(prefix) > len(current) {
			// 每次删除 prefix 最后一个字符，直到长度不大于 current。
			prefix = prefix[:len(prefix)-1]
		}

		// 逐个字符比较 prefix 和 current。
		for j := 0; j < len(prefix); j++ {
			// 遇到第一个不相等的位置时，说明公共前缀到 j-1 为止。
			if prefix[j] != current[j] {
				// 把候选前缀截断到第 j 个字符之前。
				prefix = prefix[:j]
				// 本轮比较结束，继续和下一个字符串比较。
				break
			}
		}
	}

	// 循环结束后，prefix 就是最终的最长公共前缀。
	return prefix
}
