package leetcodeDailyTopicss


/*
给你一个字符串 s ，找出其中最长的回文子序列，并返回该序列的长度。

子序列定义为：不改变剩余字符顺序的情况下，删除某些字符或者不删除任何字符形成的一个序列。

 

示例 1：

输入：s = "bbbab"
输出：4
解释：一个可能的最长回文子序列为 "bbbb" 。
示例 2：

输入：s = "cbbd"
输出：2
解释：一个可能的最长回文子序列为 "bb" 。
 

提示：

1 <= s.length <= 1000
s 仅由小写英文字母组成

 */

func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([]int, n )
	for i:=0; i<n; i++ {
		m := 0
		for j:=i-1;j>=0; j-- {
			t := dp[j]
			if s[i] == s[j] {
				dp[j] = m+2
			}
			if t > m {
				m = t
			}
		}
		dp[i] = 1
	}
	max := 0
	for _,v := range dp {
		if v > max {
			max = v
		}
	}
	return max
}
