package main

import (
	"math"
)

type Matrix struct {
	matrix [][]float32
}

func (*A Matrix) shape(){
	numRows = len(m.matrix)
	var numCols int
	switch numRows {
	case len(numRows) <1:
		numCols := 0
		
	default:
		numCols := len(m.matrix[0])
	}
	return numRows, numCols
	
}

func (*A Matrix, i int) row(){
	return A[i]
}

func (*A Matrix. i int) column(){
	columnData := make([]float32,0)
	for _, val := range A[i]{
		columnData = append(columnData,i)
	}
	return columnData
}