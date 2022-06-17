package membership

type CreateRequest struct {
	UserName       string
	MembershipType string
}

type CreateResponse struct {
	Code           int
	Message        string
	ID             string
	MembershipType string
}

type UpdateRequest struct {
	ID             string
	UserName       string
	MembershipType string
}

type UpdateResponse struct {
	ID             string
	UserName       string
	MembershipType string
}

type GetResponse struct {
	Code           int
	Message        string
	ID             string
	UserNames      string
	MembershipType string
}
