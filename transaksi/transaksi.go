package transaksi

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type Transaksi struct {
	ID            int
	Pelanggan_id  int
	Barang_id     int
	Tgl_transaksi string
	Total         int
}

type TransaksiMenu struct {
	DB *sql.DB
}

func (tm *TransaksiMenu) TambahTransaksi(newTransaksi Transaksi) (int, error) {
	addTransaksiQry, err := tm.DB.Prepare("INSERT INTO transaksi (pelanggan_id, barang_id, total,tgl_transaksi) values (?,?,?,now())")
	if err != nil {
		log.Println("prepare insert pelanggan ", err.Error())
		return 0, errors.New("prepare statement insert pelanggan error")
	}

	res, err := addTransaksiQry.Exec(newTransaksi.Pelanggan_id, newTransaksi.Barang_id, newTransaksi.Total)

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

func (tm *TransaksiMenu) MinStock(barang_id int) (int, error) {
	res := tm.DB.QueryRow("SELECT stok FROM barang WHERE id = ?", barang_id)
	var stok int
	err := res.Scan(&stok)
	if err != nil {
		log.Println("Result scan error", err.Error())
	}
	// fmt.Println(stok)
	return stok, nil

}

func (tm *TransaksiMenu) UpdateStock(barang_id int, total int) (int, error) {

	stok, err := tm.MinStock(barang_id)
	if err != nil {
		fmt.Println(err.Error())

	}

	updateStockQry, err := tm.DB.Prepare("UPDATE barang set stok = ? WHERE id = ?")
	if err != nil {
		log.Println("prepare update stock ", err.Error())
		return 0, errors.New("prepare statement update stock error")
	}
	// fmt.Println(stok)
	min := stok - total
	res, err := updateStockQry.Exec(min, barang_id)

	if err != nil {
		log.Println("update stok ", err.Error())
		return 0, errors.New("update stok error")
	}

	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after update stok ", err.Error())
		return 0, errors.New("error setelah update stok")
	}

	if affRows <= 0 {
		log.Println("no record affected")
		return 0, errors.New("no record")
	}

	id, _ := res.LastInsertId()

	return int(id), nil
}
