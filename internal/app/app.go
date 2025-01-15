package app

import (
	"TestTask/config"
	"TestTask/db/database"
	"TestTask/internal/cache"
	"TestTask/internal/handlers"
	"TestTask/internal/kafka"
	"TestTask/internal/repository"
	"TestTask/internal/routes"
	"TestTask/internal/service"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func Run() {
	config.LoadConfig("config.yaml")

	dbConfig := config.Config.Database
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Password, dbConfig.Name, dbConfig.SSLMode,
	)
	database.InitDB(connStr)

	log.Println("Database initialized")

	orderRepository := repository.NewOrderRepository(database.DB)
	productRepository := repository.NewProductRepository(database.DB)
	userRepository := repository.NewUserRepository(database.DB)
	logRepository := repository.NewLogRepository(database.DB)

	log.Println("Repositories initialized")

	kafkaConfig := config.Config.Kafka
	kafkaProducer := kafka.NewProducer(kafkaConfig.Brokers, kafkaConfig.Topic)

	log.Println("Kafka producer initialized")

	cacheService := cache.NewCacheService()
	eventService := service.NewEventService(kafkaProducer)
	orderService := service.NewOrderService(orderRepository, cacheService, eventService)
	productService := service.NewProductService(productRepository)
	userService := service.NewUserService(userRepository)
	authService := service.NewAuthService(userService)
	logService := service.NewLogService(logRepository)

	log.Println("Services initialized")

	orderHandler := handlers.NewOrderHandler(orderService, logService)
	productHandler := handlers.NewProductHandler(productService)
	authHandler := handlers.NewAuthHandlers(authService)

	log.Println("Handlers initialized")

	router := chi.NewRouter()
	apiRoutes := routes.NewRoutes(router)

	apiRoutes.SetupOrderRoutes(orderHandler)
	apiRoutes.SetupProductRoutes(productHandler)
	apiRoutes.SetupAuthRoutes(authHandler)
	apiRoutes.SetupSwagger()

	appConfig := config.Config.App
	log.Println("Server is running on port: ", appConfig.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", appConfig.Port), router)
	if err != nil {
		panic(err)
	}
}
