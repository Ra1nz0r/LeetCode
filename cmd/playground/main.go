package main

import (
	"fmt"
	"leetcode/array/Easy/binary_search"
)

func main() {
	fmt.Println(binary_search.Search([]int{-1, 0, 3, 5, 9, 12, 15, 16, 17, 18}, 12))
	fmt.Println(binary_search.Search([]int{-1, 0, 3, 5, 9, 12}, 9))
	fmt.Println(binary_search.Search([]int{-1, 0, 3, 5, 9, 12, 22, 45, 76, 80, 112, 256, 678, 888, 889, 900, 931, 945}, 889))
	fmt.Println(binary_search.Search([]int{-1, 0, 3, 5, 9, 12}, 2))
}
