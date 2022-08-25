package service

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/go-kit/log"
	"github.com/stretchr/testify/assert"
)

func TestHTTPServer(t *testing.T) {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stdout)
	logger = log.With(logger, "listen", "8081", "caller", log.DefaultCaller)
	s := NewService()
	r := NewHTTPServer(s, logger)
	srv := httptest.NewServer(r)

	for _, testcase := range []struct {
		method, url, body string
		want              int
	}{
		{"POST", srv.URL + "/hash", `{"string":"123456"}`, http.StatusOK},
		{"POST", srv.URL + "/hash", `{}`, http.StatusBadRequest},
	} {
		req, _ := http.NewRequest(testcase.method, testcase.url, strings.NewReader(testcase.body))
		resp, _ := http.DefaultClient.Do(req)
		assert.Equal(t, testcase.want, resp.StatusCode)
	}
}
