package main

import (
	"log"
)

func main() {
	nums := nums(10)

	total := sumParallel(nums, 2)

	log.Println(total)
}

func nums(amount int) []int {
	nums := []int{}

	for i := 1; i <= amount; i++ {
		nums = append(nums, i)
	}

	return nums
}

func sum(nums []int) int {
	total := 0

	for _, n := range nums {
		total += n
	}

	return total
}

func sumParallel(nums []int, concurrency int) int {
	ch := make(chan int)
	chunks := chunk(nums, len(nums)/concurrency)

	for i := 0; i < concurrency; i++ {
		go func(i int) {
			ch <- sum(chunks[i])
		}(i)
	}

	total := 0

	for i := 0; i < concurrency; i++ {
		total += <-ch
	}

	return total
}

func chunk(nums []int, size int) (chunks [][]int) {
	for size < len(nums) {
		nums, chunks = nums[size:], append(chunks, nums[0:size:size])
	}

	return append(chunks, nums)
}
