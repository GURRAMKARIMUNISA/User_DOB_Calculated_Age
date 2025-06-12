package service

import (
	"context"
	// "fmt" // Removed: "fmt" imported and not used
	"time"

	"github.com/GurramKarimunisa/go-user-api/db/sqlc" // Adjust import path
	"github.com/GurramKarimunisa/go-user-api/internal/repository" // Adjust import path

	// Import pgtype for handling database specific types
	"github.com/jackc/pgx/v5/pgtype"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// UserResponse struct to include calculated age
type UserResponse struct {
	ID   int32     `json:"id"`
	Name string    `json:"name"`
	Dob  time.Time `json:"dob"`
	Age  int       `json:"age"`
}

// CalculateAge calculates the age based on the date of birth
func CalculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()
	if now.YearDay() < dob.YearDay() {
		age--
	}
	return age
}

func (s *UserService) CreateUser(ctx context.Context, name string, dob time.Time) (db.User, error) {
	// Convert time.Time to pgtype.Date for database insertion
	dobPgType := pgtype.Date{Time: dob, Valid: true}
	arg := db.CreateUserParams{
		Name: name,
		Dob:  dobPgType, // Use pgtype.Date here
	}
	return s.userRepo.CreateUser(ctx, arg)
}

func (s *UserService) GetUserByID(ctx context.Context, id int32) (UserResponse, error) {
	user, err := s.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return UserResponse{}, err
	}
	// Convert pgtype.Date back to time.Time for age calculation and response
	userDobTime := user.Dob.Time
	age := CalculateAge(userDobTime)
	return UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Dob:  userDobTime, // Use time.Time here
		Age:  age,
	}, nil
}

func (s *UserService) ListUsers(ctx context.Context) ([]UserResponse, error) {
	users, err := s.userRepo.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	var userResponses []UserResponse
	for _, user := range users {
		// Convert pgtype.Date back to time.Time for age calculation and response
		userDobTime := user.Dob.Time
		age := CalculateAge(userDobTime)
		userResponses = append(userResponses, UserResponse{
			ID:   user.ID,
			Name: user.Name,
			Dob:  userDobTime, // Use time.Time here
			Age:  age,
		})
	}
	return userResponses, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id int32, name string, dob time.Time) (db.User, error) {
	// Convert time.Time to pgtype.Date for database update
	dobPgType := pgtype.Date{Time: dob, Valid: true}
	arg := db.UpdateUserParams{
		ID:   id,
		Name: name,
		Dob:  dobPgType, // Use pgtype.Date here
	}
	return s.userRepo.UpdateUser(ctx, arg)
}

func (s *UserService) DeleteUser(ctx context.Context, id int32) error {
	return s.userRepo.DeleteUser(ctx, id)
}
