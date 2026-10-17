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
)

// parityStub is a hand-written stand-in for MathService. No mockery.
type parityStub struct {
	isEven bool
	got    int64
}

func (s *parityStub) IsEven(n int64) bool {
	s.got = n
	return s.isEven
}

// TestParityRaw_even checks that the handler maps IsEven=true to the body "even".
func TestParityRaw_even(t *testing.T) {
	stub := &parityStub{isEven: true}
	e := echo.New()
	rest.NewMathHandler(e, stub)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/parity/4", nil)
	require.NoError(t, err)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "even", rec.Body.String())
	assert.Equal(t, int64(4), stub.got)
}

// TestParityRaw_odd checks that the handler maps IsEven=false to the body "odd".
func TestParityRaw_odd(t *testing.T) {
	stub := &parityStub{isEven: false}
	e := echo.New()
	rest.NewMathHandler(e, stub)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/parity/7", nil)
	require.NoError(t, err)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "odd", rec.Body.String())
	assert.Equal(t, int64(7), stub.got)
}
