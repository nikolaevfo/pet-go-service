package service

import(
	"github.com/nikolaevfo/pet-go-service/internal/api"
)

type Service struct{
	api *api.Api
}

func (s Service) GetUsers() string {
	users := s.api.GetUsers()

	return users
}

func NewService(api *api.Api) *Service{
	return &Service{
		api: api,
	}
}