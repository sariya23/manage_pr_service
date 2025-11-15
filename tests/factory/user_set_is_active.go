package factory

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type SetIsActiveRequest struct {
	UserID   string `json:"user_id"`
	IsActive bool   `json:"is_active"`
}

func (r SetIsActiveRequest) ToJson() io.Reader {
	const operationPlace = "factory.users.set_is_active.SetIsActiveRequest.ToJson"
	body, err := json.Marshal(r)
	if err != nil {
		panic(err.Error() + " " + operationPlace)
	}
	return bytes.NewBuffer(body)
}

type SetIsActiveResponseUserDTO struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	TeamName string `json:"team_name"`
	IsActive bool   `json:"is_active"`
}

type SetIsActiveResponse struct {
	User SetIsActiveResponseUserDTO `json:"user"`
}

func SetIsActiveFromHTTPResponseOK(resp *http.Response) SetIsActiveResponse {
	const operationPlace = "factory.users.set_is_active.SetIsActiveFromHTTPResponseOK"
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error() + " " + operationPlace)
	}

	var result SetIsActiveResponse
	if err := json.Unmarshal(body, &result); err != nil {
		panic(err.Error() + " " + operationPlace)
	}
	return result
}
