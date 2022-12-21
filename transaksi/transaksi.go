package transaksi

import (
	"database/sql"
	"errors"
	"log"
)

type Transaksi struct {
	ID            int
	Pelanggan_id  int
	Barang_id     []int
	Tgl_transaksi string
	Total         int
}

type TransaksiMenu struct {
	DB *sql.DB
}

func (tm *TransaksiMenu) TambahTransaksi(newTransaksi Transaksi) (int, error) {
	addTransaksiQry, err := tm.DB.Prepare("INSERT INTO transaksi (pelanggan_id, barang_id, tgl_transaksi, total) values (?,?,now(),?)")
	if err != nil {
		log.Println("prepare insert pelanggan ", err.Error())
		return 0, errors.New("prepare statement insert pelanggan error")
	}

	res, err := addTransaksiQry.Exec(newTransaksi.Pelanggan_id, newTransaksi.Barang_id, newTransaksi.Tgl_transaksi, newTransaksi.Total)

	if err != nil {
		log.Println("insert transaksi ", err.Error())
		return 0, errors.New("insert transaksi error")
	}

	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after insert transaksi ", err.Error())
		return 0, errors.New("error setelah insert transaksi")
	}

	if affRows <= 0 {
		log.Println("no record affected")
		return 0, errors.New("no record")
	}

	id, _ := res.LastInsertId()

	return int(id), nil
}
