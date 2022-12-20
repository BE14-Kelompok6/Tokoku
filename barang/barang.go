package barang

import (
	"database/sql"
	"errors"
	"log"
)

type Barang struct {
	ID          int
	User_id     int
	Nama_barang string
	Stok        int
	Tgl_input   string
}

type BarangMenu struct {
	DB *sql.DB
}

func (bm *BarangMenu) insertBarang(newBarang Barang) (bool, error) {
	insertBarang, err := bm.DB.Prepare("INSERT INTO barang (user_id, nama_barang, stok, tgl_input) VALUES (?,?,?,now())")
	if err != nil {
		log.Println("prepare insert barang")
		return false, errors.New("prepare statement insert barang error")
	}
	res, err := insertBarang.Exec(newBarang.User_id, newBarang.Nama_barang, newBarang.Stok, newBarang.Tgl_input)

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
