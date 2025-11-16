package request

import (
	"context"

	"github.com/sariya23/manage_pr_service/internal/middleware"
)

func GetIDKey(ctx context.Context) string {
	requestID, ok := ctx.Value(middleware.RequestIDKey).(string)
	if !ok {
		return "-1"
	}
	return requestID
}
