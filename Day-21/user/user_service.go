// user_service.go
package user

import "fmt"

type UserService struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (us *UserService) CreateUser(user *User) error {
	return us.userRepository.Create(user)
}

func (us *UserService) GetAllUsers() ([]User, error) {
	return us.userRepository.GetAll()
}

func (us *UserService) GetUserByID(userID int) (*User, error) {
	return us.userRepository.GetByID(userID)
}

func (us *UserService) UpdateUser(userID int, updatedUser *User) error {
	return us.userRepository.Update(userID, updatedUser)
}

func (us *UserService) DeleteUser(userID int) error {
	return us.userRepository.Delete(userID)
}

func (us *UserService) logError(message string, err error) {

	fmt.Printf("Error: %s - %v\n", message, err)
}
