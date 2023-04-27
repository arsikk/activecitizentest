package postgres

import "github.com/jmoiron/sqlx"

type userPostgressRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *userPostgressRepo {
	return &userPostgressRepo{db: db}
}

func (r *userPostgressRepo) SaveUser() {
	q := `INSERT INTO users(username, passord, email)
         VALUES ($1,$2,$3);`

}

func (r *userPostgressRepo) FindByUsername() {
	q := ` SELECT FROM users WHERE username = $1`

}

func (r *userPostgressRepo) RefreshToken() {
	q := `SELECT FROM refresh_token`
}
