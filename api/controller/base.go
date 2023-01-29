package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Joel-K-Muraguri/Go-Crud/api/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	// "gorm.io/driver/postgres"
)

type Server struct {
	Router *mux.Router
	DB     *gorm.DB
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "Games"
)

func (server *Server) Initialize() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	// username := os.Getenv("DB_USER")
	// password := os.Getenv("DB_PASSWORD")
	// dbName := os.Getenv("DB_NAME")
	// host := os.Getenv("DB_HOST")
	// port := os.Getenv("DB_PORT")

	DBURL := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// dburi := fmt.Sprintf("host=%s port=%s dbname=%s user='%s' password=%s sslmode=disable", host, port, dbName, username, password)
	// fmt.Println(dburi)

	server.DB, err = gorm.Open("postgres", DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", "postgres")
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", "postgres")
	}

	server.DB.Debug().AutoMigrate(&models.Game{})

	server.Router = mux.NewRouter()

	server.initializeRoutes()

}

func (server *Server) Run(port string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(port, server.Router))
}
