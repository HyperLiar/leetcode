/*
 * @lc app=leetcode id=445 lang=golang
 *
 * [445] Add Two Numbers II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    s1, s2 := []*ListNode{},[]*ListNode{}

    for ; l1 != nil; l1 = l1.Next {
        s1 = append(s1, l1)
    }

    for ; l2 != nil; l2 = l2.Next {
        s2 = append(s2, l2)
    }

    var l3, temp *ListNode
    var bias int
    for len(s1) != 0 || len(s2) != 0 || bias == 1 {
        l3 = &ListNode{Next:temp}

        sum := bias
        if len(s1) != 0 {
            sum += s1[len(s1)-1].Val
            s1 = s1[0:len(s1)-1]
        }
        if len(s2) != 0 {
            sum += s2[len(s2)-1].Val
            s2 = s2[0:len(s2)-1]
        }
        if sum >= 10 {
            sum -= 10
            bias = 1
        } else {
            bias = 0
        }
        l3.Val = sum
        temp = l3
    }

    return l3
}

func reverse(l *ListNode) *ListNode {
    var prev, temp *ListNode
    curr := l
    for curr != nil {
        temp = curr.Next
        curr.Next = prev
        prev = curr
        curr = temp
    }

    return prev
}
// @lc code=end

