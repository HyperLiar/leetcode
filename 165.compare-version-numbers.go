/*
 * @lc app=leetcode id=165 lang=golang
 *
 * [165] Compare Version Numbers
 */

// @lc code=start
func compareVersion(version1, version2 string) int{
    l1, l2 := len(version1), len(version2)
    idx1, idx2 := 0, 0
    for idx1 < l1 || idx2 < l2 {
        if idx1 < l1 && version1[idx1] == '.' {
            idx1++
        }
        
        if idx2 < l2 && version2[idx2] == '.'{
            idx2++
        }

        curr1, curr2 := 0, 0
        for idx1 < l1 && version1[idx1] != '.' {
            curr1 = 10 * curr1 + int(version1[idx1]-'0')
            idx1++
        }
        for idx2 < l2 && version2[idx2] != '.' {
            curr2 = 10 * curr2 + int(version2[idx2]-'0')
            idx2++
        }
        if curr1 > curr2 {
            return 1
        } else if curr1 < curr2 {
            return -1
        }
    }

    return 0
}

func compareVersion1(version1 string, version2 string) int {
    a1, a2 := strings.Split(version1, "."), strings.Split(version2, ".")
    swap := 1
    if len(a1) > len(a2) {
        a1, a2 = a2, a1
        swap = -1
    }

    for i := 0; i < len(a1); i++ {
        i1, i2 := moveZero(a1[i]), moveZero(a2[i])

        if i1 > i2 {
            return swap
        } else if i1 <i2 {
            return -swap
        }
    }

    for i := len(a1); i < len(a2); i++ {
        if moveZero(a2[i]) != 0 {
            return -swap
        }
    }

    return 0
}
func moveZero(s string) int {
    i := 0
    for ; i < len(s); i++ {
        if s[i] != '0' {
            break
        }
    }
    r, _ := strconv.Atoi(s[i:])
    return r
}
// @lc code=end

