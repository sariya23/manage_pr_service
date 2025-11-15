package factory

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/brianvoe/gofakeit/v7"
)

type PullRequestCreateRequest struct {
	PullRequestID   string `json:"pull_request_id"`
	PullRequestName string `json:"pull_request_name"`
	AuthorID        string `json:"author_id"`
}

func (pr *PullRequestCreateRequest) RadnomInit(prID string, prName string, authorID string) {
	if prID == "" {
		prID = gofakeit.LetterN(8)
	}

	if prName == "" {
		prName = gofakeit.AppName()
	}

	if authorID == "" {
		authorID = gofakeit.LetterN(8)
	}

	pr.PullRequestID = prID
	pr.PullRequestName = prName
	pr.AuthorID = authorID
}

func (r PullRequestCreateRequest) ToJson() io.Reader {
	const operationPlace = "factory.pull_request.create.PullRequestCreateRequest.ToJson"
	body, err := json.Marshal(r)
	if err != nil {
		panic(err.Error() + " " + operationPlace)
	}
	return bytes.NewBuffer(body)
}

type PullRequestCreateResponsePullRequestDTO struct {
	PullRequestID     string   `json:"pull_request_id"`
	PullRequestName   string   `json:"pull_request_name"`
	AuthorID          string   `json:"author_id"`
	Status            string   `json:"status"`
	AssignedReviewers []string `json:"assigned_reviewers"`
}

type PullRequestCreateResponse struct {
	PR PullRequestCreateResponsePullRequestDTO `json:"pr"`
}

func PullRequestCreateFromHTTPResponseOK(resp *http.Response) PullRequestCreateResponse {
	const operationPlace = "factory.pull_request.create.PullRequestCreateFromHTTPResponseOK"
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error() + " " + operationPlace)
	}

	var result PullRequestCreateResponse
	if err := json.Unmarshal(body, &result); err != nil {
		panic(err.Error() + " " + operationPlace)
	}
	return result
}
