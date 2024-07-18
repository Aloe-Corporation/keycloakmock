package keycloakmock

import (
	"net/http"
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
				UUID: uuid.New(),
			},
			{
				Name: "administrator",
				UUID: uuid.New(),
			},
			{
				Name: "manager",
				UUID: uuid.New(),
			},
			{
				Name: "operator",
				UUID: uuid.New(),
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
	router.Use(realmCheckMiddleware(c))

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

	router.Group("/admin/realms/:realm/roles").
		GET("", getRealmRoles(c)).
		GET("/:role_name", getRealmRoleByName(c))

	router.Group("/admin/realms/:realm/roles-by-id/:role_id").
		GET("", getRealmRoleById(c))

	server := httptest.NewServer(router)
	return server
}

func realmCheckMiddleware(conf Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		realm := c.Param("realm")
		if realm == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, "")
			return
		}

		if realm != conf.Realm || realm != "master" {
			c.AbortWithStatusJSON(http.StatusBadRequest, "")
			return
		}
		c.Next()
	}
}

func stringP(str string) *string {
	return &str
}

func stringB(b bool) *bool {
	return &b
}
