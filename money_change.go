package main

import "fmt"

func main() {
	fmt.Println(makeList(1))
	fmt.Println(makeList(2))

}

func makeList(n int) [][]int {
	res := make([][]int, 0)
	cur := []int{}
	backTrack(&res, &cur, 0, 0, n)
	return res
}

func backTrack(res *[][]int, cur *[]int, start, end, max int) {
	if len(*cur) == max*2 {
		temp := append([]int{}, *cur...)
		*res = append(*res, temp)
		return
	}

	if start < max {
		*cur = append(*cur, 5)
		backTrack(res, cur, start+1, end, max)
		*cur = (*cur)[0:len(*cur)-1]
	}

	if end < start {
		*cur = append(*cur, 10)
		backTrack(res, cur, start, end+1, max)
		*cur = (*cur)[0:len(*cur)-1]
	}
}