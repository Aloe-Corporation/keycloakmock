package keycloakmock

import (
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Lauch(c Config) *httptest.Server {
	return launch(c)
}

func LauchDefault() (*httptest.Server, Config) {
	c := Config{
		Realm:    "test",
		UserUUID: uuid.New(),
		ClientId: "client",
		Roles: []RolesConfig{
			{
				Name: "default",
				ID:   uuid.New(),
			},
			{
				Name: "administrator",
				ID:   uuid.New(),
			},
			{
				Name: "manager",
				ID:   uuid.New(),
			},
			{
				Name: "operator",
				ID:   uuid.New(),
			},
		},
		Groups: []string{
			"test",
			"tenant/test",
			"tenant/test/subgroup",
		},
	}

	return launch(c), c
}

func launch(c Config) *httptest.Server {
	router := gin.New()
	router.POST("/realms/:realm/protocol/openid-connect/token", loginClient(c))

	router.POST("/admin/realms/:realm/users", createUser())
	router.Group("/admin/realms/:realm/users/:id").
		GET("", getUserByID(c)).
		PUT("", updateUser(c)).
		DELETE("", deleteUser(c)).
		GET("/role-mappings/realm", getRealmRolesByUserID(c)).
		POST("/role-mappings/realm", addRealmRoleToUser(c)).
		DELETE("/role-mappings/realm", deleteRealmRoleFromUser(c)).
		GET("/groups", getUserGroups(c))

	router.GET("/admin/realms/:realm/roles", getRealmRoles(c))

	server := httptest.NewServer(router)
	return server
}

func stringP(str string) *string {
	return &str
}

func stringB(b bool) *bool {
	return &b
}
