package leetcodeDailyTopicss

/*
给你一个头结点为 head 的单链表和一个整数 k ，请你设计一个算法将链表分隔为 k 个连续的部分。

每部分的长度应该尽可能的相等：任意两部分的长度差距不能超过 1 。这可能会导致有些部分为 null 。

这 k 个部分应该按照在链表中出现的顺序排列，并且排在前面的部分的长度应该大于或等于排在后面的长度。

返回一个由上述 k 部分组成的数组。


示例 1：

输入：head = [1,2,3], k = 5
输出：[[1],[2],[3],[],[]]
解释：
第一个元素 output[0] 为 output[0].val = 1 ，output[0].next = null 。
最后一个元素 output[4] 为 null ，但它作为 ListNode 的字符串表示是 [] 。
示例 2：


输入：head = [1,2,3,4,5,6,7,8,9,10], k = 3
输出：[[1,2,3,4],[5,6,7],[8,9,10]]
解释：
输入被分成了几个连续的部分，并且每部分的长度相差不超过 1 。前面部分的长度大于等于后面部分的长度。

提示：

链表中节点的数目在范围 [0, 1000]
0 <= Node.val <= 1000
1 <= k <= 50

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/split-linked-list-in-parts
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
执行结果：通过 显示详情 添加备注

执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：3 MB, 在所有 Go 提交中击败了5.06%的用户
通过测试用例：
43 / 43
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

//用多一个数组的方法
func splitListToParts(head *ListNode, k int) (re []*ListNode) {
	re = make([]*ListNode, k)
	nodeSlice := make([]*ListNode, 0, k)
	for head != nil {
		nodeSlice = append(nodeSlice, head)
		head = head.Next
	}
	total := len(nodeSlice)
	if total == 0 {
		return re
	}
	offset := -1
	long, s := total%k, total/k
	if s == 0 {
		k = long
	}
	for i := 0; i < k; i++ {
		re[i] = nodeSlice[offset+1]
		offset = offset + s
		if i < long {
			offset = offset + 1
		}
		nodeSlice[offset].Next = nil
	}
	return
}

/*
执行结果：通过 显示详情 添加备注

执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.8 MB, 在所有 Go 提交中击败了100.00%的用户
通过测试用例：
43 / 43
 */
func splitListToPartss(head *ListNode, k int) (re []*ListNode) {
	re = make([]*ListNode, k)
	if head == nil {
		return re
	}
	headCopy := head
	total := 0
	for head != nil {
		total++
		head = head.Next
	}
	long, s := total%k, total/k
	for i := 0; i < k && headCopy != nil; i++ {
		re[i] = headCopy
		partSize := s
		if i < long {
			partSize++
		}
		for j := 1; j < partSize; j++ {
			headCopy = headCopy.Next
		}
		headCopy, headCopy.Next = headCopy.Next, nil
	}

	return
}
