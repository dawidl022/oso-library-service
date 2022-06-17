package resolvers

import (
	"log"
	"strconv"

	"github.com/graph-gophers/graphql-go"
	"github.com/osohq/go-oso"
	"gorm.io/gorm"

	"github.com/dawidl022/oso-library-service/models"
)

type Book struct {
	b *models.Book
}

func (b *Book) Title() string {
	return b.b.Title
}

func (b *Book) GloballyAvailable() bool {
	return b.b.GloballyAvailable
}

func (b *Book) Regions() []string {
	return b.b.Regions
}

type bookResolver struct {
	*bookQuery
	*bookMutation
}

func newBookResolver(db *gorm.DB, oso *oso.Oso) *bookResolver {
	return &bookResolver{
		newBookQuery(db, oso),
		newBookMutation(db, oso),
	}
}

func getBook(bookId graphql.ID) models.Book {
	id, err := strconv.Atoi(string(bookId))
	if err != nil {
		log.Fatal(err)
	}
	return *getAllBooks()[id-1]
}

func getAllBooks() []*models.Book {
	return []*models.Book{
		{
			Title:   "Refactoring",
			Regions: []string{"London", "Amsterdam"},
		},
		{
			Title:             "Structure and Interpretation of Computer Programs",
			GloballyAvailable: true,
		},
		{
			Title:   "Design Patterns: Elements of Reusable Object-Oriented Software",
			Regions: []string{"Paris"},
		},
	}
}
