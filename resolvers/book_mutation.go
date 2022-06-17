package resolvers

import (
	"net/http"

	"github.com/graph-gophers/graphql-go"
	"github.com/osohq/go-oso"
	"gorm.io/gorm"

	"github.com/dawidl022/oso-library-service/models"
)

type BookMutation struct {
	db  *gorm.DB
	oso *oso.Oso
}

func (b *BookMutation) getUser(userId graphql.ID) models.User {
	return models.User{
		Role:    "member",
		Regions: []string{"Madrid", "London"},
	}
}

func (b *BookMutation) getBook(userId graphql.ID) models.Book {
	return models.Book{
		GloballyAvailable: false,
		Regions:           []string{"London", "Amsterdam"},
	}
}

func (b *BookMutation) getUserAndBookModels(args bookMutationArgs) (*models.User, *models.Book) {
	user := b.getUser(args.UserId)
	book := b.getBook(args.BookId)
	return &user, &book
}

func (b *BookMutation) getAuthzHttpStatusCode(args bookMutationArgs, permission string) int32 {
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

func (b *BookMutation) ReadBook(args bookMutationArgs) int32 {
	return b.getAuthzHttpStatusCode(args, "read")
}

func (b *BookMutation) CheckoutBook(args bookMutationArgs) int32 {
	return b.getAuthzHttpStatusCode(args, "checkout")
}

func (b *BookMutation) CheckinBook(args bookMutationArgs) int32 {
	return b.getAuthzHttpStatusCode(args, "checkin")
}

func (b *BookMutation) RemoveBook(args bookMutationArgs) int32 {
	return b.getAuthzHttpStatusCode(args, "remove")
}
