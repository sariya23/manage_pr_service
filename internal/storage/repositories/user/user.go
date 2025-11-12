package repo_user

import (
	"log/slog"
	"strings"

	"github.com/sariya23/manage_pr_service/internal/storage/database"
)

const (
	UserTableUserIDField   = "user_id"
	UserTableUsernameField = "username"
	UserTableIsActiveField = "is_active"
	UserTableName          = "user"
)

var (
	UserAllFields = strings.Join([]string{UserTableUserIDField, UserTableUsernameField, UserTableIsActiveField}, ", ")
)

type UserRepository struct {
	conn *database.Database
	log  *slog.Logger
}

func NewUserRepository(conn *database.Database, log *slog.Logger) *UserRepository {
	return &UserRepository{
		conn: conn,
		log:  log,
	}
}
