package keycloakmock

import (
	"net/http"
	"slices"

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
				ID:   stringP(role.UUID.String()),
				Name: stringP(role.Name),
			})
		}

		c.JSON(http.StatusOK, &roles)
	}
}

func getRealmRolesByUserID(conf Config) gin.HandlerFunc {
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

		var roles []Role
		for _, role := range conf.Roles {
			roles = append(roles, Role{
				ID:   stringP(role.UUID.String()),
				Name: stringP(role.Name),
			})
		}

		c.JSON(http.StatusOK, &roles)
	}
}

func addRealmRoleToUser(conf Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") == "" {
			c.JSON(http.StatusUnauthorized, "")
			return
		}

		roles := []Role{}
		if err := c.ShouldBindJSON(&roles); err != nil {
			c.JSON(http.StatusBadRequest, "can't unmarshal role")
			return
		}

		for _, role := range roles {
			if !slices.ContainsFunc(conf.Roles, func(roleConf struct {
				UUID uuid.UUID
				Name string
			}) bool {
				if roleConf.UUID.String() == *role.ID && roleConf.Name == *role.Name {
					return true
				}
				return false
			}) {
				c.JSON(http.StatusBadRequest, "unknown role")
				return
			}
		}

		c.JSON(http.StatusOK, "")
	}
}

func deleteRealmRoleFromUser(conf Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") == "" {
			c.JSON(http.StatusUnauthorized, "")
			return
		}

		roles := []Role{}
		if err := c.ShouldBindJSON(&roles); err != nil {
			c.JSON(http.StatusBadRequest, "can't unmarshal role")
			return
		}

		for _, role := range roles {
			if !slices.ContainsFunc(conf.Roles, func(roleConf struct {
				UUID uuid.UUID
				Name string
			}) bool {
				if roleConf.UUID.String() == *role.ID && roleConf.Name == *role.Name {
					return true
				}
				return false
			}) {
				c.JSON(http.StatusBadRequest, "unknown role")
				return
			}
		}

		c.JSON(http.StatusOK, "")
	}
}
