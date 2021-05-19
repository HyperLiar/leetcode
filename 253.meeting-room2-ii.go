func minMeetingRooms(intervals [][]int) int {
    m := make([][]int, 0)
    for i := 0; i < len(intervals); i++ {
        start, end := []int{}, []int{}
        start = append(start, intervals[i][0], 1)
        end = append(end, intervals[i][1], -1)
        m = append(m, start, end)
    }

    sort.Slice(m, func(i, j int) bool {
        if m[i][0] < m[j][0] {
            return true
        } else if m[i][0] > m[j][0] {
            return false
        } else {
            return m[i][1] < m[j][1]
        }
    })

    res, cur := 0, 0
    for i := 0; i < len(m); i++ {
        cur += m[i][1]
        if cur > res {
            res = cur
        }
    }
    return res
}