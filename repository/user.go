package repository

import "github.com/soerjadi/golection/domain"

// UserRepository interface
type UserRepository interface {
	GetByID(id int64) (*domain.User, error)
	GetList(offset int32, limit int32) ([]*domain.User, error)
	Save(*domain.User) (*domain.User, error)
	Delete(id int64) error
	Update(user *domain.User) (*domain.User, error)
	GetByIdentityID(id int64) (*domain.User, error)
}

// AdminRole role user
const AdminRole = 1

// UserRole role for user
const UserRole = 2
