package pelanggan

import (
	"database/sql"
	"errors"
	"fmt"
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

func (pm *PelangganMenu) TampilkanPelanggan() {
	rows, err := pm.DB.Query("SELECT p.id, p.nama, p.alamat, u.nama  FROM pelanggan p INNER JOIN users u ON u.id = p.user_id")
	if err != nil {
		log.Println("tampilkan pelanggan ", err.Error())
		fmt.Println(errors.New("tampilkan pelanggan error"))
	}
	fmt.Println("ID", "\tNama Pelanggan", "\tAlamat", "\tPegawai")
	for rows.Next() {
		var id int
		var nama, alamat, pegawai string
		err = rows.Scan(&id, &nama, &alamat, &pegawai)
		if err != nil {
			log.Println("tampilkan barang ", err.Error())
			fmt.Println(errors.New("tampilkan barang error"))
		}
		fmt.Println(id, "\t", nama, "\t", alamat, "\t", pegawai)
	}

}

func (pm *PelangganMenu) HapusPelanggan(id int) (bool, error) {
	hapuspelangganQry, err := pm.DB.Prepare("DELETE FROM pelanggan where id = ?")
	if err != nil {
		log.Println("prepare hapus pelanggan ", err.Error())
		return false, errors.New("prepare statement hapus pelanggan error")
	}

	res, err := hapuspelangganQry.Exec(id)
	if err != nil {
		log.Println("hapus pelanggan", err.Error())
		return false, errors.New("hapus pelanggan error")
	}

	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after hapus pelanggan ", err.Error())
		return false, errors.New("error setelah hapus pelanggan")
	}

	if affRows <= 0 {
		log.Println("no record affected")
		return false, errors.New("no record")
	}

	return true, nil
}

func (pm *PelangganMenu) Showpelanggan() {
	rows, err := pm.DB.Query("SELECT id, nama FROM pelanggan")
	if err != nil {
		log.Println("tampilkan pelanggan ", err.Error())
		fmt.Println(errors.New("tampilkan pelanggan error"))
	}
	fmt.Println("ID", "\tNama Pelanggan")
	for rows.Next() {
		var id int
		var nama string
		err = rows.Scan(&id, &nama)
		if err != nil {
			log.Println("tampilkan barang ", err.Error())
			fmt.Println(errors.New("tampilkan barang error"))
		}
		fmt.Println(id, "\t", nama)
	}

}
