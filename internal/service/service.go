package service

import "github.com/DmitryOdintsov/awesomeProject/internal/store"

type Service struct {
	Store *store.Store
}

func NewService(s *store.Store) *Service {
	return &Service{Store: s}
}
