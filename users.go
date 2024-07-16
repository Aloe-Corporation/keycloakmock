package keycloakmock

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUserByID(conf Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") == "" {
			c.JSON(http.StatusUnauthorized, "")
			return
		}

		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusUnauthorized, "")
			return
		}

		if id != conf.UserUUID.String() {
			c.JSON(http.StatusBadRequest, "")
			return
		}

		c.JSON(http.StatusOK, &User{
			ID:        &id,
			Enabled:   stringB(true),
			FirstName: stringP("dummy firstname"),
			LastName:  stringP("dummy lasttname"),
			Email:     stringP("dummy@email.com"),
		})
	}
}

func updateUser(conf Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") == "" {
			c.JSON(http.StatusUnauthorized, "")
			return
		}

		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusUnauthorized, "")
			return
		}

		if id != conf.UserUUID.String() {
			c.JSON(http.StatusBadRequest, "")
			return
		}

		user := new(User)
		if err := c.ShouldBindJSON(user); err != nil {
			c.JSON(http.StatusBadRequest, "")
			return
		}

		c.JSON(http.StatusOK, "")
	}
}

func deleteUser(conf Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") == "" {
			c.JSON(http.StatusUnauthorized, "")
			return
		}

		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusUnauthorized, "")
			return
		}

		if id != conf.UserUUID.String() {
			c.JSON(http.StatusBadRequest, "")
			return
		}

		c.JSON(http.StatusOK, "")
	}
}
