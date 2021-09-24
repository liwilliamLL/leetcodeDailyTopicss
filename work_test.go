package leetcodeDailyTopicss

import (
	"testing"
)

func Test_serivce(t *testing.T) {

	tests := []struct {
		name    string
		args    interface{}
		wantErr bool
	}{
		{name: "TestLeetCode2021-8-11", args: []int{2,4,6,8,10}},
		{name: "TestLeetCode2021-8-11", args: []int{7, 7, 7}},
		{name: "TestLeetCode2021-8-12", args: "bbbb"},
		{name: "TestLeetCode2021-8-13", args: 1000},
		{name: "TestLeetCode2021-8-16", args: 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.name {
			case "TestLeetCode2021-8-11":
				res := numberOfArithmeticSlices(tt.args.([]int))
				t.Log(res)
			case "TestLeetCode2021-8-12":
				res := longestPalindromeSubseq(tt.args.(string))
				t.Log(res)
			case "TestLeetCode2021-8-13":
				res := countDigitOne(tt.args.(int))
				t.Log(res)
			case "TestLeetCode2021-8-16":
				testList()
				res := countArrangement(tt.args.(int))
				t.Log(res)
			}
		})
	}
}
