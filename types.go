package main

const Rows, Columns = 9, 9

type Field [Rows][Columns]int8

type SudokuError []error
