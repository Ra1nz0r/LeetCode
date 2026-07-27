package main

import (
	"fmt"
	"leetcode/array/Medium/top_k_frequent_elements"
)

func main() {
	fmt.Println("Ввели [1,1,1,2,2,3] и 2,  ожидаем ответ [1,2]")
	top_k_frequent_elements.TopKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	fmt.Println("Ввели [1] и 1,  ожидаем ответ [1]")
	top_k_frequent_elements.TopKFrequent([]int{1}, 1)
	fmt.Println("Ввели [1,2,1,2,1,2,3,1,3,2] и 2,  ожидаем ответ [1,2]")
	top_k_frequent_elements.TopKFrequent([]int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}, 2)
}
