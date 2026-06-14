package handler

import (
	"go-matrix-api/internal/domain"
	"go-matrix-api/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type MatrixHandler struct {
	usecase usecase.MatrixUsecase
}

func NewMatrixHandler(uc usecase.MatrixUsecase) *MatrixHandler {
	return &MatrixHandler{usecase: uc}
}

// ProcessMatrix godoc
// @Summary Process a matrix
// @Description Rotates a matrix, performs QR factorization, and retrieves statistics from Node.js API
// @Tags Matrix
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body domain.MatrixRequest true "Matrix data"
// @Success 200 {object} domain.BaseResponse
// @Failure 400 {object} domain.BaseResponse
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} domain.BaseResponse
// @Router /matrix [post]
func (h *MatrixHandler) ProcessMatrix(c *fiber.Ctx) error {
	var req domain.MatrixRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.BaseResponse{
			Success: false,
			Error:   "Invalid request body format",
		})
	}

	if len(req.Matrix) == 0 || len(req.Matrix[0]) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(domain.BaseResponse{
			Success: false,
			Error:   "Matrix cannot be empty",
		})
	}

	token := c.Get("Authorization")

	res, err := h.usecase.ProcessMatrix(&req, token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.BaseResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(domain.BaseResponse{
		Success: true,
		Data:    res,
	})
}
