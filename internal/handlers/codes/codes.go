package codes

import api "github.com/sariya23/manage_pr_service/internal/generated"

const (
	INVALIDREQUEST           = "INVALID_REQUEST"
	NOTFOUND                 = api.NOTFOUND
	INTERNAL                 = "INTERNAL"
	TEAMALREADYEXISTS        = api.TEAMEXISTS
	PULLREQUESTALREADYEXISTS = api.PREXISTS
	PULLREQUESTMERGED        = api.PRMERGED
	USERNOTREVIEWER          = api.NOTASSIGNED
	NOCANDIDATE              = api.NOCANDIDATE
)
