/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 */

// @lc code=start
func search(nums []int, target int) int {
    left, right := 0, len(nums)-1

    for left <= right {
        mid := (left+right) / 2
        if target == nums[mid] {
            return mid
        }

        if nums[mid] >= nums[left] {
            if target >= nums[left] && target < nums[mid] {
                right = mid-1
            } else {
                left = mid+1
            }
        } else {
            if target <= nums[right] && target > nums[mid] {
                left = mid+1
            } else {
                right = mid-1
            }
        }
    }

    return -1
}

func search1(nums []int, target int) int {
    left, right := 0, len(nums)-1

    for left <= right {
        mid := (left+right) / 2
        if nums[left] == target {
            return left
        } else if nums[right] == target {
            return right
        } else if nums[mid] == target {
            return mid
        }

        if nums[mid] > nums[left] {
            if target > nums[mid] {
                left = mid+1
            } else {
                if target < nums[left] {
                    left = mid+1
                } else {
                    right = mid-1
                }
            }
        } else {
            if target < nums[mid] {
                right = mid-1
            } else {
                if target > nums[right] {
                    right = mid-1
                } else {
                    left = mid+1
                }
            }
        }
    }

    return -1
}
// @lc code=end

