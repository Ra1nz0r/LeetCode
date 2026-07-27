package main

import (
	"fmt"
	"leetcode/array/Medium/top_k_frequent_elements"
)

func main() {
	fmt.Println("Ввели [10,10,10,20,20,30] и 2,  ожидаем ответ [10,20]")
	fmt.Println(top_k_frequent_elements.TopKFrequent([]int{10, 10, 10, 20, 20, 30}, 2))

	fmt.Println("Ввели [1] и 1,  ожидаем ответ [1]")
	fmt.Println(top_k_frequent_elements.TopKFrequent([]int{1}, 1))

	fmt.Println("Ввели [1,2,1,2,1,2,3,1,3,2] и 2,  ожидаем ответ [1,2]")
	fmt.Println(top_k_frequent_elements.TopKFrequent([]int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}, 2))
}
