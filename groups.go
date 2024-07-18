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

		for _, g := range conf.Groups {
			match := searchGroupByName(searchGroupName, g)
			if match != nil {
				c.JSON(http.StatusOK, flattenGroupConfig([]GroupConfig{*match}))
				return
			}
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
		// result = append(result, flattenGroupConfig(group.SubGroup)...)
	}
	return result
}

func searchGroupByName(name string, group GroupConfig) *GroupConfig {
	if group.Name == name {
		return &group
	}

	if group.SubGroup == nil {
		return nil
	}

	for _, sub := range group.SubGroup {
		return searchGroupByName(name, sub)
	}

	return nil
}
