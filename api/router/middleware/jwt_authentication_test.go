package middleware_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"merchant/api/router/middleware"
)

const (
	testHandlerRespBody = "{\"message\":\"Hello World!\"}"

	tokenExpired = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNmNmYjVlM2YtZDZkNC00Y2ZmLThmOTItMTY4YTQ0Y2E1M2U5IiwiZW1haWwiOiJub3JtYW5AYWxwaGFuZXR3b3Jrcy5jb20uc2ciLCJleHAiOjE1OTExNDg5NDd9.gNipvogAXjPxfnV95fJrCRMJW5XrEfQ9LlRJzh9L3Vg"
	// tokenValid   = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNmNmYjVlM2YtZDZkNC00Y2ZmLThmOTItMTY4YTQ0Y2E1M2U5IiwiZW1haWwiOiJub3JtYW5AYWxwaGFuZXR3b3Jrcy5jb20uc2ciLCJleHAiOjB9.uTNVB-tB5SdrGhgUhf1Tty5UprfQucARTjHPbGWfyc0"
)

type jwtAuthTestCase struct {
	name         string
	token        string
	expectedResp int
}

var tests = []*jwtAuthTestCase{
	{
		name:         "with no token",
		expectedResp: http.StatusUnauthorized,
	}, {
		name:         "with an invalid token",
		token:        "invalid-token",
		expectedResp: http.StatusUnauthorized,
	}, {
		name:         "with an expired token",
		token:        tokenExpired,
		expectedResp: http.StatusUnauthorized,
	},
	// TODO: FIX with valid token
	/*{
		name:         "with a valid token",
		token:        tokenValid,
		expectedResp: http.StatusOK,
	},*/
}

func TestJwtAuthentication(t *testing.T) {
	for _, tc := range tests {
		tc := tc

		r := httptest.NewRequest("GET", "/", nil)
		rr := httptest.NewRecorder()

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if tc.token != "" {
				r.Header.Set("Authorization", fmt.Sprintf("BEARER %s", tc.token))
			}

			middleware.JwtAuthentication(http.HandlerFunc(jwtTestHandlerFunc())).ServeHTTP(rr, r)

			resp := rr.Result().StatusCode
			if tc.expectedResp != resp {
				t.Errorf("Wrong response code: want %d, got %d ", tc.expectedResp, resp)
			}
		})
	}
}

func jwtTestHandlerFunc() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, testHandlerRespBody)
	}
}
