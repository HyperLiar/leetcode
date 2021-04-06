/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
    lBound := searchStart(nums, 0, len(nums)-1, target)
	rBound := searchEnd(nums, 0, len(nums)-1, target)

	return []int{lBound, rBound}
}

func searchStart(nums []int, start, end, target int) int {
	if start > end {
		return -1
	}

	if end - start <=1 {
		if nums[start] == target {
			return start
		} else if nums[end] == target {
			return end
		}

		return -1
	}
	
	mid := (start + end) >> 1
	if nums[mid] > target {
		return searchStart(nums, start, mid-1, target)
	} else if nums[mid] < target {
		return searchStart(nums, mid+1, end, target)
	} else {
		return searchStart(nums, start, mid, target)
	}
}

func searchEnd(nums []int, start, end, target int) int {
	if start > end {
		return -1
	}

	if end - start <=1 {
		if nums[end] == target {
			return end
		} else if nums[start] == target {
			return start
		}

		return -1
	}
	
	mid := (start + end) >> 1
	if nums[mid] > target {
		return searchEnd(nums, start, mid-1, target)
	} else if nums[mid] < target {
		return searchEnd(nums, mid+1, end, target)
	} else {
		return searchEnd(nums, mid, end, target)
	}
}
// @lc code=end

