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

func (f *field) AddNum(count **int) error {
	var x, y int
	var num int8
	var errList SudokuError
	fmt.Scan(&x, &y, &num)
	if !InBounds(x, y) {
		errList = append(errList, BoundsError)
	}
	if !ValueDigit(num) {
		errList = append(errList, DigitError)
	}
	if !CheckRows(Field(*f), y, num) {
		errList = append(errList, RowError)
	}
	if !CheckColumns(Field(*f), x, num) {
		errList = append(errList, ColumnError)
	}
	if !CheckGrid(Field(*f), x, y, num) {
		errList = append(errList, GridError)
	}
	if !CheckMainDigits(Field(*f), x, y) {
		errList = append(errList, MainDigitError)
	}
	if len(errList) > 0 {
		return errList
	}
	f[y][x] = num
	**count++
	return nil
}

func (f *field) DeleteNum(checkNullsField Field, count **int) error {
	var x, y int
	var errList SudokuError
	fmt.Scan(&x, &y)
	if !CheckMainDigits(checkNullsField, x, y) {
		errList = append(errList, DelMainDigitError)
	}
	if len(errList) > 0 {
		return errList
	}
	f[y][x] = 0
	**count--
	return nil
}

func switchAddNum(f *field, count *int) {
	err := f.AddNum(&count)
	fmt.Print("\033[H\033[2J")
	if err != nil {
		if errs, ok := err.(SudokuError); ok {
			fmt.Printf("%d error(s) occurred:\n", len(errs))
			for _, e := range errs {
				fmt.Printf("- %v\n", e)
			}
		}
	}
	Show(*f)
}

func switchDelNum(f *field, nf field, count *int) {
	err := f.DeleteNum(Field(nf), &count)
	fmt.Print("\033[H\033[2J")
	if err != nil {
		if errs, ok := err.(SudokuError); ok {
			fmt.Printf("%d error(s) occurred:\n", len(errs))
			for _, e := range errs {
				fmt.Printf("- %v\n", e)
			}
		}
	}
	fmt.Println()
	Show(*f)
}

func checkNullsCount(f field) int {
	var count int
	for i := 0; i < Rows; i++ {
		for j := 0; j < Columns; j++ {
			if f[i][j] == 0 {
				count++
			}
		}
	}
	return count
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
	count := checkNullsCount(f)
	var checkNullsField = f
	Show(f)
	for count > 0 {
		fmt.Printf("\n1. Добавьте число\n2. Удалите число\n")
		var choise int
		fmt.Scan(&choise)
		switch choise {
		case 1:
			fmt.Printf("\nПравило ввода:\n1 число - координата по 'Х'\n2 число - координата по 'Y'\n3 число - число, которое вы хотите вставить\n")
			switchAddNum(&f, &count)
		case 2:
			fmt.Printf("\nПравило ввода:\n1 число - координата по 'Х'\n2 число - координата по 'Y'\n")
			switchDelNum(&f, checkNullsField, &count)
		default:
			fmt.Println("Нет такого пункта.")
		}
	}
}
