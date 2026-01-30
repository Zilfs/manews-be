package app

import (
	"context"
	"log"
	"manews/config"
	"manews/internal/adapter/cloudflare"
	"manews/internal/adapter/handler"
	"manews/internal/adapter/repository"
	"manews/internal/core/service"
	"manews/lib/auth"
	"manews/lib/middleware"
	"manews/lib/pagination"
	"os/signal"
	"syscall"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"os"
)

func RunServer() {
	cfg := config.NewConfig()
	db, err := cfg.ConnectionPostgres()
	if err != nil {
		log.Fatal("Failed to connect to database: %v", err)
		return
	}

	err = os.MkdirAll("./temp/content", 0755)
	if err != nil {
		log.Fatal("Error creating temp directory: %v", err)
		return
	}

	// Cloudflare R2
	cdfR2 := cfg.LoadAwsConfig()
	s3Client := s3.NewFromConfig(cdfR2)
	r2Adapter := cloudflare.NewCloudflareR2Adapter(s3Client, cfg)

	jwt := auth.NewJwt(cfg)
	middlewareAuth := middleware.NewMiddleware(cfg)

	_ = pagination.NewPagination()

	// Repository
	authRepo := repository.NewAuthRepository(db.DB)
	categoryRepo := repository.NewCategoryRepository(db.DB)
	contentRepo := repository.NewContentRepository(db.DB)

	// Service
	authService := service.NewAuthService(authRepo, cfg, jwt)
	categoryService := service.NewCategoryService(categoryRepo)
	contentService := service.NewContentService(contentRepo, cfg, r2Adapter)

	// Handler
	authHandler := handler.NewAuthHandler(authService)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	contentHandler := handler.NewContentHandler(contentService)

	app := fiber.New()
	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] %{ip} %{status} - %{latency} ${method} ${path}\n",
	}))

	api := app.Group("/api")
	api.Post("/login", authHandler.Login)

	adminApp := api.Group("/admin")
	adminApp.Use(middlewareAuth.CheckToken())

	// Caetgory
	categoryApp := adminApp.Group("/categories")
	categoryApp.Get("/", categoryHandler.GetCategories)
	categoryApp.Post("/", categoryHandler.CreateCategory)
	categoryApp.Put("/:categoryId", categoryHandler.EditCategory)
	categoryApp.Get("/:categoryId", categoryHandler.GetCategoryByID)
	categoryApp.Delete("/:categoryId", categoryHandler.DeleteCategory)

	// Content
	contentApp := adminApp.Group("/contents")
	contentApp.Get("/", contentHandler.GetContents)
	contentApp.Post("/", contentHandler.CreateContent)
	contentApp.Put("/:contentId", contentHandler.EditContent)
	contentApp.Get("/:contentId", contentHandler.GetContentByID)
	contentApp.Delete("/:contentId", contentHandler.DeleteContent)
	contentApp.Post("/upload-image", contentHandler.UploadImageR2)

	go func() {
		if cfg.App.AppPort == "" {
			cfg.App.AppPort = os.Getenv("APP_PORT")
		}

		err := app.Listen(":" + cfg.App.AppPort)
		if err != nil {
			log.Fatal("Failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGTERM)

	<-quit

	log.Println("Server shutdown of 5 seconds")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	app.ShutdownWithContext(ctx)
}
