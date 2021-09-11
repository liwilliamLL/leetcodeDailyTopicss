package leetcodeDailyTopicss

import "sort"

/*
一个班级里有 n 个学生，编号为 0 到 n - 1 。每个学生会依次回答问题，编号为 0 的学生先回答，然后是编号为 1 的学生，以此类推，直到编号为 n - 1 的学生，然后老师会重复这个过程，重新从编号为 0 的学生开始回答问题。

给你一个长度为 n 且下标从 0 开始的整数数组 chalk 和一个整数 k 。一开始粉笔盒里总共有 k 支粉笔。当编号为 i 的学生回答问题时，他会消耗 chalk[i] 支粉笔。如果剩余粉笔数量 严格小于 chalk[i] ，那么学生 i 需要 补充 粉笔。

请你返回需要 补充 粉笔的学生 编号 。

示例 1：

输入：chalk = [5,1,5], k = 22
输出：0
解释：学生消耗粉笔情况如下：
- 编号为 0 的学生使用 5 支粉笔，然后 k = 17 。
- 编号为 1 的学生使用 1 支粉笔，然后 k = 16 。
- 编号为 2 的学生使用 5 支粉笔，然后 k = 11 。
- 编号为 0 的学生使用 5 支粉笔，然后 k = 6 。
- 编号为 1 的学生使用 1 支粉笔，然后 k = 5 。
- 编号为 2 的学生使用 5 支粉笔，然后 k = 0 。
编号为 0 的学生没有足够的粉笔，所以他需要补充粉笔。
示例 2：

输入：chalk = [3,4,1,2], k = 25
输出：1
解释：学生消耗粉笔情况如下：
- 编号为 0 的学生使用 3 支粉笔，然后 k = 22 。
- 编号为 1 的学生使用 4 支粉笔，然后 k = 18 。
- 编号为 2 的学生使用 1 支粉笔，然后 k = 17 。
- 编号为 3 的学生使用 2 支粉笔，然后 k = 15 。
- 编号为 0 的学生使用 3 支粉笔，然后 k = 12 。
- 编号为 1 的学生使用 4 支粉笔，然后 k = 8 。
- 编号为 2 的学生使用 1 支粉笔，然后 k = 7 。
- 编号为 3 的学生使用 2 支粉笔，然后 k = 5 。
- 编号为 0 的学生使用 3 支粉笔，然后 k = 2 。
编号为 1 的学生没有足够的粉笔，所以他需要补充粉笔。
 

提示：

chalk.length == n
1 <= n <= 105
1 <= chalk[i] <= 105
1 <= k <= 109

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-the-student-that-will-replace-the-chalk
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */


func chalkReplacer(chalk []int, k int) int {
	//ns :=make([]int,len(chalk))
	total := 0
	for v,t:=range chalk {
		total += t
		chalk[v]=total
	}
	k = k%total
	//这里时间充足作为练手，手撸一个二分查找的代码
	var left = 0
	var right = len(chalk) - 1
	for right - left > 1{
		var mid = (left + right) / 2
		if chalk[mid] == k {
			return mid+1
		} else if chalk[mid] > k {
			right = mid
		} else {
			left = mid
		}
	}
	if chalk[left] >= k {
		return left
	}
	return right
}



/*
执行结果：
通过
显示详情


执行用时：108 ms, 在所有 Go 提交中击败了98.86%的用户
内存消耗：8.1 MB, 在所有 Go 提交中击败了66.86%的用户
通过测试用例：
72 / 72
 */

func chalkReplacer2(chalk []int, k int) int {
	//为了节省内存在原数据的切片进行自增
	//ns :=make([]int,len(chalk))
	total := 0
	for v,t:=range chalk {
		total += t
		chalk[v]=total
	}
	k = k%total
	//这里用官方包提供的sort.SearchInts()
	return sort.SearchInts(chalk,k+1)
}




