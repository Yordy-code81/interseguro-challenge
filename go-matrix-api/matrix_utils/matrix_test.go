package matrix_utils

import (
	"math"
	"testing"
)

func assertFloatEqual(t *testing.T, a, b float64) {
	if math.Abs(a-b) > 1e-6 {
		t.Errorf("Expected %v, got %v", a, b)
	}
}

func TestRotate90DegreesRight(t *testing.T) {
	matrix := [][]float64{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	expected := [][]float64{
		{5, 3, 1},
		{6, 4, 2},
	}

	rotated, err := Rotate90DegreesRight(matrix)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	for i := range expected {
		for j := range expected[i] {
			if rotated[i][j] != expected[i][j] {
				t.Errorf("At [%d][%d]: expected %v, got %v", i, j, expected[i][j], rotated[i][j])
			}
		}
	}
}

func TestFactorizeQR_Square(t *testing.T) {
	A := [][]float64{
		{12, -51, 4},
		{6, 167, -68},
		{-4, 24, -41},
	}
	
	Q, R, err := FactorizeQR(A)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Verify Q * R ≈ A
	m := len(Q)
	n := len(R[0])
	minDim := len(R)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum := 0.0
			for k := 0; k < minDim; k++ {
				sum += Q[i][k] * R[k][j]
			}
			assertFloatEqual(t, A[i][j], sum)
		}
	}

	// Verify Q is orthogonal (Q^T * Q = I)
	for i := 0; i < minDim; i++ {
		for j := 0; j < minDim; j++ {
			sum := 0.0
			for k := 0; k < m; k++ {
				sum += Q[k][i] * Q[k][j]
			}
			expected := 0.0
			if i == j {
				expected = 1.0
			}
			assertFloatEqual(t, expected, sum)
		}
	}
}

func TestFactorizeQR_Rectangular(t *testing.T) {
	A := [][]float64{
		{1, 1},
		{2, 1},
		{1, 2},
	}

	Q, R, err := FactorizeQR(A)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Verify Q * R ≈ A
	m := len(Q)
	n := len(R[0])
	minDim := len(R)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum := 0.0
			for k := 0; k < minDim; k++ {
				sum += Q[i][k] * R[k][j]
			}
			assertFloatEqual(t, A[i][j], sum)
		}
	}
}
