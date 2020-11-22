package utils

import (
	"strings"

	"github.com/iduslab/backend/db"
	"github.com/iduslab/backend/models"
)

func GetMemberRole() (string, error) {
	return db.GetSetting("memberRole")
}

func GetRoles() (*[]models.Role, error) {
	rolesStr, err := db.GetSetting("roles")
	if err != nil {
		return nil, err
	}
	rolesSplit := strings.Split(rolesStr, "|")
	var roles []models.Role
	for _, d := range rolesSplit {
		split := strings.Split(d, ",")
		roles = append(roles, models.Role{
			Name:   split[1],
			RoleID: split[0],
		})
	}

	return &roles, nil
}
