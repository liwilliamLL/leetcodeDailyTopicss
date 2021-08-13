package leetcodeDailyTopicss


func countDigitOne(n int) (ans int) {
	for i:=1;i<=n;i=i*10{
		ans+=n/(i*10)*i
		ans+= min(max(n%(i*10)-i+1,0),i)
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

