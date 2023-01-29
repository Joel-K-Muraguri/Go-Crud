package api

import (
	"log"

	"github.com/Joel-K-Muraguri/Go-Crud/api/controller"
	"github.com/Joel-K-Muraguri/Go-Crud/api/seed"
	"github.com/joho/godotenv"
)

var server = controller.Server{}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}

func Run() {

	// var err error
	// err = godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Error getting env, %v", err)
	// } else {
	// 	fmt.Println("We are getting the env values")
	// }

	// username := os.Getenv("DB_USER")
	// password := os.Getenv("DB_PASSWORD")
	// dbName := os.Getenv("DB_NAME")
	// dbHost := os.Getenv("DB_HOST")
	// dbPort := os.Getenv("DB_PORT")

	server.Initialize()

	seed.Load(server.DB)

	server.Run(":8080")

}
