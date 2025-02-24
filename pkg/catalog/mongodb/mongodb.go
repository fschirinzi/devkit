package mongodb

import (
	"fmt"

	"github.com/adrianliechti/devkit/pkg/catalog"
	"github.com/adrianliechti/devkit/pkg/engine"
	"github.com/sethvargo/go-password/password"
)

var (
	_ catalog.Manager        = &Manager{}
	_ catalog.Decorator      = &Manager{}
	_ catalog.ShellProvider  = &Manager{}
	_ catalog.ClientProvider = &Manager{}
)

type Manager struct {
}

func (m *Manager) Name() string {
	return "mongodb"
}

func (m *Manager) Category() catalog.Category {
	return catalog.DatabaseCategory
}

func (m *Manager) DisplayName() string {
	return "MongoDB Database Server"
}

func (m *Manager) Description() string {
	return "MongoDB is a source-available cross-platform document-oriented database program."
}

const (
	DefaultShell = "/bin/bash"
)

func (m *Manager) New() (engine.Container, error) {
	image := "mongo:5-focal"

	username := "root"
	password := password.MustGenerate(10, 4, 0, false, false)

	return engine.Container{
		Image: image,

		Env: map[string]string{
			"MONGO_INITDB_ROOT_USERNAME": username,
			"MONGO_INITDB_ROOT_PASSWORD": password,
		},

		Ports: []*engine.ContainerPort{
			{
				Port:  27017,
				Proto: engine.ProtocolTCP,
			},
		},

		Mounts: []*engine.ContainerMount{
			{
				Path: "/data/db",
			},
		},
	}, nil
}

func (m *Manager) Info(instance engine.Container) (map[string]string, error) {
	username := instance.Env["MONGO_INITDB_ROOT_USERNAME"]
	password := instance.Env["MONGO_INITDB_ROOT_PASSWORD"]

	var uri string

	for _, p := range instance.Ports {
		if p.HostPort == nil || p.Port != 27017 {
			continue
		}

		uri = fmt.Sprintf("mongodb://%s:%s@localhost:%d", username, password, *p.HostPort)
	}

	return map[string]string{
		"Username": username,
		"Password": password,
		"URI":      uri,
	}, nil
}

func (m *Manager) Shell(instance engine.Container) (string, error) {
	return DefaultShell, nil
}

func (m *Manager) Client(instance engine.Container) (string, []string, error) {
	return "", []string{
		DefaultShell,
		"-c",
		"mongo --quiet --norc --username ${MONGO_INITDB_ROOT_USERNAME} --password ${MONGO_INITDB_ROOT_PASSWORD}",
	}, nil
}
