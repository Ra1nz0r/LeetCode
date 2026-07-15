package main

import (
	"fmt"
	"leetcode/array/Medium/container_with_most_water"
)

func main() {
	fmt.Println("Ввели [1,8,6,2,5,4,8,3,7],  ожидаем ответ 49")
	fmt.Println(container_with_most_water.MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println("Ввели [1,1],  ожидаем ответ 1")
	fmt.Println(container_with_most_water.MaxArea([]int{1, 1}))
}
