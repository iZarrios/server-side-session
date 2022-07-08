package controllers

import (
	"www.github.com/iZarrios/server-side-session-golang/pkg/db"
	"www.github.com/iZarrios/server-side-session-golang/pkg/models"
)

func CreateUser(user *models.User) error {
	statement := `insert into users(name, email, password) values($1, $2, $3);`

	_, err := db.DB.Exec(statement, user.Name, user.Email, user.Password)

	return err
}

func GetUser(id string) (models.User, error) {
	var user models.User

	statement := `select * from users where id=$1;`

	rows, err := db.DB.Query(statement, id)
	if err != nil {
		return models.User{}, err
	}

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func CheckEmail(email string, user *models.User) bool {
	statement := `select id, name, email, password from users where email=$1 limit 1;`

	rows, err := db.DB.Query(statement, email)
	if err != nil {
		return false
	}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return false
		}
	}

	return true
}
