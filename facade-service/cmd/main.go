package main

import (
	"facade-service/pkg/clients/items"
	"facade-service/pkg/responder"
	"log"
	"net/http"

	itemsHdl "facade-service/internal/handlers/items"
	itemsSrv "facade-service/internal/handlers/services/items"

	"go.uber.org/zap"
)

func main() {
	//utils
	r := responder.New()
	client := items.New(":50051")
	logger, _ := zap.NewDevelopment()

	//items

	itemsServices := itemsSrv.NewItemsService(client, logger)
	itemsHandlers := itemsHdl.NewItemsHandler(itemsServices, r)

	//server
	mux := http.NewServeMux()

	mux.HandleFunc("/items/create", itemsHandlers.CreatingItemRequest)

	log.Fatal(http.ListenAndServe(":8000", mux))

}
