package factory

import "sort"

type TeamMember struct {
	TeamName string
	UserID   string
}

func TeamMemberUserIDs(teamMembers []TeamMember) []string {
	res := make([]string, 0, len(teamMembers))
	for _, teamMember := range teamMembers {
		res = append(res, teamMember.UserID)
	}
	return res
}

func TeamMemberSortByUserID(teamMembers []TeamMember) {
	sort.Slice(teamMembers, func(i, j int) bool {
		return teamMembers[i].UserID < teamMembers[j].UserID
	})
}
