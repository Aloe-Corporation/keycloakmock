package keycloakmock

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func createUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") == "" {
			c.JSON(http.StatusUnauthorized, "")
			return
		}

		user := new(User)
		if err := c.ShouldBindJSON(user); err != nil {
			c.JSON(http.StatusBadRequest, "")
			return
		}

		if *user.Email == "" {
			c.JSON(http.StatusBadRequest, "")
			return
		}

		c.Writer.Header().Add("Location", "/"+uuid.NewString())

		c.JSON(http.StatusOK, "")
	}
}

func getUserGroups(conf Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") == "" {
			c.JSON(http.StatusUnauthorized, "")
			return
		}

		c.JSON(http.StatusOK, flattenGroupConfig(conf.Groups))
	}
}

func AddUserToGroup(conf Config) gin.HandlerFunc {
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

		groupID := c.Param("group_id")
		if groupID == "" {
			c.JSON(http.StatusUnauthorized, "")
			return
		}

		groups := findGroups(conf.Groups, groupContainsId, groupID)
		if len(groups) == 0 {
			c.JSON(http.StatusBadRequest, "")
			return
		}

		c.JSON(http.StatusNoContent, "")
	}
}
