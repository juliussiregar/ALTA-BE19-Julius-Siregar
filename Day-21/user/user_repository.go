package user

import (
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create inserts a new user into the database
func (r *UserRepository) Create(user *User) error {
	return r.db.Create(user).Error
}

// GetAll all users from the database
func (r *UserRepository) GetAll() ([]User, error) {
	var users []User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetByID retrieves a user by ID from the database
func (r *UserRepository) GetByID(userID int) (*User, error) {
	var user User
	if err := r.db.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Update updates a user by ID in the database
func (r *UserRepository) Update(userID int, updatedUser *User) error {
	if err := r.db.Model(&User{}).Where("id = ?", userID).Updates(updatedUser).Error; err != nil {
		return err
	}
	return nil
}

// Delete deletes a user by ID from the database
func (r *UserRepository) Delete(userID int) error {
	if err := r.db.Delete(&User{}, userID).Error; err != nil {
		return err
	}
	return nil
}
