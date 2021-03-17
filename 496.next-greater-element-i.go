/*
 * @lc app=leetcode id=496 lang=golang
 *
 * [496] Next Greater Element I
 */

// @lc code=start
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	s := make([]int, 0)
	r := make([]int, len(nums1))

	for i := 0; i < len(nums2); i++ {
		for len(s) > 0 && s[len(s)-1] < nums2[i] {
			// if num at top of stack is smaller
			m[s[len(s)-1]] = nums2[i]
			// pop the top of stack
			s = s[0 : len(s)-1]
		}
		s = append(s, nums2[i])
	}

	for i := 0; i < len(s); i++ {
		m[s[i]] = -1
	}

	for i := 0; i < len(nums1); i++ {
		r[i] = m[nums1[i]]
	}

	return r
}

// @lc code=end

