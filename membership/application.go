package membership

import (
	"errors"
	"github.com/gofrs/uuid"
	"log"
	"strings"
)

type Application struct {
	repository Repository
}

func NewApplication(repository Repository) *Application {
	return &Application{repository: repository}
}

func (app *Application) Create(request CreateRequest) (CreateResponse, error) {
	memberships := (*app).repository.data

	if strings.TrimSpace(request.UserName) == "" {
		return CreateResponse{}, errors.New("UserName cannot be empty")
	}

	if request.MembershipType == "" {
		return CreateResponse{}, errors.New("MembershipType cannot be empty")
	}

	if !(request.MembershipType == "naver" || request.MembershipType == "toss" || request.MembershipType == "payco") {
		return CreateResponse{}, errors.New("MembershipType not supported")
	}

	for _, membership := range memberships {
		if membership.UserName == request.UserName {
			return CreateResponse{}, errors.New("UserName already exists")
		}
	}

	newUUId, err := uuid.NewV4()
	newId := newUUId.String()
	if err != nil {
		log.Println(err)
		return CreateResponse{}, errors.New("error occurred while generating id")
	}

	(*app).repository.CreateMembership(newId, request)

	return CreateResponse{newId, memberships[newId].MembershipType}, nil
}

func (app *Application) Update(request UpdateRequest) (UpdateResponse, error) {
	memberships := (*app).repository.data

	if request.ID == "" {
		return UpdateResponse{}, errors.New("ID cannot be empty")
	}

	if strings.TrimSpace(request.UserName) == "" {
		return UpdateResponse{}, errors.New("UserName cannot be empty")
	}

	if request.MembershipType == "" {
		return UpdateResponse{}, errors.New("MembershipType cannot be empty")
	}

	if !(request.MembershipType == "naver" || request.MembershipType == "toss" || request.MembershipType == "payco") {
		return UpdateResponse{}, errors.New("MembershipType not supported")
	}

	for id, membership := range memberships {
		if id != request.ID && membership.UserName == request.UserName {
			return UpdateResponse{}, errors.New("UserName already exists")
		}
	}

	(*app).repository.UpdateMembership(request)

	return UpdateResponse{request.ID, memberships[request.ID].UserName, memberships[request.ID].MembershipType}, nil
}

func (app *Application) Delete(id string) error {
	memberships := (*app).repository.data

	if id == "" {
		return errors.New("ID cannot be empty")
	}

	if _, ok := memberships[id]; !ok {
		return errors.New("ID not exists")
	}

	(*app).repository.DeleteMembership(id)

	return nil
}
