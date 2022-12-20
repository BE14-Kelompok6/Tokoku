package main

import (
	"fmt"
	"tokoku/config"
	"tokoku/user"
)

func main() {
	var inputMenu int = 1
	var cfg = config.ReadConfig()
	var conn = config.ConnectSQL(*cfg)
	var authMenu = user.AuthMenu{DB: conn}

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
			// fmt.Println(userRes)

			if userRes.ID > 0 {
				isLogin := true
				adminMenu := 0
				role := userRes.ID
				for isLogin {
					if role == 1 {
						fmt.Println("## Menu Admin ##")
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
						fmt.Println("1. Tambah Barang")
						fmt.Println("2. Update Barang")
						fmt.Println("3. Data Pelanggan")
						fmt.Println("4. Transaksi")
						fmt.Println("9. Logout")
						fmt.Print("Masukkan menu : ")
						fmt.Scanln(&pegawaiMenu)
					}

				}
			}
		}
	}
}
