package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	nums := nums(60)

	if sum(nums) != 1830 {
		t.Fail()
	}
}

func TestSumParallel2(t *testing.T) {
	nums := nums(60)

	if sumParallel(nums, 2) != 1830 {
		t.Fail()
	}
}

func TestSumParallel6(t *testing.T) {
	nums := nums(60)

	if sumParallel(nums, 6) != 1830 {
		t.Fail()
	}
}

func BenchmarkSum(b *testing.B) {
	nums := nums(600000)

	for i := 0; i < b.N; i++ {
		sum(nums)
	}
}

func BenchmarkSumParallel2(b *testing.B) {
	nums := nums(600000)

	for i := 0; i < b.N; i++ {
		sumParallel(nums, 2)
	}
}

func BenchmarkSumParallel4(b *testing.B) {
	nums := nums(600000)

	for i := 0; i < b.N; i++ {
		sumParallel(nums, 4)
	}
}

func BenchmarkSumParallel6(b *testing.B) {
	nums := nums(600000)

	for i := 0; i < b.N; i++ {
		sumParallel(nums, 6)
	}
}

func BenchmarkSumParallel8(b *testing.B) {
	nums := nums(600000)

	for i := 0; i < b.N; i++ {
		sumParallel(nums, 8)
	}
}

func BenchmarkSumParallel10(b *testing.B) {
	nums := nums(600000)

	for i := 0; i < b.N; i++ {
		sumParallel(nums, 10)
	}
}
