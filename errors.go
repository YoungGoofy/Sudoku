package main

import (
	"errors"
	"strings"
)

var BoundsError = errors.New("out of bounds")
var DigitError = errors.New("invalid digit")
var RowError = errors.New("find dublicate in row")
var ColumnError = errors.New("find dublicate in column")
var GridError = errors.New("find dublicate in grid")
var MainDigitError = errors.New("this place not empty")
var DelMainDigitError = errors.New("this digit is default")

func (se SudokuError) Error() string {
	var s []string
	for _, err := range se {
		s = append(s, err.Error())
	}
	return strings.Join(s, ", ")
}

func InBounds(x, y int) bool {
	if x < 0 || x > 9 {
		return false
	}
	if y < 0 || y > 9 {
		return false
	}
	return true
}

func ValueDigit(digit int8) bool {
	if digit < 1 || digit > 9 {
		return false
	}
	return true
}

// TODO:
// надо пройтись по "у" циклом, "х" остается неизменным
func CheckRows(field Field, column int, digit int8) bool {
	var items [Rows]int8
	for i := column; i < column+1; i++ {
		for j := 0; j < Rows; j++ {
			items[j] = field[i][j]
		}
	}
	return findDubles(items, digit)
}

// TODO:
// надо пройтись по "х" циклом, "у" остается неизменным
// надо добавить числа во временный массив и сверить
// данные с digit, если есть повтор, то false, иначе true
func CheckColumns(field Field, row int, digit int8) bool {
	var items [Columns]int8
	for i := 0; i < Columns; i++ {
		for j := row; j < row+1; j++ {
			items[i] = field[i][j]
		}
	}
	return findDubles(items, digit)
}

func CheckGrid(field Field, row, column int, digit int8) bool {
	var items [Rows]int8
	var itemCount int
	for x := (column / 3) * 3; x < ((column+3)/3)*3; x++ {
		for y := (row / 3) * 3; y < ((row+3)/3)*3; y++ {
			items[itemCount] = field[x][y]
			itemCount++
		}
	}
	return findDubles(items, digit)
}

// TODO:
// реализовать проверку, есть ли такой же элемент в массиве
func findDubles(arr [Rows]int8, digit int8) bool {
	for _, item := range arr {
		if digit == item {
			return false
		}
	}
	return true
}

func CheckMainDigits(field Field, row, column int) bool {
	if field[column][row] != 0 {
		return false
	}
	return true
}
