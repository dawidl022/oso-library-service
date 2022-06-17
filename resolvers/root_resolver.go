package resolvers

import (
	"github.com/osohq/go-oso"
	"gorm.io/gorm"
)

type RootResolver struct {
	*BookMutation
}

func NewRootResolver(db *gorm.DB, oso *oso.Oso) *RootResolver {
	return &RootResolver{
		BookMutation: &BookMutation{
			db:  db,
			oso: oso,
		},
	}
}
