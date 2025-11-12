package handlers

import api "github.com/sariya23/manage_pr_service/internal/generated"

const (
	MISSING_REQUIRED_FIELD = "MISSSING_REQUIRED_FIELD"
	TYPE_MISMATCH          = "TYPE_MISMATCH"
	NOT_FOUND              = api.NOTFOUND
	INTERNAL               = "INTERNAL"
	INVALID_JSON           = "INVALID_JSON"
)
