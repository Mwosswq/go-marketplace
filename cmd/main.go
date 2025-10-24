package main

import (
	"database/sql"
	"log"
	itemHandlers "main/internal/handlers/items"
	//usersHandlers "main/internal/handlers/users"
	itemRepository "main/internal/repository/items"
	//usersRepository "main/internal/repository/users"
	itemServices "main/internal/services/items"
	//usersServices "main/internal/services/users"
	"main/pkg/responder"
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
	//v := validator.New(logger)

	//Items
	itemsRepo := itemRepository.NewItemsRepository(db)
	itemsServices := itemServices.NewItemService(itemsRepo, logger)
	itemsHandlers := itemHandlers.NewItemsHandler(itemsServices, logger, r)

	//Users
	//usersRepo := usersRepository.NewUserRepository(db)
	//usersServices := usersServices.NewUsersService(usersRepo, logger, v)
	//usersHandlers := usersHandlers.NewUserHandler(usersServices, r)

	//mux, routes and server

	mux := http.NewServeMux()

	//items endpoints
	mux.HandleFunc("/items", itemsHandlers.GetAllItems)
	mux.HandleFunc("/items/create", itemsHandlers.CreateItem)
	mux.HandleFunc("/items/{id}/remove", itemsHandlers.RemoveItem)

	//users endpoints
	//mux.HandleFunc("/auth/register", usersHandlers.RegisterUser)
	//mux.HandleFunc("/auth/login", usersHandlers.LoginUser)

	logger.Info("Starts server at 8000 port")
	log.Fatalln(http.ListenAndServe(":8000", mux))
}
