package vsslib

import (
	"gopkg.in/auth0.v5/management"
)

type Auth0RoleHandler interface {
	List() ([]Role, error)
}

type auth0role struct {
	session *management.Management
}

type Role struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewAuth0RoleSess(sess *management.Management) Auth0RoleHandler {
	a := &auth0role{session: sess}

	return a
}

func (a *auth0role) List() ([]Role, error) {
	var err error
	var role Role
	var roles []Role = []Role{}

	var page int = 0
	for {
		list, err := a.session.Role.List(
			management.Page(page),
		)
		if err != nil {
			return roles, err
		}

		for _, authRole := range list.Roles {
			role.ID = authRole.GetID()
			role.Name = authRole.GetName()
			role.Description = authRole.GetDescription()

			roles = append(roles, role)
		}

		if !list.HasNext() {
			break
		}
		page++
	}

	return roles, err
}