package factory

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type PullRequestMergeRequest struct {
	PullRequestID string `json:"pull_request_id"`
}

func (r PullRequestMergeRequest) ToJSON() io.Reader {
	const operationPlace = "factory.pull_request.merge.PullRequestMergeRequest.ToJSON"
	body, err := json.Marshal(r)
	if err != nil {
		panic(err.Error() + " " + operationPlace)
	}
	return bytes.NewBuffer(body)
}

type PullRequestMergeResponsePullRequestDTO struct {
	PullRequestID     string    `json:"pull_request_id"`
	PullRequestName   string    `json:"pull_request_name"`
	AuthorID          string    `json:"author_id"`
	Status            string    `json:"status"`
	AssignedReviewers []string  `json:"assigned_reviewers"`
	MergedAt          time.Time `json:"mergedAt"`
}

type PullRequestMergeResponse struct {
	PR PullRequestMergeResponsePullRequestDTO `json:"pr"`
}

func PullRequestMergeFromHTTPResponseOK(resp *http.Response) PullRequestMergeResponse {
	const operationPlace = "factory.pull_request.merge.PullRequestMergeFromHTTPResponseOK"
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error() + " " + operationPlace)
	}

	var result PullRequestMergeResponse
	if err := json.Unmarshal(body, &result); err != nil {
		panic(err.Error() + " " + operationPlace)
	}
	return result
}
