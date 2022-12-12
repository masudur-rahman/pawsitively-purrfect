package main

type User struct {
	FirstName string
	LastName  string
	Bio       string
	Location  string
	Avatar    string

	Username     string
	Email        string
	PasswordHash string

	IsActive bool
	IsAdmin  bool

	CreatedUnix   int64
	UpdatedUnix   int64
	LastLoginUnix int64
}

type Pet struct {
	Name           string
	Breed          string
	Gender         string
	Photo          string
	AdoptionStatus string
	ShelterID      int64
	// If UserID is set, user is the current owner,
	// but it previously belonged to the shelter with the ShelterID
	UserID int64
}

type Shelter struct {
	Name               string
	Website            string
	Location           string
	ContactInformation string
	Description        string
	Logo               string
	NumberOfPets       int64
	// Assuming a shelter can have only one owner
	OwnerID int64
}
