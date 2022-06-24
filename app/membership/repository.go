package membership

import "errors"

type Repository struct {
	data map[string]Membership
}

func NewRepository(data map[string]Membership) *Repository {
	return &Repository{data: data}
}

func (r *Repository) Create(membership Membership) {
	r.data[membership.ID] = membership
}

func (r *Repository) Update(request UpdateRequest) {
	r.data[request.ID] = Membership{ID: request.ID, UserName: request.UserName, MembershipType: request.MembershipType}
}

func (r *Repository) Delete(id string) error {
	if _, ok := r.data[id]; !ok {
		return errors.New("not found membership")
	}

	delete(r.data, id)
	return nil
}

func (r *Repository) GetById(id string) (Membership, error) {
	for _, membership := range r.data {
		if membership.ID == id {
			return membership, nil
		}
	}
	return Membership{}, errors.New("not found membership")
}

func (r *Repository) GetByUserName(userName string) (Membership, error) {
	for _, membership := range r.data {
		if membership.UserName == userName {
			return membership, nil
		}
	}
	return Membership{}, errors.New("not found membership")
}
