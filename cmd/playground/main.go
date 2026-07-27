package main

import (
	"fmt"
	"leetcode/array/Medium/sort_colors"
)

func main() {
	fmt.Println("Ввели [2,0,2,1,1,0],  ожидаем ответ [0,0,1,1,2,2]")
	sort_colors.SortColors([]int{2, 0, 2, 1, 1, 0})
	fmt.Println("Ввели [2,0,1],  ожидаем ответ [0,1,2]")
	sort_colors.SortColors([]int{2, 0, 1})
}
