package user

import (
	"database/sql"
	"errors"
	"fmt"
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
		return User{}, errors.New("nama pengguna tidak terdaftar")
	}

	res.Nama = nama
	return res, nil
}

func (am *AuthMenu) HapusPegawai(id int) (bool, error) {
	hapusQry, err := am.DB.Prepare("DELETE FROM users where id = ?")
	if err != nil {
		log.Println("prepare delete user", err.Error())
		return false, errors.New("prepare statement delete user error")
	}

	res, err := hapusQry.Exec(id)
	if err != nil {
		log.Println("delete user ", err.Error())
		return false, errors.New("delete user error")
	}

	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after delete user ", err.Error())
		return false, errors.New("error setelah delete")
	}

	if affRows <= 0 {
		log.Println("no record affected")
		return false, errors.New("no record")
	}

	return true, nil
}

func (am *AuthMenu) ShowUser() {
	rows, err := am.DB.Query("SELECT id, nama FROM users WHERE role_id = 2")
	if err != nil {
		log.Println("tampilkan user ", err.Error())
		fmt.Println(errors.New("tampilkan barang error"))
	}

	fmt.Println("ID", "\tNama Pegawai")
	for rows.Next() {
		var id int
		var nama string
		err = rows.Scan(&id, &nama)
		if err != nil {
			log.Println("tampilkan barang ", err.Error())
			fmt.Println(errors.New("tampilkan barang error"))
		}
		fmt.Println(id, "\t", nama)
	}

}
