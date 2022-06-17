package resolvers

import (
	"net/http"

	"github.com/dawidl022/oso-library-service/models"
	"github.com/graph-gophers/graphql-go"
	"github.com/osohq/go-oso"
	"gorm.io/gorm"
)

type BookMutation struct {
	db  *gorm.DB
	oso *oso.Oso
}

func (b *BookMutation) getUser(userId graphql.ID) models.User {
	return models.User{
		Role: "reader",
	}
}

func (b *BookMutation) getBook(userId graphql.ID) models.Book {
	return models.Book{}
}

type readBookArgs struct {
	ReaderId graphql.ID
	BookId   graphql.ID
}

func (b *BookMutation) ReadBook(args readBookArgs) int32 {
	user := b.getUser(args.ReaderId)
	book := b.getBook(args.BookId)

	err := b.oso.Authorize(user, "read", book)
	if err != nil {
		return http.StatusForbidden
	}
	return http.StatusOK
}
