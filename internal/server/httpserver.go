package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/jordanmarcelino/terradiscover-backend/internal/config"
	"github.com/jordanmarcelino/terradiscover-backend/internal/controller"
	"github.com/jordanmarcelino/terradiscover-backend/internal/database"
	"github.com/jordanmarcelino/terradiscover-backend/internal/logger"
	"github.com/jordanmarcelino/terradiscover-backend/internal/middleware"
	"github.com/jordanmarcelino/terradiscover-backend/internal/repository"
	"github.com/jordanmarcelino/terradiscover-backend/internal/usecase"
	"github.com/jordanmarcelino/terradiscover-backend/internal/utils/encryptutils"
	"github.com/jordanmarcelino/terradiscover-backend/internal/utils/jwtutils"
	"github.com/jordanmarcelino/terradiscover-backend/internal/utils/validationutils"
)

type HttpServer struct {
	cfg    *config.Config
	server *http.Server
}

func NewHttpServer(cfg *config.Config) *HttpServer {
	gin.SetMode(cfg.App.Environment)

	router := gin.New()
	router.ContextWithFallback = true
	router.HandleMethodNotAllowed = true

	RegisterValidators()
	RegisterMiddleware(router, cfg)

	db := database.InitStdLib(cfg)
	jwtUtil := jwtutils.NewJwtUtil(cfg.Jwt)
	bcryptEncryptor := encryptutils.NewBcryptEncryptor(cfg.App.BCryptCost)

	dataStore := repository.NewDataStore(db)
	userRepository := repository.NewUserRepository(db)
	contactRepository := repository.NewContactRepository(db)

	authUseCase := usecase.NewAuthUseCase(jwtUtil, bcryptEncryptor, dataStore, userRepository)
	contactUseCase := usecase.NewContactUseCase(contactRepository)

	authController := controller.NewAuthController(authUseCase)
	contactController := controller.NewContactController(contactUseCase)

	auth := router.Group("/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
	}

	contacts := router.Group("/contacts", middleware.Authorization(jwtUtil))
	{
		contacts.GET("", contactController.Search)
		contacts.POST("", contactController.Create)
	}

	appController := controller.NewAppController()

	appController.Route(router)

	return &HttpServer{
		cfg: cfg,
		server: &http.Server{
			Addr:    fmt.Sprintf("%s:%d", cfg.HttpServer.Host, cfg.HttpServer.Port),
			Handler: router,
		},
	}
}

func (s *HttpServer) Start() {
	logger.Log.Info("Running HTTP server on port:", s.cfg.HttpServer.Port)
	if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Log.Fatal("Error while HTTP server listening:", err)
	}
	logger.Log.Info("HTTP server is not receiving new requests...")
}

func (s *HttpServer) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.cfg.HttpServer.GracePeriod)*time.Second)
	defer cancel()

	logger.Log.Info("Attempting to shut down the HTTP server...")
	if err := s.server.Shutdown(ctx); err != nil {
		logger.Log.Fatal("Error shutting down HTTP server:", err)
	}
	logger.Log.Info("HTTP server shut down gracefully")
}

func RegisterValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(validationutils.TagNameFormatter)

		v.RegisterValidation("password", validationutils.PasswordValidator)
		v.RegisterValidation("phone_number", validationutils.PhoneNumberValidator)
	}
}

func RegisterMiddleware(router *gin.Engine, cfg *config.Config) {
	middlewares := []gin.HandlerFunc{
		middleware.Logger(),
		middleware.ErrorHandler(),
		middleware.RequestTimeout(cfg),
		cors.New(cors.Config{
			AllowMethods:     []string{"*"},
			AllowHeaders:     []string{"*"},
			AllowOrigins:     []string{"http://localhost:3000", "http://localhost:5173"},
			AllowCredentials: true,
		}),
		gin.Recovery(),
	}

	router.Use(middlewares...)
}
