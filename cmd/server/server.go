package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/99designs/gqlgen/graphql/handler"
// 	"github.com/99designs/gqlgen/graphql/playground"
// 	_ "github.com/lib/pq"
// 	"github.com/raulsilva-tech/UserEventsAPI/configs"
// 	"github.com/raulsilva-tech/UserEventsAPI/graph"
// 	"github.com/raulsilva-tech/UserEventsAPI/internal/infra/database"
// )

// func main() {

// 	//load configuration
// 	cfg, _ := configs.LoadConfig(".")

// 	//starting database connection
// 	DataSourceName := fmt.Sprintf("host=%s port=%s user=%s "+
// 		"password=%s dbname=%s sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
// 	fmt.Println(DataSourceName)

// 	db, err := sql.Open(cfg.DBDriver, DataSourceName)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()
// 	err = db.Ping()
// 	if err != nil {
// 		panic(err)
// 	}

// 	etDAO := database.NewEventTypeDAO(db)
// 	evDAO := database.NewEventDAO(db)
// 	userDAO := database.NewUserDAO(db)
// 	uaDAO := database.NewUserAddressDAO(db)

// 	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
// 		EventTypeDAO:   etDAO,
// 		EventDAO:       evDAO,
// 		UserDAO:        userDAO,
// 		UserAddressDAO: uaDAO,
// 	}}))

// 	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
// 	http.Handle("/query", srv)

// 	log.Printf("connect to http://localhost:%s/ for GraphQL playground", cfg.WebServerPort)
// 	log.Fatal(http.ListenAndServe(":"+cfg.WebServerPort, nil))
// }
