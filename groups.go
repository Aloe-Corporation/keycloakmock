package keycloakmock

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getGroups(conf Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") == "" {
			c.JSON(http.StatusUnauthorized, "")
			return
		}

		c.JSON(http.StatusOK, flattenGroupConfig(conf.Groups))
	}
}

func flattenGroupConfig(groups []GroupConfig) []Group {
	var result []Group
	for _, group := range groups {
		subGroups := flattenGroupConfig(group.SubGroup)
		result = append(result, Group{
			ID:        stringP(group.UUID.String()),
			Name:      stringP(group.Name),
			SubGroups: &subGroups,
		})
		// result = append(result, flattenGroupConfig(group.SubGroup)...)
	}
	return result
}
