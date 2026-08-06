package main

import (
	"fmt"
	"leetcode/array/Medium/maximum_subarray"
)

func main() {
	fmt.Println("Ввели [-2,1,-3,4,-1,2,1,-5,4] ожидаем ответ 6")
	fmt.Println(maximum_subarray.MaxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))

}
