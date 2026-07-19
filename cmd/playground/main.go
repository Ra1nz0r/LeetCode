package main

import (
	"fmt"
	"leetcode/array/Medium/subarray_sum_equals_k"
)

func main() {
	fmt.Println("Ввели [1,1,1],  ожидаем ответ 2")
	fmt.Println(subarray_sum_equals_k.SubarraySum([]int{1, 1, 1}, 1))
	fmt.Println("Ввели [1,2,3],  ожидаем ответ 3")
	fmt.Println(subarray_sum_equals_k.SubarraySum([]int{1, 2, 3}, 3))
}
