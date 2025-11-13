package domain

type User struct {
	UserID   string
	Username string
	TeamName string
	IsActive bool
}

func UserIDs(users []User) []string {
	ids := make([]string, 0, len(users))
	for _, user := range users {
		ids = append(ids, user.UserID)
	}
	return ids
}
