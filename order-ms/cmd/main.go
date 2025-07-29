package main

import (
	"log"
	"net"
	"net/http"
	grpc_handler "order-ms/internal/handler/grpc"
	http_handler "order-ms/internal/handler/http"
	"order-ms/internal/infrastructure/config"
	"order-ms/internal/infrastructure/database"
	"order-ms/internal/infrastructure/middleware"
	pb "order-ms/internal/infrastructure/client/pb"
	"order-ms/internal/repository"
	"order-ms/internal/usecase"

	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "order-ms/docs"
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

	conn, err := grpc.NewClient(cfg.ProductServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	productServiceClient := pb.NewProductServiceClient(conn)

	// Repository → Usecase → Handler
	repo := repository.NewMongoOrderRepository(db)
	uc := usecase.NewOrderUsecase(repo, productServiceClient)
	httpHandler := http_handler.NewOrderHandler(uc)
	grpcHandler := grpc_handler.NewOrderHandler(uc)

	conn, err = grpc.NewClient(cfg.UserServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	userClient := pb.NewUserServiceClient(conn)

	router := chi.NewRouter()
	http_handler.SetupRouter(router, httpHandler, middleware.AuthMiddleware(userClient))
	router.Get("/swagger/*", httpSwagger.WrapHandler)
	

	lis, err := net.Listen("tcp", cfg.GRPCPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, grpcHandler)

	go func() {
		log.Printf("gRPC server listening on %s", cfg.GRPCPort)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	log.Printf("HTTP server listening on %s", cfg.HTTPPort)
	if err := http.ListenAndServe(cfg.HTTPPort, router); err != nil {
		log.Fatalf("Failed to serve HTTP: %v", err)
	}

}
