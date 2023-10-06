package vsstest

import (
	"context"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var dbName string = "test"
var dbUser string = "test"
var dbPass string = "test"
var dbPort string = "5432/tcp"

func StartPostgres(ctx context.Context) (testcontainers.Container, string, string, string, string, string, error) {
	dbC, err := startPostgresContainer(ctx)
	if err != nil {
		return nil, "", "", "", "", "", err
	}
	port, err := getContainerPort(ctx, dbC)
	if err != nil {
		return nil, "", "", "", "", "", err
	}
	host, err := getContainerHost(ctx, dbC)
	if err != nil {
		return nil, "", "", "", "", "", err
	}

	return dbC, host, port, dbUser, dbPass, dbName, nil
}

func startPostgresContainer(ctx context.Context) (testcontainers.Container, error) {
	var err error

	req := testcontainers.ContainerRequest{
		Image: "postgres:12.10-alpine",
		Env: map[string]string{
			"POSTGRES_DB":       dbName,
			"POSTGRES_USER":     dbUser,
			"POSTGRES_PASSWORD": dbPass,
		},
		ExposedPorts: []string{dbPort},
		WaitingFor:   wait.ForListeningPort(nat.Port(dbPort)),
	}

	dbC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}

	return dbC, nil
}

func getContainerPort(ctx context.Context, dbC testcontainers.Container) (string, error) {
	var err error

	p, err := dbC.MappedPort(ctx, nat.Port(dbPort))
	if err != nil {
		return "", err
	}

	return p.Port(), nil
}

func getContainerHost(ctx context.Context, dbC testcontainers.Container) (string, error) {
	var err error

	host, err := dbC.Host(ctx)
	if err != nil {
		return "", err
	}

	return host, nil
}
