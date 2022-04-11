package main

func combine(n int, k int) [][]int {
	if n == k {
		r := make([][]int, 0)
		res := make([]int, 0)
		for i := 1; i <= n; i++ {
			res = append(res, i)
		}
		r = append(r, res)
		return r
	}
	if k == 1 {
		r := make([][]int, 0)
		for i := 1; i <= n; i++ {
			r = append(r, []int{i})
		}
		return r
	}

	c := combine(n-1, k-1)
	length := len(c)
	for i := 0; i < length; i++ {
		c[i] = append(c[i], n)
	}
	c = append(c, combine(n-1, k)...)
	return c
}
