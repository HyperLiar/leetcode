/*
 * @lc app=leetcode id=25 lang=golang
 *
 * [25] Reverse Nodes in k-Group
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseKGroup(head *ListNode, k int) *ListNode {
    newHead := &ListNode{}
    newHead.Next = head
    cur, pre := newHead, newHead
    var temp *ListNode
    n := -1

    for cur != nil {
        n++
        cur = cur.Next
    }

    for n >= k {
        cur = pre.Next
        for i := 0; i < k-1; i++ {
            temp = cur.Next
            cur.Next = temp.Next
            temp.Next = pre.Next
            pre.Next = temp
        }
        pre = cur
        n -= k
    }
    return newHead.Next
}

func reverseKGroup1(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	var temp, prev, lastTail, tail, newHead *ListNode
	var i int
	for head != nil {
		i = 0
		tail = head
		for i < k && head != nil {
			temp = head.Next
			head.Next = prev
			prev = head
			head = temp
			i++
		}
		if newHead == nil {
			newHead = prev
		}

		if lastTail != nil {
			lastTail.Next = prev
		}
		prev = nil

		if i == k {
			lastTail = tail
		}
	}

	//print(newHead)
	//fmt.Println(prev, tail, lastTail)
	if i != k {
		j := 0
		head = lastTail.Next
		for head != nil && j < i {
			temp = head.Next
			head.Next = prev
			prev = head
			head = temp
			j++
		}
		lastTail.Next = prev
	}

	return newHead
}

func print(l *ListNode) {
	for l != nil {
		fmt.Print(l.Val)
		l = l.Next
	}
	fmt.Println()
}

// @lc code=end

