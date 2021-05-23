/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start
func threeSum(nums []int) [][]int {
    if len(nums) == 0 {
        return [][]int{}
    }
	sort.Ints(nums)

    res := [][]int{}
    for i := 0; i < len(nums); i++ {
        if i != 0 && nums[i] == nums[i-1] {
            continue
        }

        twoSum := 0 - nums[i]
        start, end := i+1, len(nums)-1
        for start < end {
            oneSum := twoSum-nums[start]
            if nums[end] == oneSum {
                res = append(res, []int{nums[i], nums[start], nums[end]})
                start++
                end--
                for start < end && nums[start] == nums[start-1] {
                    start++
                }
            } else if nums[end] > oneSum {
                end--
            } else {
                start++
            }
        }
    }
    return res
}

// @lc code=end

