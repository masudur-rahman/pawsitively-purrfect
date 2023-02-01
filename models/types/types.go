package types

type RegisterParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`

	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
