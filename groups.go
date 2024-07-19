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

		groups := findGroupsByName(conf.Groups, searchGroupName)
		c.JSON(http.StatusOK, flattenGroupConfig(groups))
	}
}

func flattenGroupConfig(groups []GroupConfig) []Group {
	var result []Group
	for _, group := range groups {
		subGroups := flattenGroupConfig(group.SubGroups)
		result = append(result, Group{
			ID:        stringP(group.UUID.String()),
			Name:      stringP(group.Name),
			SubGroups: &subGroups,
		})
	}

	return result
}

func findGroupsByName(groups []GroupConfig, name string) []GroupConfig {
	var results []GroupConfig
	for _, group := range groups {
		if groupContainsName(group, name) {
			results = append(results, group)
		}
	}

	return results
}

func groupContainsName(group GroupConfig, name string) bool {
	if group.Name == name {
		return true
	}

	if group.SubGroups == nil {
		return false
	}

	for _, g := range group.SubGroups {
		return groupContainsName(g, name)
	}

	return false
}
