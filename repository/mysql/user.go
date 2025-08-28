package mysql

import (
	"database/sql"
	"fmt"
	"game/const/errormessage"
	"game/entity"
	"game/pkg/richerror"
	_ "github.com/go-sql-driver/mysql"
)

func (d MySqlDb) IsUniquePhoneNumber(phoneNumber string) (bool, error) {
	const op = "mysql.IsUniquePhoneNumber"
	row := d.db.QueryRow(`SELECT * FROM users WHERE phone_number = ?`, phoneNumber)

	_, err := ScanUser(row)
	if err != nil {

		if err == sql.ErrNoRows {
			return true, nil
		}
		return false, richerror.
			New(op).
			WithError(err).
			WithMassage(errormessage.CantScanQueryResult).
			WithKind(richerror.KindUnExepted)
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
	const op = "mysql.GetUserByPhoneNumber"

	row := d.db.QueryRow(`SELECT * FROM users WHERE phone_number = ?`, phoneNumber)
	user, err := ScanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, false, richerror.
				New(op).
				WithError(err).
				WithMassage(errormessage.NotFound)
		}
		return entity.User{}, false, richerror.
			New(op).
			WithError(err).
			WithMassage(errormessage.CantScanQueryResult)
	}
	return user, true, nil
}

func (d MySqlDb) GetUserByID(userID uint) (entity.User, error) {
	const op = "mysql.GetUserByID"

	row := d.db.QueryRow(`SELECT * FROM users WHERE id = ?`, userID)

	user, err := ScanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {

			return entity.User{},
				richerror.
					New(op).
					WithError(err).
					WithMassage(errormessage.NotFound).
					WithKind(richerror.KindNotFound)
		}
		return entity.User{},
			richerror.
				New(op).
				WithError(err).
				WithMassage(errormessage.CantScanQueryResult).
				WithKind(richerror.KindUnExepted)
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
