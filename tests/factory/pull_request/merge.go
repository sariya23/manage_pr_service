package pull_request

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

func (r PullRequestMergeRequest) ToJson() io.Reader {
	const operationPlace = "factory.pull_request.merge.PullRequestMergeRequest.ToJson"
	body, err := json.Marshal(r)
	if err != nil {
		panic(err.Error() + " " + operationPlace)
	}
	return bytes.NewBuffer(body)
}

type PullRequestMergeResponsePullRequestDTO struct {
	PullRequestID     string    `json:"pull_request_id"`
	PullRequestName   int       `json:"pull_request_name"`
	AuthorID          int       `json:"author_id"`
	Status            string    `json:"status"`
	AssignedReviewers int       `json:"assigned_reviewers"`
	MergedAt          time.Time `json:"merged_at"`
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
