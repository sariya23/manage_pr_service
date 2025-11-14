package handlers

import api "github.com/sariya23/manage_pr_service/internal/generated"

const (
	INVALID_REQUEST             = "INVALID_REQUEST"
	NOT_FOUND                   = api.NOTFOUND
	INTERNAL                    = "INTERNAL"
	INVALID_JSON                = "INVALID_JSON"
	TEAM_ALREADY_EXISTS         = api.TEAMEXISTS
	PULL_REQUEST_ALREADY_EXISTS = api.PREXISTS
	PULL_REQUEST_MERGED         = api.PRMERGED
	USER_NOT_REVIEWER           = api.NOTASSIGNED
	NO_CANDIDATE                = api.NOCANDIDATE
)
