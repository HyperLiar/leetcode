/*
 * @lc app=leetcode id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 */

// @lc code=start
func myAtoi(s string) int {
    var negative, setted bool
    var res int

    for i := 0; i < len(s); i++ {
        j := s[i] - '0'

        if setted && (j < 0 || j > 9) {
            break
        }
        if s[i] == ' '{
            continue
        } else if s[i] == '-' {
            setted = true
            negative = true
        } else if s[i] == '+' {
            setted = true
        } else {
            setted = true
            j := s[i] - '0'

            if j >= 0 && j <= 9 {
                if negative && res >= 0 {
                    res *= -1
                }
                res = check(res, int(s[i]-'0'), negative)
            } else {
                break
            }
        }
    }

    return res
}

func check(now, bias int, negative bool) int {
    if now == 0 {
        if negative {
            return now * 10 - bias
        } else {
            return now * 10 + bias
        }
    }
    if negative {
        if math.MinInt32 / now < 10 {
            return math.MinInt32
        }
        now *= 10
        if now - math.MinInt32 < bias {
            return math.MinInt32
        }
        return now - bias
    } else {
        if math.MaxInt32 / now < 10 {
            return math.MaxInt32
        }
        now *= 10
        if math.MaxInt32 - now < bias {
            return math.MaxInt32
        } 
        return now + bias
    }
}
// @lc code=end

