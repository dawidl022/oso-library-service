package resolvers

import (
	"github.com/dawidl022/oso-library-service/models"
	"github.com/graph-gophers/graphql-go"
	"github.com/osohq/go-oso"
	"gorm.io/gorm"
)

type bookQuery struct {
	db  *gorm.DB
	oso *oso.Oso
}

func newBookQuery(db *gorm.DB, oso *oso.Oso) *bookQuery {
	return &bookQuery{
		db:  db,
		oso: oso,
	}
}

type availableBooksQuery struct {
	UserId graphql.ID
}

func (b *bookQuery) AvailableBooks(args availableBooksQuery) []*Book {
	authorizedBooks := make([]*models.Book, 0)
	user := getUser(args.UserId)

	for _, book := range getAllBooks() {
		if b.oso.Authorize(user, "read", book) == nil {
			authorizedBooks = append(authorizedBooks, book)
		}
	}

	res := make([]*Book, 0, len(authorizedBooks))
	for _, book := range authorizedBooks {
		res = append(res, &Book{b: book})
	}
	return res
}
