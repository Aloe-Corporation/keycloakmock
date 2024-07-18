package keycloakmock

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getGroups(conf Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") == "" {
			c.JSON(http.StatusUnauthorized, "")
			return
		}

		var groups []Group

		for _, groupName := range conf.Groups {
			parts := strings.Split(groupName, "/")
			group := Group{
				ID:        stringP(uuid.New().String()),
				Name:      &parts[0],
				SubGroups: &[]Group{},
			}

			for i := 1; i < len(parts); i++ {
				subGroup := Group{
					ID:   stringP(uuid.New().String()),
					Name: &parts[i],
				}

				*group.SubGroups = append(*group.SubGroups, subGroup)
			}

			groups = append(groups, group)
		}

		c.JSON(http.StatusOK, &groups)
	}
}
