package persistance

import (
	"database/sql"

	"github.com/soerjadi/golection/domain"
	"github.com/soerjadi/golection/repository"
)

// UserRepositoryImpl User repository implementation
type UserRepositoryImpl struct {
	Conn *sql.DB
}

// UserRepositoryWithRDB return initialized UserRepositoryImpl
func UserRepositoryWithRDB(conn *sql.DB) repository.UserRepository {
	return &UserRepositoryImpl{Conn: conn}
}

// GetByID get user by id
func (repo *UserRepositoryImpl) GetByID(id int64) (*domain.User, error) {
	row, err := repo.queryRow("select * from users where id=$1", id)

	if err != nil {
		return nil, err
	}

	user := &domain.User{}
	err = row.Scan(&user.ID, &user.Username, &user.Fullname, &user.Passhash, &user.IdentityNumber, &user.IsVoted, &user.Role, &user.Loginable)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetByEmail get user by email
func (repo *UserRepositoryImpl) GetByEmail(email string) (*domain.User, error) {
	row, err := repo.queryRow("select * from users where email=$1", email)

	if err != nil {
		return nil, err
	}

	user := &domain.User{}
	err = row.Scan(&user.ID, &user.Username, &user.Fullname, &user.Passhash, &user.IdentityNumber, &user.IsVoted, &user.Role, &user.Loginable)

	return user, nil
}

// GetByIdentityID get user by identity ID
func (repo *UserRepositoryImpl) GetByIdentityID(id int64) (*domain.User, error) {
	row, err := repo.queryRow("select * from users where identity_number=$1", id)

	if err != nil {
		return nil, err
	}

	user := &domain.User{}
	err = row.Scan(&user.ID, &user.Username, &user.Fullname, &user.Passhash, &user.IdentityNumber, &user.IsVoted, &user.Role, &user.Loginable)

	return user, nil
}

// GetList get user list
func (repo *UserRepositoryImpl) GetList(offset int32, limit int32) ([]*domain.User, error) {
	rows, err := repo.query("select * from users limit $1 offset $2", limit, offset)
	// rows, err := repo.Conn.Query("select * from users limit $1 offset $2", limit, offset)

	if err != nil {
		return nil, err
	}

	users := make([]*domain.User, 0)
	for rows.Next() {
		user := &domain.User{}
		err = rows.Scan(&user.ID, &user.Username, &user.Fullname, &user.Passhash, &user.IdentityNumber, &user.IsVoted, &user.Role, &user.Loginable)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// Save this method to save user to DB
func (repo *UserRepositoryImpl) Save(u *domain.User) (*domain.User, error) {
	stmt, err := repo.Conn.Prepare("insert into users(username, fullname, passhash, identity_number, is_voted, role, loginable) values ($1, $2, $3, $4, $5, $6, $7) returning id")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var ID int64
	err = stmt.QueryRow(u.Username, u.Fullname, u.Passhash, u.IdentityNumber, u.IsVoted, u.Role, u.Loginable).Scan(&ID)

	if err != nil {
		return nil, err
	}

	newUser, err := repo.GetByID(ID)

	return newUser, err
}

// Update this method to update user with new record
func (repo *UserRepositoryImpl) Update(user *domain.User) (*domain.User, error) {
	stmt, err := repo.Conn.Prepare("update users set username=$1, fullname=$2, passhash=$3, identity_number=$4, is_voted=$5, role=$6, loginable=$7 where id=$8 returning id")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var ID int64
	err = stmt.QueryRow(user.Username, user.Fullname, user.Passhash, user.IdentityNumber, user.IsVoted, user.Role, user.Loginable, user.ID).Scan(&ID)

	if err != nil {
		return nil, err
	}

	return user, err
}

// Delete delete user by id
func (repo *UserRepositoryImpl) Delete(id int64) error {
	stmt, err := repo.Conn.Prepare("delete from users where id = $1")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	return err
}

func (repo *UserRepositoryImpl) query(q string, args ...interface{}) (*sql.Rows, error) {
	stmt, err := repo.Conn.Prepare(q)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	return stmt.Query(args...)
}

func (repo *UserRepositoryImpl) queryRow(q string, args ...interface{}) (*sql.Row, error) {
	stmt, err := repo.Conn.Prepare(q)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.QueryRow(args...), nil
}
