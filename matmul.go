package main

import (
	"errors"
	"sync"
)

type Vector []int
type Matrix []Vector

func MatMul1(A, B Matrix) (Matrix, error) {
	// Get dimensions.
	Arows := len(A)
	Acols := len(A[0])
	Brows := len(B)
	Bcols := len(B[0])

	if Acols != Brows {
		return Matrix{}, errors.New("Cols A different from Rows B.")
	}

	// Initialize the result matrix with zeros
	result := make(Matrix, Arows)
	for i := range result {
		result[i] = make(Vector, Bcols)
	}

	for i := 0; i < Arows; i++ {
		for j := 0; j < Bcols; j++ {
			sum := 0
			for k := 0; k < Acols; k++ {
				sum += A[i][k] * B[k][j]
			}
			result[i][j] = sum
		}
	}
	return result, nil
}

func MatMul2(A, B Matrix) (Matrix, error) {
	// Get dimensions.
	Arows := len(A)
	Acols := len(A[0])
	Brows := len(B)
	Bcols := len(B[0])

	if Acols != Brows {
		return Matrix{}, errors.New("Cols A different from Rows B.")
	}

	// Initialize the result matrix with zeros
	result := make(Matrix, Arows)

	for i := 0; i < Arows; i++ {
		result[i] = make(Vector, Bcols)
		for j := 0; j < Bcols; j++ {
			sum := 0
			for k := 0; k < Acols; k++ {
				sum += A[i][k] * B[k][j]
			}
			result[i][j] = sum
		}
	}
	return result, nil
}

// MatMulCon performs matrix multiplication of A and B concurrently.
func MatMulCon(A, B Matrix) (Matrix, error) {
	// Get dimensions of matrices
	Arows := len(A)
	Acols := len(A[0])
	Brows := len(B)
	Bcols := len(B[0])

	if Acols != Brows {
		return Matrix{}, errors.New("Cols A different from Rows B.")
	}

	// Initialize the result matrix with zeros
	result := make(Matrix, Arows)
	for i := range result {
		result[i] = make(Vector, Bcols)
	}

	// Use a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Function to compute a single cell in the result matrix
	computeCell := func(row, col int) {
		defer wg.Done()
		sum := 0
		for k := 0; k < Acols; k++ {
			sum += A[row][k] * B[k][col]
		}
		result[row][col] = sum
	}

	// Launch a goroutine for each cell in the result matrix
	for i := 0; i < Arows; i++ {
		for j := 0; j < Bcols; j++ {
			wg.Add(1)
			go computeCell(i, j)
		}
	}

	// Wait for all goroutines to finish
	wg.Wait()
	return result, nil
}

func (a Matrix) MatMulCon2(b Matrix) (Matrix, error) {
	if a == nil || b == nil {
		return nil, errors.New("Nil matrices not allowed")
	}

	if len(a[0]) != len(b) {
		return nil, errors.New("Inner dimensions must be equal")
	}

	res := make(Matrix, len(a))

	var wg sync.WaitGroup

	for i := range a {
		res[i] = make(Vector, len(b[0]))
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := range b[0] {
				res[i][j] = 0
				for k := range b {
					res[i][j] += a[i][k] * b[k][j]
				}
			}
		}(i)
	}
	wg.Wait()

	return res, nil
}
