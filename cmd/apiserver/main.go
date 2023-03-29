package main

import (
	"github.com/DmitryOdintsov/awesomeProject/internal/server"
	"github.com/DmitryOdintsov/awesomeProject/internal/server/handlers"
	"github.com/DmitryOdintsov/awesomeProject/internal/service"
	"github.com/DmitryOdintsov/awesomeProject/internal/store"
)

func main() {
	str := store.NewStore()
	serv := service.NewService(str)
	hand := handlers.NewHandler(serv)

	server.Run(hand)

}
