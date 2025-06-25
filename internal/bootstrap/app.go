package bootstrap

func Run() error {
	db := SetupDatabase()

	repos := SetupRepositories(db)
	usecases := SetupUsecases(repos)
	controllers := SetupControllers(usecases)
	router := SetupRouter(controllers)

	return router.Run(":8080")
}
