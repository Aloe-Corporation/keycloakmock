package keycloakmock

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
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

func getRealmRoleById(conf Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") == "" {
			c.JSON(http.StatusUnauthorized, "")
			return
		}

		roleID := c.Param("role_id")
		if roleID == "" {
			c.JSON(http.StatusBadRequest, "")
			return
		}

		roleIndex := slices.IndexFunc(conf.Roles, func(role RolesConfig) bool {
			return role.UUID.String() == roleID
		})

		if roleIndex == -1 {
			c.JSON(http.StatusBadRequest, "")
			return
		}

		c.JSON(http.StatusOK, &Role{
			ID:   stringP(conf.Roles[roleIndex].UUID.String()),
			Name: stringP(conf.Roles[roleIndex].Name),
		})
	}
}

func getRealmRoleByName(conf Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") == "" {
			c.JSON(http.StatusUnauthorized, "")
			return
		}

		roleName := c.Param("role_name")
		if roleName == "" {
			c.JSON(http.StatusBadRequest, "")
			return
		}

		roleIndex := slices.IndexFunc(conf.Roles, func(role RolesConfig) bool {
			return role.Name == roleName
		})

		if roleIndex == -1 {
			c.JSON(http.StatusBadRequest, "")
			return
		}

		c.JSON(http.StatusOK, &Role{
			ID:   stringP(conf.Roles[roleIndex].UUID.String()),
			Name: stringP(conf.Roles[roleIndex].Name),
		})
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

		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusUnauthorized, "")
			return
		}

		if id != conf.UserUUID.String() {
			c.JSON(http.StatusBadRequest, "")
			return
		}

		roles := []Role{}
		if err := c.ShouldBindJSON(&roles); err != nil {
			c.JSON(http.StatusBadRequest, "can't unmarshal role")
			return
		}

		for _, role := range roles {
			if !slices.ContainsFunc(conf.Roles, func(roleConf RolesConfig) bool {
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

		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusUnauthorized, "")
			return
		}

		if id != conf.UserUUID.String() {
			c.JSON(http.StatusBadRequest, "")
			return
		}

		roles := []Role{}
		if err := c.ShouldBindJSON(&roles); err != nil {
			c.JSON(http.StatusBadRequest, "can't unmarshal role")
			return
		}

		for _, role := range roles {
			if !slices.ContainsFunc(conf.Roles, func(roleConf RolesConfig) bool {
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
