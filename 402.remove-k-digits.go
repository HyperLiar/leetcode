/*
 * @lc app=leetcode id=402 lang=golang
 *
 * [402] Remove K Digits
 */

// @lc code=start
func removeKdigits(num string, k int) string {
    stack := []byte{}

    for i := range num {
        digit := num[i]
        for k > 0 && len(stack) > 0 && digit < stack[len(stack)-1] {
            stack = stack[:len(stack)-1]
            k--
        }

        stack = append(stack, digit)
    }
    // if ascending order, move the last k byte
    stack = stack[:len(stack)-k]
    ans := strings.TrimLeft(string(stack), "0")
    if ans == "" {
        ans = "0"
    }

    return ans
}
// @lc code=end

