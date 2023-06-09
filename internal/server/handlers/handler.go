package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/DmitryOdintsov/awesomeProject/internal/entity"
	"github.com/DmitryOdintsov/awesomeProject/internal/service"
	"github.com/go-chi/chi/v5"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	Service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{Service: s}
}

func (h *Handler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	content, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	var user entity.User

	if err = json.Unmarshal(content, &user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	userSave, err := h.Service.SaveUserService(&user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusCreated)
	idUser, err := json.Marshal(map[string]int{"id": userSave.ID})
	_, err = w.Write(idUser)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	return
	w.WriteHeader(http.StatusBadRequest)
}

type ID struct {
	SourceID int `json:"source_id"`
	TargetID int `json:"target_id"`
}

func (h Handler) AddFriendsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	content, _ := io.ReadAll(r.Body)

	defer r.Body.Close()
	var id ID
	err := json.Unmarshal(content, &id)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	sourceUser, ok := h.Service.GetUserIDService(id.SourceID)
	if !ok {
		log.Fatal("пользователь не найден")
		return
	}

	targetUser, ok := h.Service.GetUserIDService(id.TargetID)
	if !ok {
		log.Fatal("пользователь не найден")
		return

	}

	ok = sourceUser.AddFriend(targetUser)
	if !ok {
		log.Fatalln("не удалось добавить друга")
	}
	us1, _ := json.Marshal(sourceUser.Name)

	us2, _ := json.Marshal(targetUser.Name)

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(fmt.Sprintf("`%s и %s теперь друзья`", us1, us2)))
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	return

	w.WriteHeader(http.StatusBadRequest)
}

func (h *Handler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")
	idInt, _ := strconv.Atoi(id)
	users := h.Service.GetUsersService()

	userRemote, ok := h.Service.GetUserIDService(idInt)
	if !ok {
		log.Fatal("пользователя с таким ID нет")
		return
	}
	delete(users, idInt)

	for _, user := range users {
		if user.Friends != nil {
			for idUser, friend := range user.Friends {
				if friend.ID == idInt {
					user.DeleteFriend(idUser)
				}
			}
		}
	}

	userNameRemote, _ := json.Marshal(&userRemote.Name)
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(userNameRemote)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	return

	w.WriteHeader(http.StatusBadRequest)
}

func (h *Handler) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users := h.Service.GetUsersService()
	userByte, err := json.Marshal(&users)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(userByte)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	return
	w.WriteHeader(http.StatusBadRequest)
}

func (h *Handler) GetFriendHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")
	idInt, _ := strconv.Atoi(id)
	users, ok := h.Service.GetFriendsService(idInt)
	if !ok {
		log.Fatal("у этого пользователя нет друзей")
	}
	userByte, err := json.Marshal(&users)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(userByte)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	return
	w.WriteHeader(http.StatusBadRequest)
}

func (h *Handler) PutAgeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")
	idInt, _ := strconv.Atoi(id)
	newAge, err := io.ReadAll(r.Body)

	user, _ := h.Service.GetUserIDService(idInt)
	var userInput entity.User

	err = json.Unmarshal(newAge, &userInput)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()
	user.Age = userInput.Age
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("возраст успешно обновлен"))
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	return
	w.WriteHeader(http.StatusBadRequest)
}
