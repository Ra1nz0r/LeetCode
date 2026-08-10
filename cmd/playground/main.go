package main

import (
	"fmt"
	"leetcode/array/Medium/maximum_subarray"
)

func main() {
	fmt.Println("Ввели [-2,1,-3,4,-1,2,1,-5,4] ожидаем ответ 6")
	fmt.Println(maximum_subarray.MaxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))

	fmt.Println("Ввели [1] ожидаем ответ 1")
	fmt.Println(maximum_subarray.MaxSubArray([]int{1}))

	fmt.Println("Ввели [5,4,-1,7,8] ожидаем ответ 23")
	fmt.Println(maximum_subarray.MaxSubArray([]int{5, 4, -1, 7, 8}))
	fmt.Println("Ввели [5,4,-1,7,8] ожидаем ответ 23")
	fmt.Println("Ввели [5,4,-1,7,8] ожидаем ответ 23")
}
