package main

import (
	"fmt"
)

func gridTravel(row, col int) int {
	if col == 1 && row == 1 {
		return 1
	}

	if col == 0 || row == 0 {
		return 0
	}

	return gridTravel(row-1, col) + gridTravel(row, col-1)
}

func gridTravelMemo(row, col int, mem map[string]int) int {
	key := fmt.Sprintf("%d,%d", row, col)
	if v, ok := mem[key]; ok {
		return v
	}

	if col == 1 && row == 1 {
		return 1
	}

	if col == 0 || row == 0 {
		return 0
	}

	r := gridTravelMemo(row-1, col, mem) + gridTravelMemo(row, col-1, mem)

	mem[key] = r

	return r

}

func gridTravelTabu(row, col int) int {
	table := make([][]int, row+1)

	for i := range table {
		table[i] = make([]int, col+1)
	}

	table[1][1] = 1

	for i := 0; i <= row; i++ {
		for j := 0; j <= col; j++ {
			current := table[i][j]
			if i+1 <= row {
				table[i+1][j] += current
			}

			if j+1 <= col {
				table[i][j+1] += current
			}
		}
	}

	return table[row][col]
}

func main() {
	r := gridTravel(1, 1)

	fmt.Println(r)

	r = gridTravel(2, 3)

	fmt.Println(r)

	r = gridTravel(3, 2)

	fmt.Println(r)

	r = gridTravel(3, 3)

	fmt.Println(r)

	mem := make(map[string]int)

	r = gridTravelMemo(18, 18, mem)

	fmt.Println(r)

	fmt.Println(gridTravelTabu(18, 18))

}
