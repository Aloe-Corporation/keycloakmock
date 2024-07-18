package keycloakmock

import (
	"testing"

	"github.com/google/uuid"
	"gotest.tools/v3/assert"
)

func TestGroupContainsName(t *testing.T) {
	type testData struct {
		name        string
		group       GroupConfig
		groupName   string
		expectedRes bool
	}

	var testCases = [...]testData{
		{
			name:      "Success case",
			groupName: "tenant",
			group: GroupConfig{
				UUID: uuid.New(),
				Name: "tenant",
			},
			expectedRes: true,
		},
		{
			name:      "Success case: sub group",
			groupName: "test",
			group: GroupConfig{
				UUID: uuid.New(),
				Name: "tenant",
				SubGroups: []GroupConfig{
					{
						UUID: uuid.New(),
						Name: "test",
					},
				},
			},
			expectedRes: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := groupContainsName(testCase.group, testCase.groupName)
			assert.Equal(t, testCase.expectedRes, res)
		})
	}
}
