/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	var r *ListNode

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val <= l2.Val {
		r = l1
		r.Next = mergeTwoLists(l1.Next, l2)
	} else {
		r = l2
		r.Next = mergeTwoLists(l1, l2.Next)
	}
	return r
}

func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }
    
    end := len(lists) - 1
	for end > 0 {
		begin := 0
		for begin < end {
			lists[begin] = mergeTwoLists(lists[begin], lists[end])
			begin++
			end--
		}
	}
    return lists[0]
}