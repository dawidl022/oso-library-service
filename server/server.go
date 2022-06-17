package server

import (
	"log"
	"net/http"
	"os"
	"reflect"
	"time"

	"github.com/dawidl022/oso-library-service/models"
	"github.com/dawidl022/oso-library-service/resolvers"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/osohq/go-oso"
)

type server struct {
	router chi.Router
}

func StartServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	s := newServer()
	s.setup()
	log.Fatal(http.ListenAndServe(":"+port, s.router))
}

func newServer() server {
	s := server{
		router: chi.NewRouter(),
	}

	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.RealIP)
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)

	s.router.Use(middleware.Timeout(60 * time.Second))

	return s
}

func (s *server) setup() {
	b, err := os.ReadFile("server/graphql/query.graphql")
	if err != nil {
		log.Fatal("Cannot read grapql schema files:", err)
	}

	// db, err := initDB(conf)
	// if err != nil {
	// 	log.Fatal("Cannot initialise database:", err)
	// }

	oso, err := NewOso()
	if err != nil {
		log.Fatal(err)
	}

	schema := graphql.MustParseSchema(string(b), resolvers.NewRootResolver(nil, oso))
	s.routes(schema)
}

func NewOso() (*oso.Oso, error) {
	o, err := oso.NewOso()
	if err != nil {
		return nil, err
	}

	o.RegisterClass(reflect.TypeOf(models.User{}), nil)
	o.RegisterClass(reflect.TypeOf(models.Book{}), nil)

	err = o.LoadFiles([]string{"auth/main.polar"})
	if err != nil {
		return nil, err
	}

	return &o, nil
}

func (s *server) routes(schema *graphql.Schema) {
	s.router.Method(http.MethodPost, "/graphql", &relay.Handler{Schema: schema})
}
