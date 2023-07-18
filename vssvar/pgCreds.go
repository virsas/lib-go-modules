package vssvar

import (
	"github.com/virsas/lib-go-modules/vsslib"
)

type PGCreds struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}

func OPPGCreds(op vsslib.OpHandler, opPGUser string) (PGCreds, error) {
	var err error
	var pgcreds PGCreds = PGCreds{}

	pgcreds.Host, err = op.Get(opPGUser, "hostname")
	if err != nil {
		return pgcreds, err
	}

	pgcreds.Port, err = op.Get(opPGUser, "port")
	if err != nil {
		return pgcreds, err
	}

	pgcreds.User, err = op.Get(opPGUser, "username")
	if err != nil {
		return pgcreds, err
	}

	pgcreds.Pass, err = op.Get(opPGUser, "password")
	if err != nil {
		return pgcreds, err
	}

	pgcreds.Name, err = op.Get(opPGUser, "database")
	if err != nil {
		return pgcreds, err
	}

	return pgcreds, nil
}
