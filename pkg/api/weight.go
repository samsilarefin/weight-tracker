package api

// UserService contains the methods of the user service
type WeightService interface{}

// UserRepository is what lets our service do db operations without knowing anything about the implementation
type WeightRepository interface{}

type weightService struct {
	storage WeightRepository
}

func NewWeightService(userRepo UserRepository) UserService {
	return &userService{
		storage: userRepo,
	}
}
