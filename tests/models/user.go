//go:build integrations

package models

import (
	"sort"
	"time"
)

type User struct {
	UserID    string
	Username  string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func UserSortByUserID(users []User) {
	sort.Slice(users, func(i, j int) bool {
		return users[i].UserID < users[j].UserID
	})
}
