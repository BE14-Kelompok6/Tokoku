package pelanggan

import (
	"database/sql"
	"errors"
	"log"
)

type Pelanggan struct {
	User_id int
	Nama    string
	Alamat  string
}

type DaftarPelanggan struct {
	DB *sql.DB
}

func (dp *DaftarPelanggan) InsertPelanggan(newPelanggan DaftarPelanggan) (int, error) {
	insertPelanggan, err := dp.DB.Prepare("INSERT INTO pelanggan (user_id, nama, alamat) VALUES (?,?,?)")
	if err != nil {
		log.Println("prepare insert pelanggan")
		return 0, errors.New("prepare statement insert pelanggan error")
	}

	res, err := insertPelanggan.Exec(newPelanggan.User_id, newPelanggan.Nama, newPelanggan.Alamat)

	if err != nil {
		log.Println("insert barang")
		return 0, errors.New("insert barang error")
	}

	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after insert barang ", err.Error())
		return 0, errors.New("error setelah insert barang")
	}

	if affRows <= 0 {
		log.Println("no record affected")
		return 0, errors.New("no record")
	}

	id, _ := res.LastInsertId()

	return int(id), nil

}
