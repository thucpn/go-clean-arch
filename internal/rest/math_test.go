package rest_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/bxcodec/go-clean-arch/internal/rest"
	"github.com/bxcodec/go-clean-arch/internal/rest/mocks"
)

// TestParity_even checks that GET /parity/4 returns "even" when IsEven(4) is true.
// IsEven must be called exactly once.
func TestParity_even(t *testing.T) {
	mathSvc := mocks.NewMathService(t)
	mathSvc.EXPECT().IsEven(int64(4)).Return(true).Once()

	e := echo.New()
	rest.NewMathHandler(e, mathSvc)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/parity/4", nil)
	require.NoError(t, err)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "even", rec.Body.String())
}

// TestParity_odd checks that GET /parity/7 returns "odd" when IsEven(7) is false.
// IsEven must be called exactly once.
func TestParity_odd(t *testing.T) {
	mathSvc := mocks.NewMathService(t)
	mathSvc.EXPECT().IsEven(int64(7)).Return(false).Once()

	e := echo.New()
	rest.NewMathHandler(e, mathSvc)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/parity/7", nil)
	require.NoError(t, err)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "odd", rec.Body.String())
}

// TestParity_oddx2 calls GET /parity/7 twice and checks that IsEven(7) is called twice.
func TestParity_oddx2(t *testing.T) {
	mathSvc := mocks.NewMathService(t)
	mathSvc.EXPECT().IsEven(int64(7)).Return(false).Times(2)

	e := echo.New()
	rest.NewMathHandler(e, mathSvc)

	req1, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/parity/7", nil)
	require.NoError(t, err)
	rec1 := httptest.NewRecorder()
	e.ServeHTTP(rec1, req1)
	assert.Equal(t, http.StatusOK, rec1.Code)
	assert.Equal(t, "odd", rec1.Body.String())

	req2, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/parity/7", nil)
	require.NoError(t, err)
	rec2 := httptest.NewRecorder()
	e.ServeHTTP(rec2, req2)
	assert.Equal(t, http.StatusOK, rec2.Code)
	assert.Equal(t, "odd", rec2.Body.String())
}
