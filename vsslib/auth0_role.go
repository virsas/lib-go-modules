package vsslib

import (
	"errors"

	"gopkg.in/auth0.v5"
	"gopkg.in/auth0.v5/management"
)

type Auth0RoleHandler interface {
	List() ([]Role, error)
	Show(id string) (Role, error)
	Create(name string, description string) error
	Update(id string, name string, description string) error
	Delete(id string) error
}

type auth0role struct {
	session *management.Management
}

type Role struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	AuthUsers   []*management.User `json:"authUsers"`
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

func (a *auth0role) Show(id string) (Role, error) {
	var err error
	var role Role

	authRole, err := a.session.Role.Read(id)
	if err != nil {
		return role, err
	}

	role.ID = authRole.GetID()
	role.Name = authRole.GetName()
	role.Description = authRole.GetDescription()

	users, err := a.session.Role.Users(authRole.GetID())
	if err != nil {
		return role, err
	}
	role.AuthUsers = users.Users

	return role, nil
}

func (a *auth0role) Create(name string, description string) error {
	var err error

	r := &management.Role{
		Name:        auth0.String(name),
		Description: auth0.String(description),
	}

	err = a.session.Role.Create(r)
	if err != nil {
		return err
	}

	return nil
}

func (a *auth0role) Update(id string, name string, description string) error {
	var err error

	r := &management.Role{
		Name:        auth0.String(name),
		Description: auth0.String(description),
	}

	err = a.session.Role.Update(id, r)
	if err != nil {
		return err
	}

	return nil
}

func (a *auth0role) Delete(id string) error {
	var err error

	users, err := a.session.Role.Users(id)
	if err != nil {
		return err
	}

	if len(users.Users) > 0 {
		return errors.New("notEmpty")
	}

	err = a.session.Role.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
