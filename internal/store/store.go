package store

import (
	"github.com/DmitryOdintsov/awesomeProject/internal/entity"
)

type Store struct {
	Users map[int]*entity.User
}

func NewStore() *Store {
	return &Store{Users: make(map[int]*entity.User)}
}

func (s *Store) SaveUser(user *entity.User) (*entity.User, error) {
	id := s.getLastUserId()
	user.ID = id
	s.Users[id] = user
	return user, nil
}

func (s *Store) getLastUserId() int {
	return len(s.Users) + 1
}

func (s *Store) GetUsers() map[int]*entity.User {
	return s.Users
}

func (s *Store) GetUserID(id int) (*entity.User, bool) {
	user, ok := s.Users[id]
	return user, ok
}

func (s *Store) GetFriends(id int) ([]*entity.User, bool) {
	userByID, _ := s.GetUserID(id)
	if len(userByID.Friends) > 0 {
		return userByID.Friends, true
	}
	return nil, false
}
