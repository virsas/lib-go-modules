package vsslib

import (
	"errors"
	"net/http"
	"strings"

	"github.com/virsas/lib-go-modules/vssutil"
	"gopkg.in/auth0.v5"
	"gopkg.in/auth0.v5/management"
)

type Auth0Handler interface {
	RoleList() ([]*management.Role, error)
	RoleShow(id string) (*management.Role, *management.UserList, error)
	RoleCreate(name string, description string) error
	RoleUpdate(id string, name string, description string) error
	RoleDelete(id string) error
	UserList() ([]*management.User, error)
	UserShow(id string) (*management.User, error)
	UserCreate(name string, email string) error
	UserUpdate(id string, name string, email string, passreset bool) error
	UserBlock(id string) error
	UserDelete(id string) error
	PassReset(email string) error
}

type auth0Struct struct {
	session    *management.Management
	client     string
	secret     string
	domain     string
	connection string
}

func NewAuth0Session(domain string, client string, secret string, connection string) (Auth0Handler, error) {
	var err error
	var session *management.Management

	session, err = management.New(domain, management.WithClientCredentials(client, secret))
	if err != nil {
		return nil, err
	}

	a := &auth0Struct{session: session, client: client, secret: secret, domain: domain, connection: connection}

	return a, nil
}

func (a *auth0Struct) RoleList() ([]*management.Role, error) {
	var err error
	var roles []*management.Role = []*management.Role{}

	var page int = 0
	for {
		list, err := a.session.Role.List(
			management.Page(page),
		)
		if err != nil {
			return roles, err
		}

		roles = append(roles, list.Roles...)

		if !list.HasNext() {
			break
		}
		page++
	}

	return roles, err
}

func (a *auth0Struct) RoleShow(id string) (*management.Role, *management.UserList, error) {
	var err error
	var role *management.Role
	var users *management.UserList

	role, err = a.session.Role.Read(id)
	if err != nil {
		return role, users, err
	}

	users, err = a.session.Role.Users(role.GetID())
	if err != nil {
		return role, users, err
	}

	return role, users, nil
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

func (a *auth0Struct) UserCreate(name string, email string) error {
	var err error

	u := &management.User{
		Name:        auth0.String(name),
		Email:       auth0.String(email),
		Password:    auth0.String(vssutil.RandomString(32, "v1*:")),
		Connection:  auth0.String(a.connection),
		VerifyEmail: auth0.Bool(false),
	}

	err = a.session.User.Create(u)
	if err != nil {
		return err
	}

	err = a.PassReset(email)
	if err != nil {
		return err
	}

	return nil
}

func (a *auth0Struct) UserUpdate(id string, name string, email string, passreset bool) error {
	var err error

	u := &management.User{
		Name:  auth0.String(name),
		Email: auth0.String(email),
	}

	err = a.session.User.Update(id, u)
	if err != nil {
		return err
	}

	if passreset {
		err = a.PassReset(email)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *auth0Struct) UserBlock(id string) error {
	var err error

	user, err := a.session.User.Read(id)
	if err != nil {
		return err
	}

	userBlocked := user.GetBlocked()

	u := &management.User{
		Blocked: auth0.Bool(!userBlocked),
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

func (a *auth0Struct) PassReset(email string) error {
	var err error

	url := "https://" + a.domain + "/dbconnections/change_password"
	payload := strings.NewReader("{\"client_id\": \"" + a.client + "\",\"email\": \"" + email + "\",\"connection\": \"" + a.connection + "\"}")

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return err
	}
	req.Header.Add("content-type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	res.Body.Close()
	return nil
}
