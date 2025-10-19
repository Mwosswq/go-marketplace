package main

import (
	"context"
	"database/sql"
	"log"
	itemHandlers "main/internal/handlers/items"
	itemRepository "main/internal/repository/items"
	itemServices "main/internal/services/items"
	responder "main/pkg"
	"net/http"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func main() {
	//DB
	connStr := "postgres://postgres:example@db:5432/market?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln("Error while connecting database: %w", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Connection failed %w", err)
	}
	log.Println("Connected to database")

	//Logger
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	//utils
	r := responder.New(logger)

	//Items
	itemsRepo := itemRepository.NewItemsRepository(db, logger)
	itemsServices := itemServices.NewItemService(itemsRepo, logger)
	itemsHandlers := itemHandlers.NewItemsHandler(itemsServices, logger, r)

	//mux, routes and server

	mux := http.NewServeMux()

	//items endpoints
	mux.HandleFunc("/items", itemsHandlers.GetAllItems)
	mux.HandleFunc("/items/create", itemsHandlers.CreateItem)
	mux.HandleFunc("/items/{id}/remove", itemsHandlers.RemoveItem)

	logger.Info("Starts server at 8000 port")
	log.Fatalln(http.ListenAndServe(":8000", mux))
}
