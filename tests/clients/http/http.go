//go:build integrations

package httpcleint

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/sariya23/manage_pr_service/internal/config"
	"github.com/sariya23/manage_pr_service/tests/factory/teams"
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

func (c *HTTPClient) TeamsAdd(req teams.AddTeamRequest) *http.Response {
	reqJson := req.ToJson()
	resp, err := c.cl.Post(fmt.Sprintf("http://localhost:%d/api/team/add", c.port), "application/json", reqJson)
	if err != nil {
		panic(err)
	}
	return resp
}

func (c *HTTPClient) TeamGet(teamName string) *http.Response {
	resp, err := c.cl.Get(fmt.Sprintf("http://localhost:%d/api/team/get/%s", c.port, teamName))
	if err != nil {
		panic(err)
	}
	return resp
}
