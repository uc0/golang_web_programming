package membership

import "strconv"

type Application struct {
	repository Repository
}

func NewApplication(repository Repository) *Application {
	return &Application{repository: repository}
}

func (app *Application) Create(request CreateRequest) (CreateResponse, error) {
	memberships := (*app).repository.data

	newId := strconv.Itoa(len(memberships))
	newMembership := Membership{ID: newId, UserName: request.UserName, MembershipType: request.MembershipType}
	memberships[request.UserName] = newMembership

	return CreateResponse{memberships[request.UserName].ID, memberships[request.UserName].MembershipType}, nil
}

func (app *Application) Update(request UpdateRequest) (UpdateResponse, error) {
	return UpdateResponse{}, nil
}

func (app *Application) Delete(id string) error {
	return nil
}
