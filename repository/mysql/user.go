package mysql

import (
	"database/sql"
	"fmt"
	"game/entity"
	_ "github.com/go-sql-driver/mysql"
)

func (d MySqlDb) IsUniquePhoneNumber(phoneNumber string) (bool, error) {
	//user := entity.User{}
	row := d.db.QueryRow(`SELECT * FROM users WHERE phone_number = ?`, phoneNumber)

	_, err := ScanUser(row)
	if err != nil {

		if err == sql.ErrNoRows {
			return true, nil
		}
		return false, fmt.Errorf("mysql query row scan error: %w", err)
	}
	return false, nil
}

func (d MySqlDb) Register(user entity.User) (entity.User, error) {
	res, eErr := d.db.Exec(`INSERT INTO users (name, phone_number, Password) values (?, ?, ?)`,
		user.Name, user.PhoneNumber, user.Password)
	if eErr != nil {

		return entity.User{}, fmt.Errorf("can't execute in data base:%w", eErr)
	}
	id, _ := res.LastInsertId()
	user.ID = uint(id)

	return user, nil
}

func (d MySqlDb) GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error) {
	// TODO - "It would be better to split the method that handles two different tasks."

	row := d.db.QueryRow(`SELECT * FROM users WHERE phone_number = ?`, phoneNumber)
	user, err := ScanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, false, nil
		}
		return entity.User{}, false, fmt.Errorf("mysql query row scan error: %w", err)
	}
	return user, true, nil
}

func (d MySqlDb) GetUserByID(userID uint) (entity.User, error) {
	// TODO - "It would be better to split the method that handles two different tasks."

	row := d.db.QueryRow(`SELECT * FROM users WHERE id = ?`, userID)

	user, err := ScanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {

			return entity.User{}, nil
		}
		return entity.User{}, fmt.Errorf("mysql query row scan error: %w", err)
	}
	fmt.Println("name:", user.Name)
	return user, nil

}

func ScanUser(row *sql.Row) (entity.User, error) {
	var created_at []uint8
	var user entity.User

	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.Password, &created_at)

	return user, err
}
