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

func OPPGCreds(op vsslib.OpHandler, opPGItem string) (PGCreds, error) {
	var err error
	var pgcreds PGCreds = PGCreds{}

	pgcreds.Host, err = op.Get(opPGItem, "hostname")
	if err != nil {
		return pgcreds, err
	}

	pgcreds.Port, err = op.Get(opPGItem, "port")
	if err != nil {
		return pgcreds, err
	}

	pgcreds.User, err = op.Get(opPGItem, "username")
	if err != nil {
		return pgcreds, err
	}

	pgcreds.Pass, err = op.Get(opPGItem, "password")
	if err != nil {
		return pgcreds, err
	}

	pgcreds.Name, err = op.Get(opPGItem, "database")
	if err != nil {
		return pgcreds, err
	}

	return pgcreds, nil
}
