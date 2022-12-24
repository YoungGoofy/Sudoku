package main

import (
	"errors"
)

var BoundsError = errors.New("out of bounds")
var DigitError = errors.New("invalid digit")

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
func CheckRows(field Field, row, column int, digit int8) bool {
	return true
}

// TODO:
// надо пройтись по "х" циклом, "у" остается неизменным
// надо добавить числа во временный массив и сверить
// данные с digit, если есть повтор, то false, иначе true
func CheckColumns(field Field, row, column int, digit int8) bool {
	return true
}

func CheckGrid(field Field, row, column int, digit int8) bool {
	return true
}

// TODO:
// реализовать проверку, есть ли такой же элемент в массиве
func findDubles(arr [Rows]int8, row, column int, digit int8) bool {
	return true
}
