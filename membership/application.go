package membership

import (
	"errors"
	"strconv"
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
	var err error

	// 사용자 이름을 입력하지 않은 경우 실패한다.
	if strings.TrimSpace(request.UserName) == "" {
		err = errors.New("UserName cannot be empty")
		return CreateResponse{}, err
	}

	// 멤버십 타입을 입력하지 않은 경우 실패한다.
	if request.MembershipType == "" {
		err = errors.New("MembershipType cannot be empty")
		return CreateResponse{}, err
	}

	// naver/toss/payco 이외의 타입을 입력한 경우 실패한다.
	if !(request.MembershipType == "naver" || request.MembershipType == "toss" || request.MembershipType == "payco") {
		err = errors.New("MembershipType not supported")
		return CreateResponse{}, err
	}

	// 이미 등록된 사용자 이름이 존재할 경우 실패한다.
	for _, membership := range memberships {
		if membership.UserName == request.UserName {
			return CreateResponse{}, errors.New("UserName already exists")
		}
	}

	// 멤버십을 생성한다.
	newId := strconv.Itoa(len(memberships))
	newMembership := Membership{ID: newId, UserName: request.UserName, MembershipType: request.MembershipType}
	memberships[newId] = newMembership

	return CreateResponse{newId, memberships[newId].MembershipType}, nil
}

func (app *Application) Update(request UpdateRequest) (UpdateResponse, error) {
	memberships := (*app).repository.data

	// 멤버십 아이디를 입력하지 않은 경우, 예외 처리한다.
	if request.ID == "" {
		return UpdateResponse{}, errors.New("ID cannot be empty")
	}

	// 사용자 이름을 입력하지 않은 경우, 예외 처리한다.
	if strings.TrimSpace(request.UserName) == "" {
		return UpdateResponse{}, errors.New("UserName cannot be empty")
	}

	// 멤버쉽 타입을 입력하지 않은 경우, 예외 처리한다.
	if request.MembershipType == "" {
		return UpdateResponse{}, errors.New("MembershipType cannot be empty")
	}

	// 주어진 멤버쉽 타입이 아닌 경우, 예외 처리한다.
	if !(request.MembershipType == "naver" || request.MembershipType == "toss" || request.MembershipType == "payco") {
		return UpdateResponse{}, errors.New("MembershipType not supported")
	}

	// 이미 등록된 사용자 이름이 존재할 경우 실패한다.
	for id, membership := range memberships {
		if id != request.ID && membership.UserName == request.UserName {
			return UpdateResponse{}, errors.New("UserName already exists")
		}
	}

	// membership 정보를 갱신한다.
	memberships[request.ID] = Membership{ID: request.ID, UserName: request.UserName, MembershipType: request.MembershipType}

	return UpdateResponse{request.ID, request.UserName, request.MembershipType}, nil
}

func (app *Application) Delete(id string) error {
	memberships := (*app).repository.data

	// id를 입력하지 않았을 때 예외 처리한다.
	if id == "" {
		return errors.New("ID cannot be empty")
	}

	// 입력한 id가 존재하지 않을 때 예외 처리한다.
	if _, ok := memberships[id]; !ok {
		return errors.New("ID not exists")
	}

	// 멤버십을 삭제한다.
	delete(memberships, id)

	return nil
}
