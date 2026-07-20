package main

import (
	"fmt"
	"leetcode/array/Medium/rotate_array"
)

func main() {
	fmt.Println("Ввели [1,2,3,4,5,6,7] и 3,  ожидаем ответ [5,6,7,1,2,3,4]")
	rotate_array.Rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	fmt.Println("Ввели [-1,-100,3,99] и 2,  ожидаем ответ [3,99,-1,-100]")
	rotate_array.Rotate([]int{-1, -100, 3, 99}, 2)
}
