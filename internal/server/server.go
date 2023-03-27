package server

import (
	"github.com/DmitryOdintsov/awesomeProject/internal/server/handlers"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func Run(h *handlers.Handler) {
	router := chi.NewRouter()
	router.Post("/create", h.Create)
	router.Post("/make_friend", h.AddFriends)
	router.Get("/get", h.GetAll)
	router.Get("/friends/{id}", h.GetFriend)
	router.Delete("/delete/{id}", h.DeletePost)
	router.Put("/put/{id}", h.PutAge)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalln(err)
	}
}
