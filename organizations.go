package keycloakmock

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getOrganizationsMembers(conf Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") == "" {
			c.JSON(http.StatusUnauthorized, "auth header is empty")
			return
		}

		orgID := c.Param("org_id")
		if orgID != conf.OrganizationId.String() {
			c.JSON(http.StatusBadRequest, "invalid param org_id")
			return
		}

		type Response struct {
			Id        string `json:"id"`
			Firstname string `json:"firstName"`
		}

		c.JSON(http.StatusOK, []Response{
			{
				Id:        conf.UserUUID.String(),
				Firstname: "dummy firstname",
			},
		})
	}
}
