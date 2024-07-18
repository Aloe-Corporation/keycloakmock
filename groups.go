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
		searchGroupName, ok := c.GetQuery("search")
		if !ok {
			c.JSON(http.StatusOK, flattenGroupConfig(conf.Groups))
			return
		}

		match := findGroupByName(conf.Groups, searchGroupName)
		if match != nil {
			c.JSON(http.StatusOK, flattenGroupConfig([]GroupConfig{*match}))
		}

		c.JSON(http.StatusNoContent, "")
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
	}

	return result
}

func findGroupByName(groups []GroupConfig, name string) *GroupConfig {
	for _, group := range groups {
		if group.Name == name {
			return &group
		}
		if found := findGroupByName(group.SubGroup, name); found != nil {
			return found
		}
	}
	return nil
}
