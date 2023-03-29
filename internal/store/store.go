package store

import (
	"github.com/DmitryOdintsov/awesomeProject/internal/entity"
	"sync"
)

type Store struct {
	users map[int]*entity.User
	id    int
	sync.Mutex
}

func NewStore() *Store {
	return &Store{users: make(map[int]*entity.User)}
}

func (s *Store) SaveUser(user *entity.User) (*entity.User, error) {
	s.Lock()
	id := s.incrementId()
	user.ID = id
	s.users[id] = user
	s.Unlock()
	return user, nil
}

func (s *Store) incrementId() int {
	s.id += 1
	return s.id
}
func (s *Store) GetUsers() map[int]*entity.User {
	return s.users
}

func (s *Store) GetUserID(id int) (*entity.User, bool) {
	user, ok := s.users[id]
	return user, ok
}

func (s *Store) GetFriends(id int) ([]*entity.User, bool) {
	userByID, _ := s.GetUserID(id)
	if len(userByID.Friends) > 0 {
		return userByID.Friends, true
	}
	return nil, false
}
