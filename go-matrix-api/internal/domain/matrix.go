package domain

// MatrixRequest represents the incoming request with the initial matrix.
type MatrixRequest struct {
	Matrix [][]float64 `json:"matrix" validate:"required"`
}

// QRFactorization holds the resulting Q and R matrices.
type QRFactorization struct {
	Q [][]float64 `json:"Q"`
	R [][]float64 `json:"R"`
}

// NodeStatistics represents the stats returned by the Node.js API.
type NodeStatistics struct {
	MaxValue   float64         `json:"max_value" example:"6"`
	MinValue   float64         `json:"min_value" example:"-7.81"`
	Average    float64         `json:"average" example:"0.45"`
	TotalSum   float64         `json:"total_sum" example:"2.56"`
	IsDiagonal map[string]bool `json:"is_diagonal" example:"Q:false,R:false"`
}

// MatrixResponse is the consolidated response returned to the client.
type MatrixResponse struct {
	OriginalMatrix  [][]float64     `json:"original_matrix"`
	RotatedMatrix   [][]float64     `json:"rotated_matrix"`
	QRFactorization QRFactorization `json:"qr_factorization"`
	Statistics      *NodeStatistics `json:"statistics,omitempty"`
}

// BaseResponse wrapper for API responses
type BaseResponse struct {
	Success bool            `json:"success" example:"true"`
	Data    *MatrixResponse `json:"data,omitempty"`
	Error   string          `json:"error,omitempty" example:"matrix cannot be empty"`
}

// NodeClient defines the interface for communicating with the Node.js API
type NodeClient interface {
	GetStatistics(q, r [][]float64, token string) (*NodeStatistics, error)
}
