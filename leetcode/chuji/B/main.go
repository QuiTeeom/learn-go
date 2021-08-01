package main

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
}

func moveZeroes(nums []int) {
	var i, j, k int
	for j < len(nums) {
		if nums[j] == 0 {
			copy(nums[k:], nums[i:j])
			k = k + j - i
			i = j + 1
		}
		j++
	}
	copy(nums[k:], nums[i:j])
	k = k + j - i
	for k < len(nums) {
		nums[k] = 0
	}
}
