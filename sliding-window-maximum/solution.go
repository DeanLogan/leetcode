package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindow([]int{1}, 1))
}

func maxSlidingWindow(nums []int, k int) []int {
	output := []int{}
	queue := make([]int, 0)
	left, right := 0, 0

	for right < len(nums) {
		queueLen := len(queue)
		for queueLen != 0 && nums[queue[queueLen-1]] < nums[right] {
			queue = queue[:queueLen-1]
		}
		queue = append(queue, right)

		if left > queue[0] {
			queue = queue[1:]
		}

		if (right + 1) >= k {
			output = append(output, nums[queue[0]])
			left++
		}
		right++
	}
	return output
}