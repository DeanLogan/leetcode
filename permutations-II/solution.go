package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(permuteUnique([]int{1,1,2}))
	fmt.Println(permuteUnique([]int{1,2,3}))
}

func permuteUnique(nums []int) [][]int {
    var result [][]int
    sort.Ints(nums)
    used := make([]bool, len(nums))
    backtrack(nums, []int{}, &result, used)
    return result
}

func backtrack(nums, current []int, result *[][]int, used []bool) {
    if len(current) == len(nums) {
        perm := make([]int, len(current))
        copy(perm, current)
        *result = append(*result, perm)
        return
    }

    for i := 0; i < len(nums); i++ {
        if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i-1]) {
            continue
        }
        used[i] = true
        current = append(current, nums[i])
        backtrack(nums, current, result, used)
        current = current[:len(current)-1]
        used[i] = false
    }
}
