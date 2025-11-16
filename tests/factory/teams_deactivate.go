package factory

import (
	"bytes"
	"encoding/json"
	"io"
)

type TeamsDeactivateRequest struct {
	TeamName string   `json:"team_name"`
	UserIDs  []string `json:"user_ids"`
}

func (r *TeamsDeactivateRequest) ToJSON() io.Reader {
	const operationPlace = "factory.teams.add.AddTeamRequest.ToJSON"
	body, err := json.Marshal(r)
	if err != nil {
		panic(err.Error() + " " + operationPlace)
	}
	return bytes.NewBuffer(body)
}
