package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestValidate(t *testing.T) {
	api := API{}

	testCases := []struct {
		name        string
		reqBody     []byte
		respBody    string
		method      string
		contentType string
		code        int
	}{
		{
			name:        "case success",
			reqBody:     []byte(`[ "[()]{}{[()()]()}", "[(])", "[({", ")}]", "[()]{}{[()()]()}}}}}}}}" ]`),
			respBody:    `{"success":true,"message":"","data":[{"value":"[()]{}{[()()]()}","is_valid":true},{"value":"[(])","is_valid":false},{"value":"[({","is_valid":false},{"value":")}]","is_valid":false},{"value":"[()]{}{[()()]()}}}}}}}}","is_valid":false}]}`,
			method:      "POST",
			contentType: "application/json",
			code:        http.StatusOK,
		},
		{
			name: "case array len more than 20",
			reqBody: []byte(`[ "{}", "[]", "{}", "[]", "{}", "()", "[]", "(][", ")}", ")(", "{}", "}{", ")}", "][",
							"()", "[]", "[]", "[]", "[]", "[]", "[]", "[]", "()", "()", ")(" ]`),
			respBody:    `{"success":false,"message":"more than 20 values"}`,
			method:      "POST",
			contentType: "application/json",
			code:        http.StatusBadRequest,
		},
		{
			name:        "case method is not POST",
			reqBody:     []byte(`[ "{}", "[]", "{}", "[]" ]`),
			respBody:    `{"success":false,"message":"only POST method is supported"}`,
			method:      "DELETE",
			contentType: "application/json",
			code:        http.StatusMethodNotAllowed,
		},
		{
			name:        "case wrong content-type",
			reqBody:     []byte(`[ "()" ]`),
			respBody:    `{"success":false,"message":"Content-Type should be an application/json"}`,
			method:      "POST",
			contentType: "multipart/form-data",
			code:        http.StatusUnsupportedMediaType,
		},
	}

	handler := http.HandlerFunc(api.validate)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(tc.method, "/api/v1/validate", bytes.NewReader(tc.reqBody))
			req.Header.Set("Content-Type", tc.contentType)

			handler.ServeHTTP(rec, req)

			if rec.Code != tc.code {
				t.Errorf("\nexpected: %d\nactual: %d", tc.code, rec.Code)
			}

			if strings.TrimSpace(rec.Body.String()) != tc.respBody {
				t.Errorf("\nexpected: %s\nactual: %s", tc.respBody, strings.TrimSpace(rec.Body.String()))
			}
		})
	}
}

func TestFix(t *testing.T) {
	api := API{}

	testCases := []struct {
		name        string
		reqBody     []byte
		respBody    string
		method      string
		contentType string
		code        int
	}{
		{
			name: "case array len more than 20",
			reqBody: []byte(`[ "{}", "[]", "{}", "[]", "{}", "()", "[]", "(][", ")}", ")(", "{}", "}{", ")}", "][",
							"()", "[]", "[]", "[]", "[]", "[]", "[]", "[]", "()", "()", ")(" ]`),
			respBody:    `{"success":false,"message":"more than 20 values"}`,
			method:      "POST",
			contentType: "application/json",
			code:        http.StatusBadRequest,
		},
		{
			name:        "case method is not POST",
			reqBody:     []byte(`[ "{}", "[]", "{}", "[]" ]`),
			respBody:    `{"success":false,"message":"only POST method is supported"}`,
			method:      "DELETE",
			contentType: "application/json",
			code:        http.StatusMethodNotAllowed,
		},
		{
			name:        "case wrong content-type",
			reqBody:     []byte(`[ "()" ]`),
			respBody:    `{"success":false,"message":"Content-Type should be an application/json"}`,
			method:      "POST",
			contentType: "multipart/form-data",
			code:        http.StatusUnsupportedMediaType,
		},
	}

	handler := http.HandlerFunc(api.fix)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(tc.method, "/api/v1/fix", bytes.NewReader(tc.reqBody))
			req.Header.Set("Content-Type", tc.contentType)

			handler.ServeHTTP(rec, req)

			if rec.Code != tc.code {
				t.Errorf("\nexpected: %d\nactual: %d", tc.code, rec.Code)
			}

			if strings.TrimSpace(rec.Body.String()) != tc.respBody {
				t.Errorf("\nexpected: %s\nactual: %s", tc.respBody, strings.TrimSpace(rec.Body.String()))
			}
		})
	}
}
