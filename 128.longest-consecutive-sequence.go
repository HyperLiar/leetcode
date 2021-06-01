/*
 * @lc app=leetcode id=128 lang=golang
 *
 * [128] Longest Consecutive Sequence
 */

// @lc code=start
func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    uf := UF{m:map[int]int{}, cnt:map[int]int{}}

    for _, num := range nums {
        uf.m[num] = num
        uf.cnt[num] = 1
    }
    res := 1
    for _, num := range nums {
        if _, ok := uf.m[num+1]; ok {
            y :=  uf.merge(num, num+1)
            if y > res {
                res = y
            }
        }
    }

    return res
}

type UF struct {
    m map[int]int
    cnt map[int]int
}

func (this *UF) find(x int) int {
    if x == this.m[x] {
        return x
    } else {
        return this.find(this.m[x])
    }
}

func (this *UF) merge(x, y int) int {
    x, y = this.find(x), this.find(y)
    if x == y {
        return this.cnt[x]
    }
    this.m[y] = x
    this.cnt[x] += this.cnt[y]
    return this.cnt[x]
}
// @lc code=end

