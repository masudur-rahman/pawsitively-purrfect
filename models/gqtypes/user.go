package gqtypes

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	FullName string `json:"fullName"`
	Bio      string `json:"bio"`
	Location string `json:"location"`
	Avatar   string `json:"avatar"`
	IsActive bool   `json:"isActive"`
	IsAdmin  bool   `json:"isAdmin"`
}
