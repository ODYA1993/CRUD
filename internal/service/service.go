package service

import (
	"github.com/DmitryOdintsov/awesomeProject/internal/entity"
	"github.com/DmitryOdintsov/awesomeProject/internal/store"
)

type Service struct {
	Store store.IStore
}

func NewService(s store.IStore) *Service {
	return &Service{s}
}

func (s *Service) SaveUserService(user *entity.User) (*entity.User, error) {
	return s.Store.SaveUser(user)
}

func (s *Service) GetUsersService() map[int]*entity.User {
	return s.Store.GetUsers()
}

func (s *Service) GetUserIDService(id int) (*entity.User, bool) {
	return s.Store.GetUserID(id)
}

func (s *Service) GetFriendsService(id int) ([]*entity.User, bool) {
	return s.Store.GetFriends(id)
}
