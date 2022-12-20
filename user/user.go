package user

import (
	"database/sql"
	"errors"
	"log"
)

type User struct {
	ID       int
	Nama     string
	Password string
	Role_id  int
}

type AuthMenu struct {
	DB *sql.DB
}

func (am *AuthMenu) Duplicate(name string) bool {
	res := am.DB.QueryRow("SELECT id, role_id FROM users where nama = ?", name)
	var idExist int
	err := res.Scan(&idExist)
	if err != nil {
		log.Println("Result scan error", err.Error())
		return false
	}
	return true
}

func (am *AuthMenu) TambahPegawai(newUser User) (bool, error) {

	registerQry, err := am.DB.Prepare("INSERT INTO users (nama, password, role_id) values (?,?,2)")
	if err != nil {
		log.Println("prepare insert user ", err.Error())
		return false, errors.New("prepare statement insert user error")
	}

	if am.Duplicate(newUser.Nama) {
		log.Println("duplicated information")
		return false, errors.New("nama sudah digunakan")
	}

	res, err := registerQry.Exec(newUser.Nama, newUser.Password)
	if err != nil {
		log.Println("insert user ", err.Error())
		return false, errors.New("insert user error")
	}

	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after insert user ", err.Error())
		return false, errors.New("error setelah insert")
	}

	if affRows <= 0 {
		log.Println("no record affected")
		return false, errors.New("no record")
	}

	return true, nil
}

func (am *AuthMenu) Login(nama string, password string) (User, error) {
	loginQry, err := am.DB.Prepare("SELECT id FROM users WHERE nama = ? AND password = ?")
	if err != nil {
		log.Println("prepare insert user ", err.Error())
		return User{}, errors.New("prepare statement insert user error")
	}

	row := loginQry.QueryRow(nama, password)

	if row.Err() != nil {
		log.Println("login query ", row.Err().Error())
		return User{}, errors.New("tidak bisa login, data tidak ditemukan")
	}
	res := User{}
	err = row.Scan(&res.ID)

	if err != nil {
		log.Println("after login query ", err.Error())
		return User{}, errors.New("tidak bisa login, kesalahan setelah error")
	}

	res.Nama = nama
	return res, nil
}

func (am *AuthMenu) GantiPassword(newPassword string, id int) (bool, error) {
	gantiPassQry, err := am.DB.Prepare("UPDATE users SET password = ? WHERE id = ?")
	if err != nil {
		log.Println("prepare update password ", err.Error())
		return false, errors.New("prepare statement update password error")
	}

	res, err := gantiPassQry.Exec(newPassword, id)
	if err != nil {
		log.Println("update password ", err.Error())
		return false, errors.New("update password error")
	}
	// Cek berapa baris yang terpengaruh query diatas
	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after update password ", err.Error())
		return false, errors.New("error setelah update password")
	}

	if affRows <= 0 {
		log.Println("no record affected")
		return false, errors.New("no record")
	}

	return true, nil
}
