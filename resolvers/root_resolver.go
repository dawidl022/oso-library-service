package resolvers

import (
	"github.com/osohq/go-oso"
	"gorm.io/gorm"
)

type rootResolver struct {
	*bookResolver
}

func NewRootResolver(db *gorm.DB, oso *oso.Oso) *rootResolver {
	return &rootResolver{
		newBookResolver(db, oso),
	}
}
