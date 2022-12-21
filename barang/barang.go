package barang

import (
	"database/sql"
	"errors"
	"fmt"
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

func (bm *BarangMenu) UpdateBarang(newNamaBarang string, newStok int, id int) (bool, error) {
	updateBarangQry, err := bm.DB.Prepare("UPDATE barang SET nama_barang = ?, stok = ? , tgl_input = now() WHERE id = ?")
	if err != nil {
		log.Println("prepare update barang ", err.Error())
		return false, errors.New("prepare statement update barang error")
	}

	res, err := updateBarangQry.Exec(newNamaBarang, newStok, id)
	if err != nil {
		log.Println("update barang", err.Error())
		return false, errors.New("update barang error")
	}

	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after update barang ", err.Error())
		return false, errors.New("error setelah update barang")
	}

	if affRows <= 0 {
		log.Println("no record affected")
		return false, errors.New("no record")
	}

	return true, nil
}

func (bm *BarangMenu) TampilkanBarang() {
	rows, err := bm.DB.Query("SELECT b.id, b.nama_barang, b.stok, b.tgl_input, u.nama  FROM barang b INNER JOIN users u ON u.id = b.user_id")
	if err != nil {
		log.Println("tampilkan barang ", err.Error())
		fmt.Println(errors.New("tampilkan barang error"))
	}
	fmt.Println("ID", "\tNama Barang", "\tStok", "\tTanggal Input", "\t\tPegawai")
	for rows.Next() {
		var id, stok int
		var nama_barang, pegawai, tgl_input string
		err = rows.Scan(&id, &nama_barang, &stok, &tgl_input, &pegawai)
		if err != nil {
			log.Println("tampilkan barang ", err.Error())
			fmt.Println(errors.New("tampilkan barang error"))
		}
		fmt.Println(id, "\t", nama_barang, "\t", stok, "\t", tgl_input, "\t", pegawai)
	}

}

func (bm *BarangMenu) Showbarang() {
	rows, err := bm.DB.Query("SELECT id, nama_barang, stok FROM barang")
	if err != nil {
		log.Println("tampilkan barang ", err.Error())
		fmt.Println(errors.New("tampilkan barang error"))
	}
	fmt.Println("ID", "\tNama Barang", "\tStok")
	for rows.Next() {
		var id, stok int
		var nama_barang string
		err = rows.Scan(&id, &nama_barang, &stok)
		if err != nil {
			log.Println("tampilkan barang ", err.Error())
			fmt.Println(errors.New("tampilkan barang error"))
		}
		fmt.Println(id, "\t", nama_barang, "\t", stok)
	}

}
func (bm *BarangMenu) Hapusbarang(id int) (bool, error) {
	hapusBarangQry, err := bm.DB.Prepare("DELETE FROM barang where id = ?")
	if err != nil {
		log.Println("prepare hapus barang ", err.Error())
		return false, errors.New("prepare statement hapus barang error")
	}

	res, err := hapusBarangQry.Exec(id)
	if err != nil {
		log.Println("hapus barang", err.Error())
		return false, errors.New("hapus barang error")
	}

	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after hapus barang ", err.Error())
		return false, errors.New("error setelah hapus barang")
	}

	if affRows <= 0 {
		log.Println("no record affected")
		return false, errors.New("no record")
	}

	return true, nil
}
