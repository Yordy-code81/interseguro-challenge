package matrix_utils

import (
	"errors"
	"math"
)

// FactorizeQR performs QR factorization on an m x n matrix A using the Modified Gram-Schmidt process.
// It returns Q (m x min(m,n)) and R (min(m,n) x n).
func FactorizeQR(A [][]float64) (Q [][]float64, R [][]float64, err error) {
	m := len(A)
	if m == 0 {
		return nil, nil, errors.New("matrix is empty")
	}
	n := len(A[0])
	if n == 0 {
		return nil, nil, errors.New("matrix has empty rows")
	}

	for _, row := range A {
		if len(row) != n {
			return nil, nil, errors.New("matrix is not rectangular")
		}
	}

	minDim := m
	if n < m {
		minDim = n
	}

	// Initialize Q (m x minDim)
	Q = make([][]float64, m)
	for i := 0; i < m; i++ {
		Q[i] = make([]float64, minDim)
	}

	// Initialize R (minDim x n)
	R = make([][]float64, minDim)
	for i := 0; i < minDim; i++ {
		R[i] = make([]float64, n)
	}

	// Make a copy of A into V to avoid mutating the original matrix
	V := make([][]float64, m)
	for i := 0; i < m; i++ {
		V[i] = make([]float64, n)
		copy(V[i], A[i])
	}

	for k := 0; k < minDim; k++ {
		// Calculate norm of V[:, k]
		norm := 0.0
		for i := 0; i < m; i++ {
			norm += V[i][k] * V[i][k]
		}
		norm = math.Sqrt(norm)

		if norm < 1e-10 {
			return nil, nil, errors.New("linearly dependent columns found or zero vector encountered")
		}

		R[k][k] = norm

		// Q[:, k] = V[:, k] / norm
		for i := 0; i < m; i++ {
			Q[i][k] = V[i][k] / norm
		}

		// Update remaining columns of V
		for j := k + 1; j < n; j++ {
			dot := 0.0
			for i := 0; i < m; i++ {
				dot += Q[i][k] * V[i][j]
			}
			R[k][j] = dot

			for i := 0; i < m; i++ {
				V[i][j] = V[i][j] - dot*Q[i][k]
			}
		}
	}

	return Q, R, nil
}
