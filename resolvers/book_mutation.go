package resolvers

import (
	"net/http"

	"github.com/graph-gophers/graphql-go"
	"github.com/osohq/go-oso"
	"gorm.io/gorm"

	"github.com/dawidl022/oso-library-service/models"
)

type bookMutation struct {
	db  *gorm.DB
	oso *oso.Oso
}

func newBookMutation(db *gorm.DB, oso *oso.Oso) *bookMutation {
	return &bookMutation{
		db:  db,
		oso: oso,
	}
}

func (b *bookMutation) getUserAndBookModels(args bookMutationArgs) (*models.User, *models.Book) {
	user := getUser(args.UserId)
	book := getBook(args.BookId)
	return &user, &book
}

func (b *bookMutation) getAuthzHttpStatusCode(args bookMutationArgs, permission string) int32 {
	user, book := b.getUserAndBookModels(args)
	err := b.oso.Authorize(user, permission, book)
	if err != nil {
		return http.StatusForbidden
	}
	return http.StatusOK
}

type bookMutationArgs struct {
	UserId graphql.ID
	BookId graphql.ID
}

func (b *bookMutation) ReadBook(args bookMutationArgs) int32 {
	return b.getAuthzHttpStatusCode(args, "read")
}

func (b *bookMutation) CheckoutBook(args bookMutationArgs) int32 {
	return b.getAuthzHttpStatusCode(args, "checkout")
}

func (b *bookMutation) CheckinBook(args bookMutationArgs) int32 {
	return b.getAuthzHttpStatusCode(args, "checkin")
}

func (b *bookMutation) RemoveBook(args bookMutationArgs) int32 {
	return b.getAuthzHttpStatusCode(args, "remove")
}
