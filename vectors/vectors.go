package main

import (
	"fmt";
	"errors"
)

func vectorAdd(v1,v2  []float32) ([]float32,error){
	if len(v1) != len(v2){
		return []float32{-1}, errors.New("Vectors must be the same length")
	}
	resultVector := make([]float32,0)
	
	for idx,_ := range v1{
		elem1 := v1[idx]
		elem2 := v2[idx]
		elemSum := elem1 + elem2

		resultVector = append(resultVector,elemSum)
	}
	return resultVector,nil

}

func vectorSub(v1,v2 []float32)([]float32,error){
	if len(v1) != len(v2){
		return []float32{-1}, errors.New("Vectors must be the same length")
	}

	resultVector := make([]float32,0)

	for idx,_ := range v1{
		elem1 :=  v1[idx]
		elem2 := v2[idx]
		elemDiff := elem1-elem2
		
		resultVector = append(resultVector,elemDiff)
	}
	return resultVector,nil

}

func vectorSum(vectors [][]float32)([]float32){
	initialVector := vectors[0]
	
	for i:= 1; i<len(vectors); i++ {
		initialVector,_ = vectorAdd(initialVector,vectors[i])
	}
	return initialVector

}

func scalarMultiply(c float32, v []float32)([]float32){
	resultVector := make([]float32,0)
	for idx,_ := range v{
		result := c * v[idx]
		resultVector = append(resultVector,result)
	}
	return resultVector
}

func vectorMean(vectors [][]float32) ([]float32){
	n := float32(len(vectors))
	return scalarMultiply(1.0/n,vectorSum(vectors))
}
