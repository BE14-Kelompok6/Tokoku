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

			res, err := authMenu.Login(inputNama, inputPassword)
			if err != nil {
				fmt.Println(err.Error())
			}

			if res.ID > 0 {
				isLogin := true
				adminMenu := 0
				for isLogin {
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

					} else if adminMenu == 2 {
						fmt.Println("Halo")
						fmt.Println("Nama :", res.Nama)
					} else if adminMenu == 3 {
						var inputPass string
						fmt.Print("Masukkan password baru")
						fmt.Scanln(&inputPass)
						isChanged, err := authMenu.GantiPassword(inputPass, res.ID)
						if err != nil {
							fmt.Println(err.Error())
						}
						if isChanged {
							fmt.Println("Berhasil ganti password")
							isLogin = false
						}
					} else if adminMenu == 9 {
						isLogin = false
					}
				}
			}
		}
	}
}
