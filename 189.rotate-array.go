/*
 * @lc app=leetcode id=189 lang=golang
 *
 * [189] Rotate Array
 */

// @lc code=start
func rotate(nums []int, k int)  {
    l := len(nums)
    k = k % l
    r(nums)
    r(nums[0:k])
    r(nums[k:])
}

func r(nums []int) {
    for i := 0; i < len(nums)/2; i++ {
        nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
    }
}
// @lc code=end

