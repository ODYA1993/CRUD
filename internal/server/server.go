package server

import (
	"github.com/DmitryOdintsov/awesomeProject/internal/server/handlers"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func Run(h *handlers.Handler) {
	router := chi.NewRouter()
	router.Post("/users", h.Create)
	router.Post("/friends", h.AddFriends)
	router.Get("/users", h.GetUsers)
	router.Get("/user/{id}/friends", h.GetFriend)
	router.Delete("/user/{id}", h.DeleteUser)
	router.Put("/user/{id}", h.PutAge)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalln(err)
	}
}
