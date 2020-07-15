package str

//(中等)
//无重复字符的最长子串

//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

//示例 1:
//
//输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//
//输入: "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//示例 3:
//
//输入: "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

func lengthOfLongestSubstring1(str string) int {
	strs := []rune(str)
	length := 0
	strLen := len(strs)
	for i := 0; i < strLen; i++ {
		m := make(map[string]*int, 0)
		if length > strLen-i {
			for j := i; j < strLen; j++ {
				if m[string(strs[j])] == nil {
					m[string(strs[j])] = &j
				} else {
					break
				}
			}
		}
		if len(m) > length {
			length = len(m)
		}
	}

	return length
}
