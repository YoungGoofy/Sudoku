package main

import (
	"errors"

	types "github.com/YoungGoofy/sudoku/types"
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

func RepeatDigits(field types.Field) {

}
