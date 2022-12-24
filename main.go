package main

import (
	"fmt"
)

type field Field

func Show(f field) {
	for x := 0; x < Columns; x++ {
		for y := 0; y < Rows; y++ {
			if y%3 == 0 {
				fmt.Print("| ")
			}
			fmt.Printf("%d ", f[x][y])
		}
		fmt.Println()
	}
}

func (f field) GenerateField() {
	Show(f)
}

func (f *field) AddNum() error {
	var x, y int
	var num int8
	fmt.Scan(&x, &y, &num)
	if !InBounds(x, y) {
		return BoundsError
	}
	if !ValueDigit(num) {
		return DigitError
	}
	f[y][x] = num
	return nil
}

func main() {
	var f field = [Rows][Columns]int8{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}
	f.GenerateField()
	err := f.AddNum()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	Show(f)
}
