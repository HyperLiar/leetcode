/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */

// @lc code=start
func subsets(nums []int) [][]int {
	result := make([][]int, 1)
	result[0] = []int{}

	for i := 0; i < len(nums); i++ {
		tempResult := make([][]int, 0)
		for j := 0; j < len(result); j++ {
			temp := make([]int, len(result[j]))
			copy(temp, result[j])
			tempResult = append(tempResult, append(temp, nums[i]))
		}
		result = append(result, tempResult...)
	}
	return result
}

public void backtrack(int first, ArrayList<Integer> curr, int[] nums) {
    // if the combination is done
    if (curr.size() == k) {
      output.add(new ArrayList(curr));
      return;
    }
    for (int i = first; i < n; ++i) {
      // add i into the current combination
      curr.add(nums[i]);
      // use next integers to complete the combination
      backtrack(i + 1, curr, nums);
      // backtrack
      curr.remove(curr.size() - 1);
    }
  }

  public List<List<Integer>> subsets(int[] nums) {
    n = nums.length;
    for (k = 0; k < n + 1; ++k) {
      backtrack(0, new ArrayList<Integer>(), nums);
    }
    return output;
  }
// @lc code=end

