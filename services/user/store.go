package user

import (
	"database/sql"
	"fmt"
	"standardApi/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {

	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	u := new(types.User)

	for rows.Next() {
		u, err = ScamRowIntoUser(rows)

		if err != nil {
			return nil, err
		}

	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found ")
	}

	return u, nil
}

// Add data from row to an user variable
func ScamRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.Email,
		&user.CreateAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Store) GetUserById(id int) (*types.User, error) {
	return nil, nil
}

func (s *Store) CreateUser(u types.User) error {
	return nil
}
