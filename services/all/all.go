package all

import (
	"github.com/masudur-rahman/pawsitively-purrfect/infra/database/nosql"
	"github.com/masudur-rahman/pawsitively-purrfect/infra/logr"
	"github.com/masudur-rahman/pawsitively-purrfect/repos/pet"
	"github.com/masudur-rahman/pawsitively-purrfect/repos/shelter"
	"github.com/masudur-rahman/pawsitively-purrfect/repos/user"
	"github.com/masudur-rahman/pawsitively-purrfect/services"
	petsvc "github.com/masudur-rahman/pawsitively-purrfect/services/pet"
	sheltersvc "github.com/masudur-rahman/pawsitively-purrfect/services/shelter"
	usersvc "github.com/masudur-rahman/pawsitively-purrfect/services/user"
)

type Services struct {
	User    services.UserService
	Pet     services.PetService
	Shelter services.ShelterService
}

func GetNoSQLServices(db nosql.Database, logger logr.Logger) *Services {
	userRepo := user.NewNoSQLUserRepository(db, logger)
	petRepo := pet.NewNoSQLPetRepository(db, logger)
	shelterRepo := shelter.NewNoSQLShelterRepository(db, logger)

	userSvc := usersvc.NewUserService(userRepo)
	petSvc := petsvc.NewPetService(petRepo, userRepo)
	shelterSvc := sheltersvc.NewShelterService(shelterRepo)

	return &Services{
		User:    userSvc,
		Pet:     petSvc,
		Shelter: shelterSvc,
	}
}
