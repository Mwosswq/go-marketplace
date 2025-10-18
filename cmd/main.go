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
	r := responder.New()
	ctx := context.Background()

	//Items
	itemsRepo := itemRepository.NewItemsRepository(db, logger)
	itemsServices := itemServices.NewItemService(itemsRepo, logger)
	itemsHandlers := itemHandlers.NewItemsHandler(itemsServices, logger, r)

	//mux, routes and server

	mux := http.NewServeMux()

	mux.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) { itemsHandlers.GetAllItems(ctx, w, r) })
	mux.HandleFunc("/items/create", func(w http.ResponseWriter, r *http.Request) { itemsHandlers.CreateItem(ctx, w, r) })

	logger.Info("Starts server at 8000 port")
	log.Fatalln(http.ListenAndServe(":8000", mux))
}
