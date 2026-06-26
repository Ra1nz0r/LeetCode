package main

import (
	"fmt"
	"leetcode/array/Easy/majority_element"
)

func main() {
	fmt.Println(majority_element.MajorityElementBoyerMooreType([]int{3, 2, 3}))
	fmt.Println(majority_element.MajorityElementBoyerMooreType([]int{2, 2, 1, 1, 1, 2, 2}))
	fmt.Println(majority_element.MajorityElementBoyerMooreType([]int{1}))
}
