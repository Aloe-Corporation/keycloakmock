package keycloakmock

type JWT struct {
	AccessToken      string `json:"access_token"`
	IDToken          string `json:"id_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	NotBeforePolicy  int    `json:"not-before-policy"`
	SessionState     string `json:"session_state"`
	Scope            string `json:"scope"`
}

// User represents the Keycloak User Structure
type User struct {
	ID                         *string                     `json:"id,omitempty"`
	CreatedTimestamp           *int64                      `json:"createdTimestamp,omitempty"`
	Username                   *string                     `json:"username,omitempty"`
	Enabled                    *bool                       `json:"enabled,omitempty"`
	Totp                       *bool                       `json:"totp,omitempty"`
	EmailVerified              *bool                       `json:"emailVerified,omitempty"`
	FirstName                  *string                     `json:"firstName,omitempty"`
	LastName                   *string                     `json:"lastName,omitempty"`
	Email                      *string                     `json:"email,omitempty"`
	FederationLink             *string                     `json:"federationLink,omitempty"`
	Attributes                 *map[string][]string        `json:"attributes,omitempty"`
	DisableableCredentialTypes *[]interface{}              `json:"disableableCredentialTypes,omitempty"`
	RequiredActions            *[]string                   `json:"requiredActions,omitempty"`
	Access                     *map[string]bool            `json:"access,omitempty"`
	ClientRoles                *map[string][]string        `json:"clientRoles,omitempty"`
	RealmRoles                 *[]string                   `json:"realmRoles,omitempty"`
	Groups                     *[]string                   `json:"groups,omitempty"`
	ServiceAccountClientID     *string                     `json:"serviceAccountClientId,omitempty"`
	Credentials                *[]CredentialRepresentation `json:"credentials,omitempty"`
}

type CredentialRepresentation struct {
	// Common part
	CreatedDate *int64  `json:"createdDate,omitempty"`
	Temporary   *bool   `json:"temporary,omitempty"`
	Type        *string `json:"type,omitempty"`
	Value       *string `json:"value,omitempty"`

	// <= v7
	Algorithm         *string             `json:"algorithm,omitempty"`
	Config            *MultiValuedHashMap `json:"config,omitempty"`
	Counter           *int32              `json:"counter,omitempty"`
	Device            *string             `json:"device,omitempty"`
	Digits            *int32              `json:"digits,omitempty"`
	HashIterations    *int32              `json:"hashIterations,omitempty"`
	HashedSaltedValue *string             `json:"hashedSaltedValue,omitempty"`
	Period            *int32              `json:"period,omitempty"`
	Salt              *string             `json:"salt,omitempty"`

	// >= v8
	CredentialData *string `json:"credentialData,omitempty"`
	ID             *string `json:"id,omitempty"`
	Priority       *int32  `json:"priority,omitempty"`
	SecretData     *string `json:"secretData,omitempty"`
	UserLabel      *string `json:"userLabel,omitempty"`
}

type MultiValuedHashMap struct {
	Empty      *bool    `json:"empty,omitempty"`
	LoadFactor *float32 `json:"loadFactor,omitempty"`
	Threshold  *int32   `json:"threshold,omitempty"`
}

// CompositesRepresentation represents the composite roles of a role
type CompositesRepresentation struct {
	Client *map[string][]string `json:"client,omitempty"`
	Realm  *[]string            `json:"realm,omitempty"`
}

// Role is a role
type Role struct {
	ID                 *string                   `json:"id,omitempty"`
	Name               *string                   `json:"name,omitempty"`
	ScopeParamRequired *bool                     `json:"scopeParamRequired,omitempty"`
	Composite          *bool                     `json:"composite,omitempty"`
	Composites         *CompositesRepresentation `json:"composites,omitempty"`
	ClientRole         *bool                     `json:"clientRole,omitempty"`
	ContainerID        *string                   `json:"containerId,omitempty"`
	Description        *string                   `json:"description,omitempty"`
	Attributes         *map[string][]string      `json:"attributes,omitempty"`
}

type Group struct {
	ID          *string              `json:"id,omitempty"`
	Name        *string              `json:"name,omitempty"`
	Path        *string              `json:"path,omitempty"`
	SubGroups   *[]Group             `json:"subGroups,omitempty"`
	Attributes  *map[string][]string `json:"attributes,omitempty"`
	Access      *map[string]bool     `json:"access,omitempty"`
	ClientRoles *map[string][]string `json:"clientRoles,omitempty"`
	RealmRoles  *[]string            `json:"realmRoles,omitempty"`
}
