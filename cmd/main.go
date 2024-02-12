package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/patrickchagastavares/rinha-backend-2024/internal/controllers"
	"github.com/patrickchagastavares/rinha-backend-2024/internal/handlers"
	"github.com/patrickchagastavares/rinha-backend-2024/internal/repositories"
	"github.com/patrickchagastavares/rinha-backend-2024/internal/services"
	"github.com/patrickchagastavares/rinha-backend-2024/pkg/httpRouter"
	"github.com/patrickchagastavares/rinha-backend-2024/pkg/logger"
)

func main() {
	godotenv.Load(".env")
	db := os.Getenv("DATABASE_URL")

	var (
		log          = logger.NewLogrusLogger()
		repositories = repositories.New(repositories.Options{
			DB_URL: db,
			Log:    log,
		})
		services = services.New(services.Options{
			Repo: repositories,
			Log:  log,
		})
		controllers = controllers.New(controllers.Options{
			Srv: services,
			Log: log,
		})
		router = httpRouter.NewEchoRouter()
	)

	handlers.NewRouter(handlers.Options{
		Router: router,
		Ctrl:   controllers,
	})

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = ":8000"
	}
	log.Info("start serve in port:", port)
	if err := router.Server(port); err != nil {
		log.Fatal(err)
	}

}
