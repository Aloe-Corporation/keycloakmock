package keycloakmock

import "github.com/google/uuid"

type Config struct {
	Realm    string
	UserUUID uuid.UUID
	ClientId string
	Roles    []struct {
		UUID uuid.UUID
		Name string
	}
	Groups []string
}
