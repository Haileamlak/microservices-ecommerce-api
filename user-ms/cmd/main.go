package main

import (
	grpc_handler "user-ms/internal/handler/grpc"
	http_handler "user-ms/internal/handler/http"
	"user-ms/internal/infrastructure/database"
	"user-ms/internal/infrastructure/config"
	"user-ms/internal/repository"
	"user-ms/internal/usecase"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"net"
	"google.golang.org/grpc"
	pb "user-ms/internal/infrastructure/client/pb"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "user-ms/docs"
)


func main() {
	cfg := config.LoadConfig()

	mongoClient, err := database.ConnectMongo(cfg.MongoURI)
	if err != nil {
		log.Fatalf("MongoDB connection failed: %v", err)
	}
	db := mongoClient.Database(cfg.MongoDatabase)

	repo := repository.NewMongoUserRepository(db)
	usecase := usecase.NewUserUsecase(repo)
	httpHandler := http_handler.NewUserHandler(usecase)
	grpcHandler := grpc_handler.NewUserHandler(usecase)

	
	lis, err := net.Listen("tcp", cfg.GRPCPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, grpcHandler)

	go func() {
		log.Printf("gRPC server listening on %s", cfg.GRPCPort)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/swagger/*", httpSwagger.WrapHandler)
	
	r.Post("/register", httpHandler.RegisterUser)
	r.Post("/login", httpHandler.LoginUser)
	

	port := cfg.HTTPPort
	log.Printf("Server is running on port %s", port)
	log.Fatal(http.ListenAndServe(port, r))
	
}