package main

import "fmt"

func main() {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
}
func intersect(nums1 []int, nums2 []int) []int {
	var res []int
	base := nums1
	o := nums2
	if len(base) > len(nums2) {
		base = nums2
		o = nums1
	}
	m := make(map[int]int)
	for _, i := range base {
		if c, ok := m[i]; ok {
			m[i] = c + 1
		} else {
			m[i] = 1
		}
	}

	for _, i := range o {
		if c, ok := m[i]; ok {
			res = append(res, i)
			if c == 1 {
				delete(m, i)
			} else {
				m[i] = c - 1
			}
		}
	}
	return res
}
