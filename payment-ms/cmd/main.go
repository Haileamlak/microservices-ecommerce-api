package main

import (
	"log"
	"net/http"
	http_handler "payment-ms/internal/handler/http"
	"payment-ms/internal/infrastructure"
	"payment-ms/internal/infrastructure/config"
	"payment-ms/internal/infrastructure/database"
	"payment-ms/internal/infrastructure/middleware"
	pb "payment-ms/internal/infrastructure/client/pb"
	"payment-ms/internal/repository"
	"payment-ms/internal/usecase"

	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/swaggo/http-swagger"
	_ "payment-ms/docs"
)


// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type: Bearer token. Example: "Bearer {token}"
func main() {
	cfg := config.LoadConfig()

	mongoClient, err := database.ConnectMongo(cfg.MongoURI)
	if err != nil {
		log.Fatalf("MongoDB connection failed: %v", err)
	}
	db := mongoClient.Database(cfg.MongoDatabase)

	// Repository → Usecase → Handler
	repo := repository.NewMongoPaymentRepository(db)
	paymentService := infrastructure.NewPaymentService(cfg.StripeSecretKey)

	conn, err := grpc.NewClient(cfg.OrderServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	orderClient := pb.NewOrderServiceClient(conn)

	uc := usecase.NewPaymentUseCase(repo, paymentService, orderClient)
	httpHandler := http_handler.NewPaymentHandler(uc)

	conn, err = grpc.NewClient(cfg.UserServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	userClient := pb.NewUserServiceClient(conn)

	router := chi.NewRouter()
	http_handler.SetupRouter(router, httpHandler, middleware.AuthMiddleware(userClient))
	router.Get("/swagger/*", httpSwagger.WrapHandler)


	log.Printf("HTTP Payment server listening on %s", cfg.HTTPPort)
	if err := http.ListenAndServe(cfg.HTTPPort, router); err != nil {
		log.Fatalf("Failed to serve HTTP: %v", err)
	}

}
