package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, n := range nums {
		c := target - n

		if _, prs := m[c]; prs {
			return []int{m[c], i}
		}

		m[n] = i
	}

	panic("No two sum solution")
}

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}
