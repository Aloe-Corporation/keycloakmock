package keycloakmock

import "github.com/google/uuid"

type Config struct {
	Realm    string
	UserUUID uuid.UUID
	ClientId string
	Roles    []RolesConfig
	Groups   []string
}

type RolesConfig struct {
	ID   uuid.UUID
	Name string
}
