package main

import (
	"fmt"
	"leetcode/array/Medium/product_of_array_except_self"
)

func main() {
	fmt.Println("Ввели [1,2,3,4],  ожидаем ответ [24,12,8,6]")
	fmt.Println(product_of_array_except_self.ProductExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println("Ввели [-1,1,0,-3,3],  ожидаем ответ [0,0,9,0,0]")
	fmt.Println(product_of_array_except_self.ProductExceptSelf([]int{-1, 1, 0, -3, 3}))
}
