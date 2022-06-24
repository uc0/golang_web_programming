package membership

import (
	"github.com/google/uuid"
	"net/http"
	"strings"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (service *Service) Create(request CreateRequest) CreateResponse {
	if strings.TrimSpace(request.UserName) == "" {
		return CreateResponse{Code: http.StatusPaymentRequired, Message: "UserName cannot be empty"}
	}

	if request.MembershipType == "" {
		return CreateResponse{Code: http.StatusPaymentRequired, Message: "MembershipType cannot be empty"}
	}

	if !(request.MembershipType == "naver" || request.MembershipType == "toss" || request.MembershipType == "payco") {
		return CreateResponse{Code: http.StatusBadRequest, Message: "MembershipType not supported"}
	}

	existingMembership, err := service.repository.GetByUserName(request.UserName)
	if err == nil && existingMembership.UserName == request.UserName {
		return CreateResponse{Code: http.StatusBadRequest, Message: "UserName already exists"}
	}

	newMembership := Membership{uuid.New().String(), request.UserName, request.MembershipType}
	service.repository.Create(newMembership)
	return CreateResponse{
		Code:           http.StatusCreated,
		Message:        http.StatusText(http.StatusCreated),
		ID:             newMembership.ID,
		MembershipType: newMembership.MembershipType,
	}
}

func (service *Service) GetByID(id string) GetResponse {
	membership, err := service.repository.GetById(id)
	if err != nil {
		return GetResponse{Code: http.StatusBadRequest, Message: err.Error()}
	}
	return GetResponse{
		Code:           http.StatusOK,
		Message:        http.StatusText(http.StatusOK),
		ID:             membership.ID,
		UserName:       membership.UserName,
		MembershipType: membership.MembershipType,
	}
}

func (service *Service) Update(request UpdateRequest) UpdateResponse {
	if request.ID == "" {
		return UpdateResponse{Code: http.StatusBadRequest, Message: "ID cannot be empty"}
	}

	if strings.TrimSpace(request.UserName) == "" {
		return UpdateResponse{Code: http.StatusBadRequest, Message: "UserName cannot be empty"}
	}

	if request.MembershipType == "" {
		return UpdateResponse{Code: http.StatusBadRequest, Message: "MembershipType cannot be empty"}
	}

	if !(request.MembershipType == "naver" || request.MembershipType == "toss" || request.MembershipType == "payco") {
		return UpdateResponse{Code: http.StatusBadRequest, Message: "MembershipType not supported"}
	}

	existingMembership, err := service.repository.GetByUserName(request.UserName)
	if err == nil && existingMembership.ID != request.ID && existingMembership.UserName == request.UserName {
		return UpdateResponse{Code: http.StatusBadRequest, Message: "UserName already exists"}
	}

	service.repository.Update(request)
	return UpdateResponse{
		Code:           http.StatusOK,
		Message:        http.StatusText(http.StatusOK),
		ID:             request.ID,
		UserName:       request.UserName,
		MembershipType: request.MembershipType,
	}
}

func (service *Service) Delete(id string) DeleteResponse {
	if id == "" {
		return DeleteResponse{Code: http.StatusBadRequest, Message: "ID cannot be empty"}
	}

	if err := service.repository.Delete(id); err != nil {
		return DeleteResponse{Code: http.StatusBadRequest, Message: err.Error()}
	}

	return DeleteResponse{Code: http.StatusOK, Message: http.StatusText(http.StatusOK)}
}
