package main

import (
	"database/sql"
	itemHandlers "items-service/internal/handlers/items"
	itemRepository "items-service/internal/repository/items"
	itemServices "items-service/internal/services/items"
	pb "items-service/proto"
	"net"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	//Logger
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	//DB
	connStr := "postgres://postgres:example@db:5432/market?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Fatal("Error while connecting database: ", zap.Error(err))
	}

	if err := db.Ping(); err != nil {
		logger.Fatal("Connection failed: ", zap.Error(err))
	}
	logger.Info("Connected to database")

	//Items
	repo := itemRepository.NewItemsRepository(db)
	services := itemServices.NewItemService(repo, logger)
	handlers := itemHandlers.NewItemsHandler(services)

	//gRPC server
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		logger.Fatal("failed to listen: ", zap.Error(err))
	}

	grpcServer := grpc.NewServer()
	pb.RegisterItemServiceServer(grpcServer, handlers)

	logger.Info("starting gRPC on 50051 port")

	if err := grpcServer.Serve(lis); err != nil {
		logger.Fatal("failed to serve: ", zap.Error(err))
	}
}
