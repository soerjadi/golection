package repository

import "github.com/soerjadi/golection/domain"

// CandidateRepository interface
type CandidateRepository interface {
	GetByID(id int64) (*domain.Candidate, error)
	GetList(offset int32, limit int32) ([]*domain.Candidate, error)
	Save(*domain.Candidate) (*domain.Candidate, error)
	Delete(id int64) error
	Update(candidate *domain.Candidate) (*domain.Candidate, error)
}
