package main

import (
	"log"
	"net"
	"net/http"
	grpc_handler "product-ms/internal/handler/grpc"
	pb "product-ms/internal/infrastructure/client/pb"
	http_handler "product-ms/internal/handler/http"
	"product-ms/internal/infrastructure/config"
	"product-ms/internal/infrastructure/database"
	"product-ms/internal/infrastructure/middleware"
	"product-ms/internal/repository"
	"product-ms/internal/usecase"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	_ "product-ms/docs"
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
	repo := repository.NewMongoProductRepository(db)
	uc := usecase.NewProductUsecase(repo)
	httpHandler := http_handler.NewProductHandler(uc)

	conn, err := grpc.NewClient(cfg.UserServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	userClient := pb.NewUserServiceClient(conn)

	router := chi.NewRouter()
	http_handler.SetupRouter(router, httpHandler, middleware.AuthMiddleware(userClient))
	router.Get("/swagger/*", httpSwagger.WrapHandler)

	go func() {
		log.Printf("HTTP server listening on %s", cfg.HTTPPort)
		if err := http.ListenAndServe(cfg.HTTPPort, router); err != nil {
			log.Fatalf("Failed to serve HTTP: %v", err)
		}
	}()

	grpcHandler := grpc_handler.NewProductHandler(uc)

	grpcListener, err := net.Listen("tcp", cfg.GRPCPort)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", cfg.GRPCPort, err)
	}

	s := grpc.NewServer()
	pb.RegisterProductServiceServer(s, grpcHandler)

	log.Printf("Product service listening on %s", cfg.GRPCPort)
	if err := s.Serve(grpcListener); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
