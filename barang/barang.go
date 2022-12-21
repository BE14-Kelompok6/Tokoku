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

func (bm *BarangMenu) TambahBarang(newBarang Barang) (int, error) {
	insertActivityQry, err := bm.DB.Prepare("INSERT INTO barang (user_id, nama_barang, stok, tgl_input) values (?,?,?,now())")
	if err != nil {
		log.Println("prepare insert barang ", err.Error())
		return 0, errors.New("prepare statement insert barang error")
	}

	res, err := insertActivityQry.Exec(newBarang.User_id, newBarang.Nama_barang, newBarang.Stok)

	if err != nil {
		log.Println("insert barang ", err.Error())
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
