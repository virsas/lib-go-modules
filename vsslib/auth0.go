package vsslib

import (
	"errors"

	"github.com/virsas/lib-go-modules/vssutil"
	"gopkg.in/auth0.v5"
	"gopkg.in/auth0.v5/management"
)

type Auth0Handler interface {
	RoleList() ([]Auth0Role, error)
	RoleShow(id string) (Auth0Role, error)
	RoleCreate(name string, description string) error
	RoleUpdate(id string, name string, description string) error
	RoleDelete(id string) error
	UserList() ([]*management.User, error)
	UserShow(id string) (*management.User, error)
	UserCreate(name string, email string, connection string) error
	UserUpdate(id string, name string, description string) error
	UserBlock(id string) error
	UserUnblock(id string) error
	UserDelete(id string) error
}

type auth0Struct struct {
	session *management.Management
}

type Auth0Role struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	AuthUsers   []*management.User `json:"authUsers"`
}

func NewAuth0Session(domain string, client string, secret string) (Auth0Handler, error) {
	var err error
	var session *management.Management

	session, err = management.New(domain, management.WithClientCredentials(client, secret))
	if err != nil {
		return nil, err
	}

	a := &auth0Struct{session: session}

	return a, nil
}

func (a *auth0Struct) RoleList() ([]Auth0Role, error) {
	var err error
	var role Auth0Role
	var roles []Auth0Role = []Auth0Role{}

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

func (a *auth0Struct) RoleShow(id string) (Auth0Role, error) {
	var err error
	var role Auth0Role

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

func (a *auth0Struct) RoleCreate(name string, description string) error {
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

func (a *auth0Struct) RoleUpdate(id string, name string, description string) error {
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

func (a *auth0Struct) RoleDelete(id string) error {
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

func (a *auth0Struct) UserList() ([]*management.User, error) {
	var err error
	var users []*management.User = []*management.User{}

	var page int = 0
	for {
		list, err := a.session.User.List(
			management.Page(page),
		)
		if err != nil {
			return users, err
		}

		users = append(users, list.Users...)

		if !list.HasNext() {
			break
		}
		page++
	}

	return users, err
}

func (a *auth0Struct) UserShow(id string) (*management.User, error) {
	var err error
	var user *management.User

	user, err = a.session.User.Read(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (a *auth0Struct) UserCreate(name string, email string, connection string) error {
	var err error

	u := &management.User{
		Name:        auth0.String(name),
		Email:       auth0.String(email),
		Password:    auth0.String(vssutil.RandomString(32, "v1*:")),
		Connection:  auth0.String(connection),
		VerifyEmail: auth0.Bool(false),
	}

	err = a.session.User.Create(u)
	if err != nil {
		return err
	}

	return nil
}

func (a *auth0Struct) UserUpdate(id string, name string, email string) error {
	var err error

	u := &management.User{
		Name:  auth0.String(name),
		Email: auth0.String(email),
	}

	err = a.session.User.Update(id, u)
	if err != nil {
		return err
	}

	return nil
}

func (a *auth0Struct) UserBlock(id string) error {
	var err error

	u := &management.User{
		Blocked: auth0.Bool(true),
	}

	err = a.session.User.Update(id, u)
	if err != nil {
		return err
	}

	return nil
}

func (a *auth0Struct) UserUnblock(id string) error {
	var err error

	u := &management.User{
		Blocked: auth0.Bool(false),
	}

	err = a.session.User.Update(id, u)
	if err != nil {
		return err
	}

	return nil
}

func (a *auth0Struct) UserDelete(id string) error {
	err := a.session.User.Delete(id)
	if err != nil {
		return err
	}

	return nil
}