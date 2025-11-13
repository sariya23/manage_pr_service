package repo_user

import (
	"github.com/sariya23/manage_pr_service/internal/storage/database"
)

const (
	UserTableUserIDField    = "user_id"
	UserTableUsernameField  = "username"
	UserTableIsActiveField  = "is_active"
	UserTableCreatedAtField = "created_at"
	UserTableUpdatedAtField = "updated_at"
	UserTableName           = "user"
)

type UserRepository struct {
	conn *database.Database
}

func NewUserRepository(conn *database.Database) *UserRepository {
	return &UserRepository{
		conn: conn,
	}
}
