package users

import (
	"github.com/coby9241/frontend-service/internal/models/permissions"
	"github.com/coby9241/frontend-service/internal/models/users"
	"github.com/jinzhu/gorm"
)

// UserRepository is
type UserRepository interface {
	GetUserByUID(uid string) (*users.User, error)
}

// UserRepositoryImpl is
type UserRepositoryImpl struct {
	DB *gorm.DB
}

// NewUserRepositoryImpl is
func NewUserRepositoryImpl(storage *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: storage,
	}
}

// GetUserByUID is
func (r *UserRepositoryImpl) GetUserByUID(uid string) (*users.User, error) {
	i := users.User{}
	err := r.DB.Where("UID = ?", uid).First(&i).Error
	if err != nil {
		return nil, err
	}

	var role permissions.Role
	if err = r.DB.Model(i).Related(&role).Error; err != nil {
		return nil, err
	}

	i.Role = role

	return &i, err
}
