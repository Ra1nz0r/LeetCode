package container_with_most_water

import (
	"fmt"
)

func MaxArea(height []int) int {
	frontPos := 0
	backPos := 0
	res := 1

	for k, v := range height {
		if v >= height[frontPos] {
			frontPos = k
			fmt.Println("perezapisali front na positiu", k)
		}

		if v > height[backPos] && backPos < frontPos {
			backPos = k
			fmt.Println("perezapisali back na positiu", k)
		}

		s := 1

		if height[frontPos] < height[backPos] {
			s = height[frontPos] * (frontPos - backPos)
			fmt.Println("Считаем значение для сравнения, когда передняя позиция меньше задней", s)
		} else if height[frontPos] > height[backPos] {
			s = height[backPos] * (frontPos - backPos)
			fmt.Println("Считаем значение для сравнения, когда задняя позиция меньше передней", s)
		}

		if s > res {
			res = s
			fmt.Println("Предварительный результат больше, чем имеем в итоге, перезаписываем", res)
		} else {
			frontPos = k
			fmt.Println("Предварительный результат меньше, чем имеем в итоге, обновляем позицию передней грани", frontPos)
		}

		fmt.Println("++++++++++++++++++++++++++++++++++++")
	}

	return res
}
