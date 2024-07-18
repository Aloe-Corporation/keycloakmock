package keycloakmock

import "github.com/google/uuid"

type Config struct {
	Realm    string
	UserUUID uuid.UUID
	ClientId string
	Roles    []RolesConfig
	Groups   []GroupConfig
}

type RolesConfig struct {
	UUID uuid.UUID
	Name string
}

type GroupConfig struct {
	UUID      uuid.UUID
	Name      string
	SubGroups []GroupConfig
}
