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

		groups := findGroups(conf.Groups, groupContainsName, searchGroupName)
		c.JSON(http.StatusOK, flattenGroupConfig(groups))
	}
}

func getGroupMembers(conf Config) gin.HandlerFunc {
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

		groups := findGroups(conf.Groups, groupContainsId, id)
		if len(groups) == 0 {
			c.JSON(http.StatusBadRequest, "")
			return
		}

		users := []*User{
			{
				ID:        stringP(conf.UserUUID.String()),
				Enabled:   stringB(true),
				FirstName: stringP("dummy firstname"),
				LastName:  stringP("dummy lasttname"),
				Email:     stringP("dummy@email.com"),
			},
		}

		c.JSON(http.StatusOK, &users)
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

type groupFinder func(g GroupConfig, attr string) bool

func findGroups(groups []GroupConfig, findFn groupFinder, attr string) []GroupConfig {
	var results []GroupConfig
	for _, group := range groups {
		if findFn(group, attr) {
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

func groupContainsId(group GroupConfig, id string) bool {
	if group.UUID.String() == id {
		return true
	}

	if group.SubGroups == nil {
		return false
	}

	for _, g := range group.SubGroups {
		return groupContainsName(g, id)
	}

	return false
}
