package main

import (
	"fmt"
	"leetcode/array/Medium/longest_consecutive_sequence"
)

func main() {
	fmt.Println("Ввели [1,8,6,2,5,4,8,3,7],  ожидаем ответ 49")
	fmt.Println(longest_consecutive_sequence.LongestConsecutive([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println("Ввели [1,1],  ожидаем ответ 1")
	fmt.Println(longest_consecutive_sequence.LongestConsecutive([]int{1, 1}))
}
