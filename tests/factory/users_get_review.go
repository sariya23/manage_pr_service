package factory

import (
	"encoding/json"
	"io"
	"net/http"
)

type UsersGetReviewResponsePullRequestDTO struct {
	PullRequestID   string `json:"pull_request_id"`
	PullRequestName string `json:"pull_request_name"`
	AuthorID        string `json:"author_id"`
	Status          string `json:"status"`
}

type UsersGetReviewResponse struct {
	UserID       string                                 `json:"user_id"`
	PullRequests []UsersGetReviewResponsePullRequestDTO `json:"pull_requests"`
}

func GetReviewFromHTTPResponseOK(resp *http.Response) UsersGetReviewResponse {
	const operationPlace = "factory.users.set_is_active.SetIsActiveFromHTTPResponseOK"
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error() + " " + operationPlace)
	}

	var result UsersGetReviewResponse
	if err := json.Unmarshal(body, &result); err != nil {
		panic(err.Error() + " " + operationPlace)
	}
	return result
}
