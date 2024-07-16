package keycloakmock

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func loginClient(conf Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientID := c.PostForm("client_id")
		if clientID == "" {
			c.JSON(http.StatusUnauthorized, "")
			return
		}

		if clientID != conf.ClientId {
			c.JSON(http.StatusUnauthorized, "")
			return
		}

		grantType := c.PostForm("grant_type")
		if grantType == "" {
			c.JSON(http.StatusUnauthorized, "")
			return
		}

		c.JSON(http.StatusOK, &JWT{
			AccessToken: "blablabla",
			IDToken:     "blablabla",
		})
	}
}
