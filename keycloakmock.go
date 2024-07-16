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
	}

	return launch(c), c
}

func launch(c Config) *httptest.Server {
	router := gin.New()
	router.POST("/realms/:realm/protocol/openid-connect/token", loginClient(c))

	router.Group("/admin/realms/:realm/users/:id").
		GET("", getUserByID(c)).
		PUT("", updateUser(c)).
		DELETE("", deleteUser(c))

	server := httptest.NewServer(router)
	return server
}

func stringP(str string) *string {
	return &str
}

func stringB(b bool) *bool {
	return &b
}
