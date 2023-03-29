package service

import (
	"github.com/DmitryOdintsov/awesomeProject/internal/entity"
	"github.com/DmitryOdintsov/awesomeProject/internal/store"
)

type Service struct {
	Store *store.Store
}

func NewService(s *store.Store) *Service {
	return &Service{Store: s}
}

func (s *Service) SaveUser(user *entity.User) (*entity.User, error) {
	return s.Store.SaveUser(user)
}

func (s *Service) GetUsers() map[int]*entity.User {
	return s.Store.GetUsers()
}

func (s *Service) GetUserID(id int) (*entity.User, bool) {
	return s.Store.GetUserID(id)
}

func (s *Service) GetFriends(id int) ([]*entity.User, bool) {

	return s.Store.GetFriends(id)
}
