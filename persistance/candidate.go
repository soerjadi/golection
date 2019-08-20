package persistance

import (
	"database/sql"

	"github.com/soerjadi/golection/domain"
	"github.com/soerjadi/golection/repository"
)

// CandidateRepositoryImpl Candidate Repository Implementation
type CandidateRepositoryImpl struct {
	Conn *sql.DB
}

// CandidateRepositoryWithRDB return initialize CandidateRepository
func CandidateRepositoryWithRDB(conn *sql.DB) repository.CandidateRepository {
	return &CandidateRepositoryImpl{Conn: conn}
}

// GetByID get candidate by ID
func (repo *CandidateRepositoryImpl) GetByID(id int64) (*domain.Candidate, error) {
	row := repo.Conn.QueryRow("select * from candidates where id=$1", id)

	candidate := &domain.Candidate{}
	err := row.Scan(&candidate.ID, &candidate.Name, &candidate.Picture, &candidate.Desc, &candidate.VoteCount)

	if err != nil {
		return nil, err
	}

	return candidate, nil
}

// GetList get list candidate
func (repo *CandidateRepositoryImpl) GetList(offset int32, limit int32) ([]*domain.Candidate, error) {
	rows, err := repo.Conn.Query("select * from candidates limit $1 offset $2", limit, offset)

	candidates := make([]*domain.Candidate, 0)

	defer rows.Close()

	for rows.Next() {
		candidate := &domain.Candidate{}
		err = rows.Scan(&candidate.ID, &candidate.Name, &candidate.Picture, &candidate.Desc, &candidate.VoteCount)

		if err != nil {
			return nil, err
		}

		candidates = append(candidates, candidate)
	}

	return candidates, nil

}

// Save candidate to DB
func (repo *CandidateRepositoryImpl) Save(candidate *domain.Candidate) (*domain.Candidate, error) {
	stmt, err := repo.Conn.Prepare("insert into candidates(name, description, pic) values ($1, $2, $3) returning id")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var ID int64
	err = stmt.QueryRow(candidate.Name, candidate.Desc, candidate.Picture).Scan(&ID)

	if err != nil {
		return nil, err
	}

	newUser, err := repo.GetByID(ID)

	return newUser, err
}

// Update Candidate object
func (repo *CandidateRepositoryImpl) Update(candidate *domain.Candidate) (*domain.Candidate, error) {
	stmt, err := repo.Conn.Prepare("update candidates set name=$1, description=$2, pic=$3, vote_count=$4 where id=$5 returning id")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var ID int64
	err = stmt.QueryRow(candidate.Name, candidate.Desc, candidate.Picture, candidate.VoteCount, candidate.ID).Scan(&ID)

	if err != nil {
		return nil, err
	}

	return candidate, err
}

// Delete candidate by ID
func (repo *CandidateRepositoryImpl) Delete(id int64) error {
	stmt, err := repo.Conn.Prepare("delete from candidates where id=$1")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	return err

}
