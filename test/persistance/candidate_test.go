package persistance_test

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/soerjadi/golection/database"
	"github.com/soerjadi/golection/domain"
	"github.com/soerjadi/golection/persistance"
	"github.com/stretchr/testify/assert"
)

func TestSaveCandidate(t *testing.T) {
	conn := database.RDB().DBTest()
	db := database.DBTestRepository(conn)

	db.Clean("candidates")

	repo := persistance.CandidateRepositoryWithRDB(conn)

	candidate := &domain.Candidate{
		ID:   1,
		Name: "Person1",
		// Desc:      "",
		// VoteCount: 0,
	}

	newCandidate, err := repo.Save(candidate)

	assert.NoError(t, err)
	assert.Equal(t, newCandidate, candidate)
}

func TestUpdateCandidate(t *testing.T) {

	conn := database.RDB().DBTest()
	db := database.DBTestRepository(conn)

	db.Clean("candidates")

	repo := persistance.CandidateRepositoryWithRDB(conn)

	candidate := &domain.Candidate{
		ID:   1,
		Name: "Person2",
		// Desc:      "",
		// VoteCount: 0,
	}

	_, err := repo.Save(candidate)

	if err != nil {
		t.Fatalf("error when save candidate : %#v", err)
	}

	updateCandidate := &domain.Candidate{
		ID:   1,
		Name: "Person1",
		// Desc:      "",
		VoteCount: sql.NullInt64{Int64: 1, Valid: true},
	}

	newCandidate, err := repo.Update(updateCandidate)

	if err != nil {
		t.Fatalf("got error when update %#v", err)
	}

	if reflect.DeepEqual(newCandidate, candidate) {
		t.Errorf("candidate not updated old : %#v, new : %#v", candidate, newCandidate)
	}

	assert.NoError(t, err)

}

func TestDeleteCandidate(t *testing.T) {
	conn := database.RDB().DBTest()
	db := database.DBTestRepository(conn)

	db.Clean("candidates")

	repo := persistance.CandidateRepositoryWithRDB(conn)

	candidate := &domain.Candidate{
		ID:   1,
		Name: "Person2",
		// Desc:      "",
		// VoteCount: 0,
	}

	_, err := repo.Save(candidate)

	if err != nil {
		t.Fatalf("error when save candidate : %#v", err)
	}

	err = repo.Delete(int64(1))

	assert.NoError(t, err)
}

func TestGetListCandidate(t *testing.T) {
	conn := database.RDB().DBTest()
	db := database.DBTestRepository(conn)

	db.Clean("candidates")

	repo := persistance.CandidateRepositoryWithRDB(conn)

	candidates := []*domain.Candidate{
		{
			ID:   1,
			Name: "Person1",
			// Desc:      "",
			// VoteCount: 0,
		},
		{
			ID:   2,
			Name: "Person2",
			// Desc:      "",
			// VoteCount: 0,
		},
	}

	for _, candidate := range candidates {
		_, _ = repo.Save(candidate)
	}

	candidates, err := repo.GetList(0, 10)

	assert.NoError(t, err)
	assert.Equal(t, len(candidates), 2)
}

func TestGetByIDCandidate(t *testing.T) {
	conn := database.RDB().DBTest()
	db := database.DBTestRepository(conn)

	db.Clean("candidates")

	repo := persistance.CandidateRepositoryWithRDB(conn)

	candidate := &domain.Candidate{
		ID:   1,
		Name: "Person2",
		// Desc:      "",
		// VoteCount: 0,
	}

	_, err := repo.Save(candidate)

	if err != nil {
		t.Fatal("error when save candidate")
	}

	c, err := repo.GetByID(int64(1))

	if err != nil {
		t.Fatalf("want error %#v, got %#v", nil, err)
	}

	if !reflect.DeepEqual(c.ID, candidate.ID) {
		t.Errorf("want %d, got %d", candidate.ID, c.ID)
	}
}
