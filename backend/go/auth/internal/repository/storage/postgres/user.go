package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"webscrapper/auth/internal/model"
)

type User interface {
	GetByEmailOrUsername(username string) (model.User, error)
	GetByID(id string) (model.User, error)
	Create(user model.User) (model.User, error)
}

type user struct {
	conn   *sqlx.DB
	schema string
	table  string
}

func NewUser(conn *sqlx.DB, schema string) User {
	return &user{
		conn:   conn,
		schema: schema,
		table:  "users",
	}
}

func (r *user) GetByEmailOrUsername(username string) (model.User, error) {
	query := fmt.Sprintf(`
SELECT * FROM %s.%s where username = $1 or email = $1
`, r.schema, r.table)

	var dbUser model.User
	err := r.conn.Get(&dbUser, query, username)
	if err != nil {
		return model.User{}, err
	}

	return dbUser, nil
}

func (r *user) Create(user model.User) (model.User, error) {
	query := fmt.Sprintf(`
INSERT INTO %s.%s 
(username, email, password, first_name, last_name, is_archive, is_admin) 
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, username, email, password, first_name, last_name, is_archive, is_admin, created_at, updated_at;
`, r.schema, r.table)

	var dbUser model.User
	// Create row
	row := r.conn.QueryRow(
		query,
		user.Username,
		user.Email,
		user.Password,
		user.FirstName,
		user.LastName,
		user.IsArchive,
		user.IsAdmin,
	)
	// Scan result of query to dbUser
	err := row.Scan(
		&dbUser.ID,
		&dbUser.Username,
		&dbUser.Email,
		&dbUser.Password,
		&dbUser.FirstName,
		&dbUser.LastName,
		&dbUser.IsArchive,
		&dbUser.IsAdmin,
		&dbUser.CreatedAt,
		&dbUser.UpdatedAt,
	)
	if err != nil {
		return model.User{}, err
	}

	return dbUser, nil
}

func (r *user) GetByID(id string) (model.User, error) {
	query := fmt.Sprintf(`
SELECT * FROM %s.%s where id = $1
`, r.schema, r.table)

	var dbUser model.User
	err := r.conn.Get(&dbUser, query, id)
	if err != nil {
		return model.User{}, err
	}

	return dbUser, nil
}
