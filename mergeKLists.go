package lesson3

type ListNode struct {
	Val  int
	Next *ListNode
}

//MergeKLists: 合并K个升序链表
/*parameter
lists: 链表数组
return: 新的有序链表
*/
func MergeKLists(lists []*ListNode) *ListNode {

	return merge(lists, 0, len(lists)-1)
}

//merge: 递归合并K个升序链表
/*parameter
lists: 链表数组
l, r: 左右边界
return: 新的有序链表
*/
func merge(lists []*ListNode, l, r int) *ListNode {

	switch {
	case l == r:
		return lists[l]
	case l > r:
		return nil
	}

	mid := (l + r) >> 1
	return mergeTwoLists(merge(lists, l, mid), merge(lists, mid+1, r))
}

//mergeTwoLists: 合并两个有序链表
/*parameter
list1: 有序链表1
list2: 有序链表2
return: 新的有序链表
*/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	//如果有一个链表为空则返回另一链表，两个都为空则返回Nil
	if list1 == nil {

		return list2
	}

	if list2 == nil {

		return list1
	}

	newList := &ListNode{
		Val: -100,
	}
	node := newList

	for list1 != nil && list2 != nil {

		if list1.Val <= list2.Val {

			node.Next = list1
			list1 = list1.Next
		} else {

			node.Next = list2
			list2 = list2.Next
		}

		node = node.Next
	}

	if list1 == nil {

		node.Next = list2
	} else {

		node.Next = list1
	}

	return newList.Next
}
