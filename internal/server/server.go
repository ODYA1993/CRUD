package server

import (
	"github.com/DmitryOdintsov/awesomeProject/internal/server/handlers"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func Run(h *handlers.Handler) {
	router := chi.NewRouter()
	router.Post("/users", h.CreateHandler)
	router.Post("/friends", h.AddFriendsHandler)
	router.Get("/users", h.GetUsersHandler)
	router.Get("/user/{id}/friends", h.GetFriendHandler)
	router.Delete("/user/{id}", h.DeleteUserHandler)
	router.Put("/user/{id}", h.PutAgeHandler)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalln(err)
	}
}
