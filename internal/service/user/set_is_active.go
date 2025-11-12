package serviceuser

import (
	"context"

	"github.com/sariya23/manage_pr_service/internal/models"
)

func (s *UserService) SetIsActive(ctx context.Context, userId int64, isActive bool) (models.User, error) {
	panic("implement me")
}
