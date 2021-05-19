/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int)  {
    a, b := m-1, n-1
    for i := m+n-1; i >= 0; i-- {
        if a < 0 {
            nums1[i] = nums2[b]
            b--
        } else if b < 0 {
            nums1[i] = nums1[a]
            a--
        } else if nums1[a] > nums2[b] {
            nums1[i] = nums1[a]
            a--
        } else {
            nums1[i] = nums2[b]
            b--
        }
    }
}
// @lc code=end

