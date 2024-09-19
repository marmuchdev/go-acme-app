package user

import (
	"acme/model"
	"errors"
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type PostgresUserRepository struct {
	DB *sqlx.DB
}

func NewPostgresUserRepository(db *sqlx.DB) *PostgresUserRepository {
	return &PostgresUserRepository{DB: db}
}

func (repo *PostgresUserRepository) GetUsers() ([]model.User, error) {

	users := []model.User{}

	err := sqlx.Select(repo.DB, &users, "SELECT * FROM users")
	if err != nil {
		fmt.Println("Error querying the database:", err)
		return []model.User{}, errors.New("Database could not be queried")
	}

	return users, nil
}

// GetUser retrieves a user by ID from the database.
func (repo *PostgresUserRepository) GetUser(id int) (model.User, error) {
	user := []model.User{}
	err := sqlx.Select(repo.DB, &user, "SELECT * FROM users WHERE id = ($1)", strconv.Itoa(id))
	if err != nil {
		fmt.Println("Error querying the database:", err)
		return model.User{}, errors.New("Database could not be queried")
	}

	return user[0], nil
}

func (repo *PostgresUserRepository) AddUser(user model.User) (id int, err error) {

	err = repo.DB.QueryRow("INSERT INTO users (name) VALUES ($1) RETURNING id", user.Name).Scan(&id)
	if err != nil {
		fmt.Println("Error inserting user into the database:", err)
		return 0, errors.New("Could not insert user")
	}

	return id, nil
}
func (repo *PostgresUserRepository) UpdateUser(id int, user *model.User) (model.User, error) {
	query := "UPDATE users SET name = ($1) WHERE id = ($2) RETURNING id, name"
	rows, err := repo.DB.Queryx(query, user.Name, id)

	if err != nil {
		fmt.Println("Error querying the database:", err)
		return model.User{}, errors.New("Database could not be queried")
	}
	defer rows.Close()

	var updatedUser []model.User

	for rows.Next() {
		var u model.User
		if err := rows.StructScan(&u); err != nil {
			return model.User{}, err
		}
		updatedUser = append(updatedUser, u)
	}
	return updatedUser[0], nil
}

func (repo *PostgresUserRepository) DeleteUser(id int) error {

	user := []model.User{}
	err := sqlx.Select(repo.DB, &user, "DELETE FROM users WHERE id = ($1)", strconv.Itoa(id))
	if err != nil {
		fmt.Println("Error querying the database:", err)
		return errors.New("Database could not be queried")
	}
	return nil
}
func (repo *PostgresUserRepository) Close() {

}
