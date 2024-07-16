package keycloakmock

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getRealmRoles(conf Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") == "" {
			c.JSON(http.StatusUnauthorized, "")
			return
		}

		var roles []Role
		for _, role := range conf.Roles {
			roles = append(roles, Role{
				ID:   stringP(uuid.NewString()),
				Name: stringP(role),
			})
		}

		c.JSON(http.StatusOK, &roles)
	}
}
