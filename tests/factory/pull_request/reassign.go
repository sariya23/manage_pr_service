//go:build integrations

package pull_request

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"slices"

	"github.com/sariya23/manage_pr_service/tests/helpers/random"
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
	PullRequestID     string   `json:"pull_request_id"`
	PullRequestName   string   `json:"pull_request_name"`
	AuthorID          string   `json:"author_id"`
	Status            string   `json:"status"`
	AssignedReviewers []string `json:"assigned_reviewers"`
	ReplacedBy        string   `json:"replaced_by"`
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

func MembersWithoutAuthorID(members []string, authorID string) []string {
	withoutAuthor := make([]string, 0, len(members)-1)

	for _, member := range members {
		if member != authorID {
			withoutAuthor = append(withoutAuthor, member)
		}
	}
	return withoutAuthor
}

func PickTeamUserButNotReviewer(members []string, reviewers []string) string {
	withoutReviewer := make([]string, 0, len(members)-len(reviewers))
	for _, member := range members {
		if !slices.Contains(reviewers, member) {
			withoutReviewer = append(withoutReviewer, member)
		}
	}
	return random.Choice(withoutReviewer)
}
