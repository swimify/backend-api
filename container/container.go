package container

import (
	"net/http"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/swimlibre/backend-api/routes"

	"github.com/gorilla/mux"
)

type Container struct {
	DB     *gorm.DB
	Router *mux.Router
}

func BuildContainer() Container {
	var c Container
	c.DB = BuildDBService()
	c.Router = BuildRouterService()

	return c
}

func BuildDBService() *gorm.DB {
	var db *gorm.DB
	var err error

	switch os.Getenv("APP_DB_DRIVER") {
	case "sqlite3":
		db, err = gorm.Open(sqlite.Open("swimlibre.db"), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})
		if err != nil {
			panic("failed to connect SQLite3 database")
		}
	case "postgres":
		panic("@todo - Postgres not configured yet")
	}

	return db
}

func BuildRouterService() *mux.Router {
	//create a new router
	router := mux.NewRouter()
	routes.AttachRoutes(router)
	http.Handle("/", router)

	return router
}
