package main

import (
	"fmt"
	"tokoku/barang"
	"tokoku/config"
	"tokoku/pelanggan"
	"tokoku/user"
)

func main() {
	var inputMenu int = 1
	var cfg = config.ReadConfig()
	var conn = config.ConnectSQL(*cfg)
	var authMenu = user.AuthMenu{DB: conn}
	var barangMenu = barang.BarangMenu{DB: conn}
	var pelangganMenu = pelanggan.PelangganMenu{DB: conn}

	for inputMenu != 0 {
		fmt.Println("## TOKOKU ##")
		fmt.Println("1. Login")
		fmt.Println("0. Exit")
		fmt.Print("Masukan Pilihan : ")
		fmt.Scanln(&inputMenu)
		if inputMenu == 1 {
			var inputNama, inputPassword string
			fmt.Print("Masukkan nama : ")
			fmt.Scanln(&inputNama)
			fmt.Print("Masukkan password : ")
			fmt.Scanln(&inputPassword)

			userRes, err := authMenu.Login(inputNama, inputPassword)
			if err != nil {
				fmt.Println(err.Error())
			}

			if userRes.ID > 0 {
				isLogin := true
				adminMenu := 0
				role := userRes.ID
				for isLogin {
					if role == 1 {
						fmt.Println("## Menu Admin ##")
						fmt.Println("Selamat datang", userRes.Nama)
						fmt.Println("1. Tambah Pegawai")
						fmt.Println("2. Hapus Data Pegawai")
						fmt.Println("3. Hapus Data Barang")
						fmt.Println("4. Hapus Data Pelanggan")
						fmt.Println("5. Hapus Data Transaksi")
						fmt.Println("9. Logout")
						fmt.Print("Masukkan menu : ")
						fmt.Scanln(&adminMenu)
						if adminMenu == 1 {
							var newUser user.User
							fmt.Print("Masukkan nama : ")
							fmt.Scanln(&newUser.Nama)
							fmt.Print("Masukkan password : ")
							fmt.Scanln(&newUser.Password)
							res, err := authMenu.TambahPegawai(newUser)
							if err != nil {
								fmt.Println(err.Error())
							}
							if res {
								fmt.Println("Sukses mendaftarkan data")
							} else {
								fmt.Println("Gagal mendaftarn data")
							}

						} else if adminMenu == 2 {
							fmt.Println("Halo")
							fmt.Println("Nama :", userRes.Nama)
						} else if adminMenu == 3 {

						} else if adminMenu == 9 {
							isLogin = false
						}
					} else if role > 1 {
						pegawaiMenu := 0
						fmt.Println("## Menu Pegawai ##")
						fmt.Println("Selamat datang", userRes.Nama)
						fmt.Println("1. Tambah Barang")
						fmt.Println("2. Update Barang")
						fmt.Println("3. Lihat Barang")
						fmt.Println("4. Data Pelanggan")
						fmt.Println("5. Lihat Pelanggan")
						fmt.Println("6. Transaksi")
						fmt.Println("9. Logout")
						fmt.Print("Masukkan menu : ")
						fmt.Scanln(&pegawaiMenu)
						if pegawaiMenu == 1 {
							var newBarang barang.Barang
							fmt.Println("## Tambah Barang ##")
							fmt.Print("Masukan nama barang : ")
							fmt.Scanln(&newBarang.Nama_barang)
							fmt.Print("Masukan stok barang : ")
							fmt.Scanln(&newBarang.Stok)
							newBarang.User_id = userRes.ID

							brgRes, err := barangMenu.TambahBarang(newBarang)
							if err != nil {
								fmt.Println(err.Error())
							}
							newBarang.ID = brgRes
							if brgRes != 0 {
								fmt.Println("Sukses menambahkan barang")
							} else {
								fmt.Println("Gagal menambahkan barang")
							}

						} else if pegawaiMenu == 2 {
							var newNamaBarang string
							var newStok, idBrg int
							fmt.Println("## Update Barang ##")
							fmt.Println("Berikut list barang Tokoku :")
							barangMenu.Showbarang()
							fmt.Print("Masukkan id : ")
							fmt.Scanln(&idBrg)
							fmt.Print("Masukan nama barang : ")
							fmt.Scanln(&newNamaBarang)
							fmt.Print("Masukan stok barang : ")
							fmt.Scanln(&newStok)
							uptBrg, err := barangMenu.UpdateBarang(newNamaBarang, newStok, idBrg)
							if err != nil {
								fmt.Println(err.Error())
							}

							if uptBrg {
								fmt.Println("Sukses mengupdate barang")
							} else {
								fmt.Println("Gagal mengupdate barang")
							}

						} else if pegawaiMenu == 3 {

							fmt.Println("Berikut daftar Barang Tokoku :")
							barangMenu.TampilkanBarang()
							fmt.Print("Tekan enter untuk melanjutkan : ")
							fmt.Scanln()
						} else if pegawaiMenu == 4 {
							var newPelanggan pelanggan.Pelanggan
							fmt.Println("## Data Pelanggan ##")
							fmt.Print("Masukan nama pelanggan   : ")
							fmt.Scanln(&newPelanggan.Nama)
							fmt.Print("Masukan alamat pelanggan : ")
							fmt.Scanln(&newPelanggan.Alamat)
							newPelanggan.User_id = userRes.ID

							plgRes, err := pelangganMenu.TambahPelanggan(newPelanggan)
							if err != nil {
								fmt.Println(err.Error())
							}
							newPelanggan.ID = plgRes
							if plgRes != 0 {
								fmt.Println("Sukses menambahkan pelanggan")
							} else {
								fmt.Println("Gagal menambahkan pelanggan")
							}

						} else if pegawaiMenu == 5 {
							fmt.Println("Berikut daftar pelanggan Tokoku :")
							pelangganMenu.TampilkanPelanggan()
							fmt.Print("Tekan enter untuk melanjutkan : ")
							fmt.Scanln()
						} else if pegawaiMenu == 6 {
							fmt.Println("## Data Transaksi ##")
						} else if pegawaiMenu == 9 {
							isLogin = false
						}
					}

				}
			}
		}
	}
}
