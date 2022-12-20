package barang

import (
	"database/sql"
	"errors"
	"log"
)

type Barang struct {
	ID          int
	User_ID     int
	Nama_barang string
	Stok        int
	Tgl_input   string
}

type BarangMenu struct {
	DB *sql.DB
}

func (bm *BarangMenu) insertBarang(nama string, stok int, tanggal string) (int, error) {
	insertBarang, err := bm.DB.Prepare("INSERT INTO barang (nama, stok, tgl_input) VALUES (,?,?,now())")
	// insertBarang, err := bm.DB.Prepare("INSERT INTO barang (user_id, nama_barang, stok, tgl_input) VALUES (?,?,?,now())")
	if err != nil {
		log.Println("prepare insert barang")
		return 0, errors.New("prepare statement insert barang error")
	}
	res, err := insertBarang.Exec(nama, stok, tanggal)
	// res, err := insertBarang.Exec(newBarang.User_ID, newBarang.Nama_barang, newBarang.Stok, newBarang.Tgl_input)

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
