package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/raulsilva-tech/UserEventsAPI/configs"
	"github.com/raulsilva-tech/UserEventsAPI/internal/infra/database"
	"github.com/raulsilva-tech/UserEventsAPI/internal/infra/webserver/handler"
)

func main() {

	//load configuration
	cfg, _ := configs.LoadConfig(".")

	//starting database connection
	DataSourceName := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	fmt.Println(DataSourceName)

	db, err := sql.Open(cfg.DBDriver, DataSourceName)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	createRoutes(router, db)

	err = router.Run(":" + cfg.WebServerPort)
	if err != nil {
		log.Fatalln(err)
	}

}

func createRoutes(router *gin.Engine, db *sql.DB) {
	etDAO := database.NewEventTypeDAO(db)
	etHandler := handler.NewEventTypeHandler(etDAO)
	etGroup := router.Group("/event_types")
	etGroup.POST("/", etHandler.CreateEventType)
	etGroup.GET("/:id", etHandler.GetEventType)
	etGroup.GET("/", etHandler.GetAllEventType)
	etGroup.PUT("/:id", etHandler.UpdateEventType)
	etGroup.DELETE("/:id", etHandler.DeleteEventType)

	evDAO := database.NewEventDAO(db)
	evHandler := handler.NewEventHandler(evDAO)
	evGroup := router.Group("/events")
	evGroup.POST("/", evHandler.CreateEvent)
	evGroup.GET("/:id", evHandler.GetEvent)
	evGroup.GET("/", evHandler.GetAllEvent)
	evGroup.PUT("/:id", evHandler.UpdateEvent)
	evGroup.DELETE("/:id", evHandler.DeleteEvent)

	userDAO := database.NewUserDAO(db)
	userHandler := handler.NewUserHandler(userDAO)
	userGroup := router.Group("/users")
	userGroup.POST("/", userHandler.CreateUser)
	userGroup.GET("/:id", userHandler.GetUser)
	userGroup.GET("/", userHandler.GetAllUser)
	userGroup.PUT("/:id", userHandler.UpdateUser)
	userGroup.DELETE("/:id", userHandler.DeleteUser)

	uaDAO := database.NewUserAddressDAO(db)
	uaHandler := handler.NewUserAddressHandler(uaDAO)
	uaGroup := router.Group("/user_address")
	uaGroup.POST("/", uaHandler.CreateUserAddress)
	uaGroup.GET("/:id", uaHandler.GetUserAddress)
	uaGroup.GET("/", uaHandler.GetAllUserAddress)
	uaGroup.PUT("/:id", uaHandler.UpdateUserAddress)
	uaGroup.DELETE("/:id", uaHandler.DeleteUserAddress)

}
