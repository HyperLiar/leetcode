/*
 * @lc app=leetcode id=1239 lang=golang
 *
 * [1239] Maximum Length of a Concatenated String with Unique Characters
 */

// @lc code=start
func maxLength(arr []string) int {
    return dfs(arr, 0, 0)
}

func dfs(arr []string, start, bitMask int) int {
    if start == len(arr) {
        return 0
    }

    ans :=  0

    for i := start; i < len(arr); i++ {
        bit := getBitMask(arr[i])
        if bit == 0 || bitMask & bit != 0 {
            continue
        }
        k := dfs(arr, i+1, bitMask | bit) + len(arr[i])
        if k > ans {
            ans = k
        }
    }

    return ans
}

func getBitMask(s string) int {
    bitMask := 0
    for i := 0; i < len(s); i++{
        bit := 1 << (s[i] - 'a')
        if (bit & bitMask) != 0 {
            return 0
        }
        bitMask |= bit
    }

    return bitMask
}

func maxLength1(arr []string) int {
    m := make(map[byte]struct{})
    maxLen := 0
    dfs(arr, 0, &maxLen, &m)
    return maxLen
}

func dfs(arr []string, idx int, maxLen *int, m *map[byte]struct{}) {
    if len(*m) > *maxLen {
        *maxLen = len(*m)
    } 
    if idx == len(arr) {
        return
    }

    // do not connect
    dfs(arr, idx+1, maxLen, m)
    // do connect
    c := true
    m1 := make(map[byte]struct{})
    for i := 0; i < len(arr[idx]); i++ {
        if _, ok := (*m)[arr[idx][i]]; ok {
            c = false
            break
        } else if _, ok := m1[arr[idx][i]]; ok {
            c = false
            break
        }
        m1[arr[idx][i]] = struct{}{}
    }
    if !c {
        return
    }

    for i := 0; i < len(arr[idx]);i++ {
        (*m)[arr[idx][i]] = struct{}{}
    }
    dfs(arr, idx+1, maxLen, m)
    for i := 0; i < len(arr[idx]);i++ {
        delete(*m, arr[idx][i])
    }
}

// @lc code=end

