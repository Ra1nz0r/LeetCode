package main

import (
	"fmt"
	"leetcode/array/Medium/longest_consecutive_sequence"
)

func main() {
	fmt.Println("Ввели [100,4,200,1,3,2],  ожидаем ответ 4")
	fmt.Println(longest_consecutive_sequence.LongestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println("Ввели [0,3,7,2,5,8,4,6,0,1],  ожидаем ответ 9")
	fmt.Println(longest_consecutive_sequence.LongestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	fmt.Println("Ввели [1,0,1,2],  ожидаем ответ 3")
	fmt.Println(longest_consecutive_sequence.LongestConsecutive([]int{1, 0, 1, 2}))
	fmt.Println("Ввели [1,100,6,101,102,99], ожидаем ответ 3")
	fmt.Println(longest_consecutive_sequence.LongestConsecutive([]int{1, 100, 6, 101, 102, 99}))
}
