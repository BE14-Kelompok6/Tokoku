package pelanggan

import (
	"database/sql"
	"errors"
	"log"
)

type Pelanggan struct {
	ID      int
	User_id int
	Nama    string
	Alamat  string
}

type PelangganMenu struct {
	DB *sql.DB
}

func (pm *PelangganMenu) insertPelanggan(newPelanggan PelangganMenu) (bool, error) {
	insertPelanggan, err := pm.DB.Prepare("INSERT INTO pelanggan (user_id, nama, alamat) VALUES (?,?,?)")
	if err != nil {
		log.Println("prepare insert pelanggan")
		return false, errors.New("prepare statement insert pelanggan error")
	}

	res, err := insertPelanggan.Exec(newPelanggan.User_id, newPelanggan.Nama, newPelanggan.Alamat)

	if err != nil {
		log.Println("insert barang")
		return false, errors.New("insert barang error")
	}

	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after insert barang ", err.Error())
		return false, errors.New("error setelah insert barang")
	}

	if affRows <= 0 {
		log.Println("no record affected")
		return false, errors.New("no record")
	}

	return true, nil

}
