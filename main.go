package main

import (
	"os"

	"backend/app/article/delivery/http"
	"backend/app/article/repository/postgres"
	"backend/app/article/repository/redis"
	"backend/app/article/usecase"

	user_http "backend/app/user/delivery/http"
	user_repo_posgres "backend/app/user/repository/postgres"
	user_usecase_posgres "backend/app/user/usecase"

	"backend/pkg/configs"
	"backend/pkg/middleware"

	"backend/pkg/routes"
	"backend/pkg/utils"

	"backend/platform/cache"
	"backend/platform/database"

	// "github.com/gofiber/fiber/v2"
	fiber "github.com/gofiber/fiber/v2"

	_ "backend/docs"

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

// @title API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email khalidalhabiie07@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath
func main() {

	// set platform ex: database
	redisClient := cache.RedisConnection()
	postgresDB := database.PostgreSQLConnection()

	// repository
	articlePostgres := postgres.New(postgresDB)
	articleCache := redis.New(redisClient)
	userPostgres := user_repo_posgres.New(postgresDB)

	// usecase
	articleUsecase := usecase.New(articlePostgres, articleCache)
	userUsecase := user_usecase_posgres.New(userPostgres)

	// handlers
	articleHandler := http.New(articleUsecase)
	userHandler := user_http.New(userUsecase)

	// Define Fiber config
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Register Routes
	userHandler.Register(app)
	articleHandler.Register(app) // articles handler
	routes.SwaggerRoute(app)     // Register a route for API Docs (Swagger).
	routes.PromotheusrRoute(app) // promotheus handler
	routes.NotFoundRoute(app)    // Register route for 404 Error.

	// Start server (with or without graceful shutdown).
	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
