package bootstrap

import (
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
	return router.Run(":8080")
}
