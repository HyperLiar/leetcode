/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] ZigZag Conversion
 */

// @lc code=start
func convert(s string, numRows int) string {
    l := len(s)
    if numRows == 1 {
        return s
    }
    count := (l - numRows) / (2 * numRows - 2) + 1
    width := 1 + count * (numRows-1)

    m := make([][]byte, numRows)
    for i := 0; i < numRows; i++ {
        m[i] = make([]byte, width)
    }

    i, j, k := 0, 0, 0
    for {
        for j < numRows && i < len(s) {
            m[j][k] = s[i]
            i, j = i+1, j+1
        }
        j, k = j-2, k+1
        for j >= 1 && k < width-1 && i < len(s) {
            m[j][k] = s[i]
            i, j, k = i+1, j-1, k+1
        }
        
        if i == len(s) {
            break
        }
    }
    res := []byte{}
    for i := 0; i < numRows; i++ {
        for j := 0; j < width; j++ {
            if int(m[i][j]) != 0 {
                res = append(res, m[i][j])
            }
        }
    }

    return string(res)
}
// @lc code=end

