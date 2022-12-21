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

func (pm *PelangganMenu) TambahPelanggan(newPelanggan Pelanggan) (int, error) {
	insertPelangganQry, err := pm.DB.Prepare("INSERT INTO pelanggan (user_id, nama, alamat) values (?,?,?)")
	if err != nil {
		log.Println("prepare insert pelanggan ", err.Error())
		return 0, errors.New("prepare statement insert pelanggan error")
	}

	res, err := insertPelangganQry.Exec(newPelanggan.User_id, newPelanggan.Nama, newPelanggan.Alamat)

	if err != nil {
		log.Println("insert pelanggan ", err.Error())
		return 0, errors.New("insert pelanggan error")
	}

	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after insert pelanggan ", err.Error())
		return 0, errors.New("error setelah insert pelanggan")
	}

	if affRows <= 0 {
		log.Println("no record affected")
		return 0, errors.New("no record")
	}

	id, _ := res.LastInsertId()

	return int(id), nil
}
