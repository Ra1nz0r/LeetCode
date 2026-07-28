package main

import (
	"fmt"
	"leetcode/array/Medium/group_anagrams"
)

func main() {
	fmt.Println("Ввели [`eat`,`tea`,`tan`,`ate`,`nat`,`bat`],  ожидаем ответ [[`bat`],[`nat`,`tan`],[`ate`,`eat`,`tea`]]")
	fmt.Println(group_anagrams.GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))

	fmt.Println("Ввели [``],  ожидаем ответ [[``]]")
	fmt.Println(group_anagrams.GroupAnagrams([]string{""}))

	fmt.Println("Ввели [`a`],  ожидаем ответ [[`a`]]")
	fmt.Println(group_anagrams.GroupAnagrams([]string{"a"}))
}
