package matrix_utils

import "errors"

// Rotate90DegreesRight rotates an m x n matrix 90 degrees to the right,
// returning an n x m matrix.
func Rotate90DegreesRight(matrix [][]float64) ([][]float64, error) {
	m := len(matrix)
	if m == 0 {
		return nil, errors.New("matrix is empty")
	}
	n := len(matrix[0])
	if n == 0 {
		return nil, errors.New("matrix has empty rows")
	}

	// Verify all rows have the same length
	for _, row := range matrix {
		if len(row) != n {
			return nil, errors.New("matrix is not rectangular")
		}
	}

	rotated := make([][]float64, n)
	for i := 0; i < n; i++ {
		rotated[i] = make([]float64, m)
		for j := 0; j < m; j++ {
			rotated[i][j] = matrix[m-1-j][i]
		}
	}

	return rotated, nil
}
