package main

import (
	"fmt"
	three_sum "leetcode/array/Medium/3_sum"
)

func main() {
	fmt.Println("Ввели [-1,0,1,2,-1,-4],  ожидаем ответ [[-1,-1,2],[-1,0,1]]")
	fmt.Println(three_sum.ThreeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println("Ввели [0,1,1],  ожидаем ответ []")
	fmt.Println(three_sum.ThreeSum([]int{0, 1, 1}))
	fmt.Println("Ввели [0,0,0],  ожидаем ответ [[0,0,0]]")
	fmt.Println(three_sum.ThreeSum([]int{0, 0, 0}))
}
