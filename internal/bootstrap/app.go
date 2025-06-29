package bootstrap

import (
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func BuildApp(db *mongo.Database) *gin.Engine {
	repos := SetupRepositories(db)
	usecases := SetupUsecases(repos)
	controllers := SetupControllers(usecases)
	return SetupRouter(controllers)
}

func Run() error {
	db := SetupDatabase()
	router := BuildApp(db)
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "80" // Default port if not set
	}
	return router.Run(`:` + port)
}
