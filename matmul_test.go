package main

import (
	"testing"
)

var A = Matrix{
	Vector{1, 2, 3, 4, 5},
	Vector{1, 2, 3, 4, 5},
	Vector{1, 2, 3, 4, 5},
	Vector{1, 2, 3, 4, 5},
	Vector{1, 2, 3, 4, 5},
}
var B = Matrix{
	Vector{1, 0, 0, 0, 0},
	Vector{0, 1, 0, 0, 0},
	Vector{0, 0, 1, 0, 0},
	Vector{0, 0, 0, 1, 0},
	Vector{0, 0, 0, 0, 1},
}

// go test -bench=. -benchtime=10 -cpu=1,12
var matrixSize int = 100

func getMatrix(b *testing.B, size int) Matrix {
	b.Helper()
	m := make(Matrix, size)
	for i := range m {
		m[i] = make(Vector, size)
		for j := range m[i] {
			m[i][j] = i * j
		}
	}
	return m
}

func BenchmarkMatMul1(b *testing.B) {
	m1 := getMatrix(b, matrixSize)
	m2 := getMatrix(b, matrixSize)

	for i := 0; i < b.N; i++ {
		MatMul1(m1, m2)
	}
}

func BenchmarkMatMul2(b *testing.B) {
	m1 := getMatrix(b, matrixSize)
	m2 := getMatrix(b, matrixSize)

	for i := 0; i < b.N; i++ {
		MatMul2(m1, m2)
	}
}

func BenchmarkMatMulCon(b *testing.B) {
	m1 := getMatrix(b, matrixSize)
	m2 := getMatrix(b, matrixSize)
	for i := 0; i < b.N; i++ {
		MatMulCon(m1, m2)
	}
}

func BenchmarkMatMulCon2(b *testing.B) {
	m1 := getMatrix(b, matrixSize)
	m2 := getMatrix(b, matrixSize)

	for i := 0; i < b.N; i++ {
		m1.MatMulCon2(m2)
	}

}
