//go:build integrations

package httpcleint

import (
	"net/http"
	"path/filepath"
	"time"

	"github.com/sariya23/manage_pr_service/internal/config"
)

type HTTPClient struct {
	cl   *http.Client
	port int
}

func NewHTTPClient() *HTTPClient {
	cfg := config.MustLoadByPath(filepath.Join("..", "..", "..", "..", "config", "test.env"))
	cl := &http.Client{Timeout: time.Second}
	return &HTTPClient{cl: cl, port: cfg.HTTPServerPort}
}

func (c *HTTPClient) GetClient() *http.Client {
	return c.cl
}
