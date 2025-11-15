//go:build integrations

package httpcleint

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/sariya23/manage_pr_service/internal/config"
	"github.com/sariya23/manage_pr_service/tests/factory"
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

func (c *HTTPClient) TeamsAdd(req factory.AddTeamRequest) *http.Response {
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

func (c *HTTPClient) UsersSetIsActive(req factory.SetIsActiveRequest) *http.Response {
	reqJson := req.ToJson()
	resp, err := c.cl.Post(fmt.Sprintf("http://localhost:%d/api/users/setIsActive", c.port), "application/json", reqJson)
	if err != nil {
		panic(err)
	}
	return resp
}

func (c *HTTPClient) UsersGetReview(userID string) *http.Response {
	resp, err := c.cl.Get(fmt.Sprintf("http://localhost:%d/api/users/getReview/%s", c.port, userID))
	if err != nil {
		panic(err)
	}
	return resp
}

func (c *HTTPClient) PullRequestCreate(req factory.PullRequestCreateRequest) *http.Response {
	reqJson := req.ToJson()
	resp, err := c.cl.Post(fmt.Sprintf("http://localhost:%d/api/pullRequest/create", c.port), "application/json", reqJson)
	if err != nil {
		panic(err)
	}
	return resp
}

func (c *HTTPClient) PullRequestMerge(req factory.PullRequestMergeRequest) *http.Response {
	reqJson := req.ToJson()
	resp, err := c.cl.Post(fmt.Sprintf("http://localhost:%d/api/pullRequest/merge", c.port), "application/json", reqJson)
	if err != nil {
		panic(err)
	}
	return resp
}

func (c *HTTPClient) PullRequestReassign(req factory.PullRequestReassignRequest) *http.Response {
	reqJson := req.ToJson()
	resp, err := c.cl.Post(fmt.Sprintf("http://localhost:%d/api/pullRequest/reassign", c.port), "application/json", reqJson)
	if err != nil {
		panic(err)
	}
	return resp
}
