package actTransaksi

import (
	"database/sql"
	"errors"
	"log"
)

type ActTransaksi struct {
	Transaksi_id int
	Barang_id    int
	Qty          int
}

type ActTransaksiMenu struct {
	DB *sql.DB
}

func (atm *ActTransaksiMenu) TambahActTransaksi(newActTransaksi ActTransaksi) (bool, error) {
	addTransaksiQry, err := atm.DB.Prepare("INSERT INTO aktivitas_transaksi (transaksi_id,barang_id, qty) values (?,?,?)")
	if err != nil {
		log.Println("prepare insert pelanggan ", err.Error())
		return false, errors.New("prepare statement insert pelanggan error")
	}

	res, err := addTransaksiQry.Exec(newActTransaksi.Transaksi_id, newActTransaksi.Barang_id, newActTransaksi.Qty)

	if err != nil {
		log.Println("insert transaksi ", err.Error())
		return false, errors.New("insert transaksi error")
	}

	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after insert transaksi ", err.Error())
		return false, errors.New("error setelah insert transaksi")
	}

	if affRows <= 0 {
		log.Println("no record affected")
		return false, errors.New("no record")
	}

	// id, _ := res.LastInsertId()

	return true, nil
}
