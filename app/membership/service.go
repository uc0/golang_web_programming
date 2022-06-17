package membership

import (
	"github.com/google/uuid"
	"net/http"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (service *Service) Create(request CreateRequest) CreateResponse {
	membership := Membership{uuid.New().String(), request.UserName, request.MembershipType}
	service.repository.Create(membership)
	return CreateResponse{
		Code:           http.StatusCreated,
		Message:        http.StatusText(http.StatusCreated),
		ID:             membership.ID,
		MembershipType: membership.MembershipType,
	}
}

func (service *Service) GetByID(id string) GetResponse {
	membership, err := service.repository.GetById(id)
	if err != nil {
		return GetResponse{
			Code:    http.StatusBadRequest,
			Message: "not found id",
		}
	}
	return GetResponse{
		Code:           http.StatusOK,
		Message:        http.StatusText(http.StatusOK),
		ID:             membership.ID,
		UserNames:      membership.UserName,
		MembershipType: membership.MembershipType,
	}
}
