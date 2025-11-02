package main

import (
	"items-service/config"
	itemHandlers "items-service/internal/handlers/items"
	itemRepository "items-service/internal/repository/items"
	itemServices "items-service/internal/services/items"
	"items-service/pkg/postgres"
	"net"

	pb "github.com/marketplace-go/contracts/items"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	//Logger
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	//Config
	cfg := config.Get()

	//DB
	connStr := cfg.DBurl
	db := postgres.New(logger, connStr)

	if err := db.Ping(); err != nil {
		logger.Fatal("Connection failed: ", zap.Error(err))
	}
	logger.Info("Connected to database")

	//Items
	repo := itemRepository.NewItemsRepository(db)
	services := itemServices.NewItemService(repo, logger)
	handlers := itemHandlers.NewItemsHandler(services)

	//gRPC server
	lis, err := net.Listen("tcp", ":"+cfg.Port)

	if err != nil {
		logger.Fatal("failed to listen: ", zap.Error(err))
	}

	grpcServer := grpc.NewServer()
	pb.RegisterItemServiceServer(grpcServer, handlers)

	logger.Info("starting gRPC on ", zap.String("port", cfg.Port))

	if err := grpcServer.Serve(lis); err != nil {
		logger.Fatal("failed to serve: ", zap.Error(err))
	}
}
