package persitance_test

import (
	"reflect"
	"testing"

	"github.com/soerjadi/golection/database"
	"github.com/soerjadi/golection/domain"
	"github.com/soerjadi/golection/persistance"
	"github.com/stretchr/testify/assert"
)

func TestSaveUser(t *testing.T) {
	conn := database.RDB().DBTest()
	db := database.DBTestRepository(conn)

	db.Clean("users")

	repo := persistance.UserRepositoryWithRDB(conn)

	user := &domain.User{
		ID:             1,
		Username:       "admin",
		Fullname:       "Administrator",
		Passhash:       "123123",
		IdentityNumber: 123123,
		IsVoted:        false,
		Loginable:      false,
	}

	newUser, err := repo.Save(user)

	assert.NoError(t, err)
	assert.Equal(t, newUser, user)
}

func TestGetByID(t *testing.T) {
	conn := database.RDB().DBTest()
	db := database.DBTestRepository(conn)

	db.Clean("users")

	repo := persistance.UserRepositoryWithRDB(conn)

	user := &domain.User{
		ID:             1,
		Username:       "admin",
		Fullname:       "Administrator",
		Passhash:       "123123",
		IdentityNumber: 123123,
		IsVoted:        false,
		Loginable:      false,
	}

	_, err := repo.Save(user)

	if err != nil {
		t.Fatalf("error when save user")
	}

	u, err := repo.GetByID(int64(1))

	if err != nil {
		t.Fatalf("want error %#v, got %#v", nil, err)
	}

	if !reflect.DeepEqual(u.ID, user.ID) {
		t.Errorf("want %d, got %d", user.ID, u.ID)
	}
}

func TestGetByIdentityID(t *testing.T) {
	conn := database.RDB().DBTest()
	db := database.DBTestRepository(conn)

	db.Clean("users")

	repo := persistance.UserRepositoryWithRDB(conn)

	user := &domain.User{
		ID:             1,
		Username:       "admin",
		Fullname:       "Administrator",
		Passhash:       "123123",
		IdentityNumber: 123123,
		IsVoted:        false,
		Loginable:      false,
	}

	_, err := repo.Save(user)

	if err != nil {
		t.Fatalf("error when save user")
	}

	u, err := repo.GetByIdentityID(int64(123123))

	if err != nil {
		t.Fatalf("want error %#v, got %#v", nil, err)
	}

	if !reflect.DeepEqual(u.ID, user.ID) {
		t.Errorf("want %d, got %d", user.ID, u.ID)
	}
}

func TestUpdate(t *testing.T) {
	conn := database.RDB().DBTest()
	db := database.DBTestRepository(conn)

	db.Clean("users")

	repo := persistance.UserRepositoryWithRDB(conn)

	user := &domain.User{
		ID:             1,
		Username:       "admin",
		Fullname:       "Administrator",
		Passhash:       "123123",
		IdentityNumber: 123123,
		IsVoted:        false,
		Loginable:      false,
	}

	_, err := repo.Save(user)

	if err != nil {
		t.Fatalf("error when save user")
	}

	updateUser := &domain.User{
		ID:             1,
		Username:       "admin",
		Fullname:       "Administrator",
		Passhash:       "123123",
		IdentityNumber: 1231234,
		IsVoted:        false,
		Loginable:      false,
	}

	newUser, err := repo.Update(updateUser)

	if err != nil {
		t.Fatalf("got error when update %#v", err)
	}

	if reflect.DeepEqual(newUser, user) {
		t.Errorf("user not updated old : %#v, new : %#v", user, newUser)
	}

	assert.NoError(t, err)
	if !reflect.DeepEqual(updateUser.IdentityNumber, newUser.IdentityNumber) {
		t.Errorf("want %d, got %d", updateUser.IdentityNumber, newUser.IdentityNumber)
	}
}

func TestDelete(t *testing.T) {
	conn := database.RDB().DBTest()
	db := database.DBTestRepository(conn)

	db.Clean("users")

	repo := persistance.UserRepositoryWithRDB(conn)

	user := &domain.User{
		ID:             1,
		Username:       "admin",
		Fullname:       "Administrator",
		Passhash:       "123123",
		IdentityNumber: 123123,
		IsVoted:        false,
		Loginable:      false,
	}

	_, err := repo.Save(user)

	if err != nil {
		t.Fatalf("error when save user")
	}

	err = repo.Delete(int64(1))

	assert.NoError(t, err)
}

func TestGetList(t *testing.T) {
	conn := database.RDB().DBTest()
	db := database.DBTestRepository(conn)

	db.Clean("users")

	repo := persistance.UserRepositoryWithRDB(conn)

	users := []*domain.User{
		{
			ID:             1,
			Username:       "admin",
			Fullname:       "Administrator",
			Passhash:       "123123",
			IdentityNumber: 123123,
			IsVoted:        false,
			Loginable:      false,
		},
		{
			ID:             2,
			Username:       "user",
			Fullname:       "User",
			Passhash:       "123123",
			IdentityNumber: 123123,
			IsVoted:        false,
			Loginable:      false,
		},
	}

	for _, user := range users {
		_, _ = repo.Save(user)
	}

	users, err := repo.GetList(0, 10)

	assert.NoError(t, err)
	assert.Equal(t, len(users), 2)

}
