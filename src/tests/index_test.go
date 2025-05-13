package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/mikietechie/gopeople/api/app"
	"github.com/stretchr/testify/assert"
)

func TestIndexRoute(t *testing.T) {
	testcases := []struct {
		route        string
		expectedCode int
		desc         string
	}{
		{
			route:        "/",
			expectedCode: fiber.StatusOK,
			desc:         "Index should always return ok",
		},
	}
	app := app.New()
	for _, tc := range testcases {
		req := httptest.NewRequest("GET", tc.route, nil)
		res, _ := app.Test(req, 3)
		assert.Equal(t, tc.expectedCode, res.StatusCode, tc.desc)
	}
}
