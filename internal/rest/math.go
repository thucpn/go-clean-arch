package rest

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// ParityChecker is the math capability this handler needs.
type ParityChecker interface {
	IsEven(n int64) bool
}

// MathHandler maps a number to "even" or "odd".
type MathHandler struct {
	Math ParityChecker
}

// NewMathHandler registers GET /parity/:n.
func NewMathHandler(e *echo.Echo, math ParityChecker) {
	handler := &MathHandler{Math: math}
	e.GET("/parity/:n", handler.Parity)
}

// Parity reads :n and returns "even" or "odd" from Math.IsEven.
func (h *MathHandler) Parity(c echo.Context) error {
	n, err := strconv.ParseInt(c.Param("n"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: "invalid number"})
	}
	if h.Math.IsEven(n) {
		return c.String(http.StatusOK, "even")
	}
	return c.String(http.StatusOK, "odd")
}
