package membership

type Repository struct {
	data map[string]Membership
}

func NewRepository(data map[string]Membership) *Repository {
	return &Repository{data: data}
}

func (repo Repository) CreateMembership(newId string, request CreateRequest) {
	memberships := repo.data
	newMembership := Membership{ID: newId, UserName: request.UserName, MembershipType: request.MembershipType}
	memberships[newId] = newMembership
}

func (repo Repository) UpdateMembership(request UpdateRequest) {
	memberships := repo.data
	memberships[request.ID] = Membership{ID: request.ID, UserName: request.UserName, MembershipType: request.MembershipType}
}

func (repo Repository) DeleteMembership(id string) {
	memberships := repo.data
	delete(memberships, id)
}
