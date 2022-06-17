package resolvers

import (
	"github.com/graph-gophers/graphql-go"

	"github.com/dawidl022/oso-library-service/models"
)

func getUser(userId graphql.ID) models.User {
	return models.User{
		Role:    "member",
		Regions: []string{"Madrid", "London"},
	}
}
