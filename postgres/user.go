package postgres

import (
	"database/sql"
	"log"

	"github.com/monirz/gojwt"
)

var _ gojwt.UserService = (*UserService)(nil)

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{db: db}
}

func (u *UserService) CreateUser(user *gojwt.User) (int64, error) {
	stmt := `INSERT INTO users (uuid, username, email, password, created_at, updated_at)
			 VALUES ($1, $2, $3, $4, $5, $6, $7)
			 RETURNING id`

	var id int64
	err := u.db.QueryRow(stmt, user.UUID, user.UserName, user.Email, user.Password, user.CreatedAt, user.UpdatedAt).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *UserService) FindUserByID(uuid string) (*gojwt.User, error) {
	return nil, nil
}

func (u *UserService) FindByEmail(email string) (*gojwt.User, error) {
	query := "SELECT id, uuid, username, email, password, created_at, updated_at FROM users WHERE email = $1"
	row := u.db.QueryRow(query, email)

	user := &gojwt.User{}
	err := row.Scan(&user.ID, &user.UUID, &user.UserName, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println("debug")
	return user, nil
}

func (u *UserService) GetUsers() (*[]gojwt.User, error) {
	query := "SELECT * FROM users"
	rows, err := u.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []gojwt.User{}

	for rows.Next() {
		user := gojwt.User{}
		err := rows.Scan(
			&user.ID,
			&user.UUID,
			&user.UserName,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &users, nil
}

func (u *UserService) UpdateUser(*gojwt.User) error {
	return nil
}
