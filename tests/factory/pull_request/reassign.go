//go:build integrations

package pull_request

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type PullRequestReassignRequest struct {
	PullRequestID string `json:"pull_request_id"`
	OldUserID     string `json:"old_user_id"`
}

func (r PullRequestReassignRequest) ToJson() io.Reader {
	const operationPlace = "factory.pull_request.reassign.PullRequestReassignRequest.ToJson"
	body, err := json.Marshal(r)
	if err != nil {
		panic(err.Error() + " " + operationPlace)
	}
	return bytes.NewBuffer(body)
}

type PullRequestReassignResponsePullRequestDTO struct {
	PullRequestID     string `json:"pull_request_id"`
	PullRequestName   int    `json:"pull_request_name"`
	AuthorID          int    `json:"author_id"`
	Status            string `json:"status"`
	AssignedReviewers int    `json:"assigned_reviewers"`
	ReplacedBy        string `json:"replaced_by"`
}

type PullRequestReassignResponse struct {
	PR PullRequestReassignResponsePullRequestDTO `json:"pr"`
}

func PullRequestReassignFromHTTPResponseOK(resp *http.Response) PullRequestReassignResponse {
	const operationPlace = "factory.pull_request.reassign.PullRequestReassignFromHTTPResponseOK"
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error() + " " + operationPlace)
	}

	var result PullRequestReassignResponse
	if err := json.Unmarshal(body, &result); err != nil {
		panic(err.Error() + " " + operationPlace)
	}
	return result
}
