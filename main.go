package main

import (
	"errors"
	"fmt"
	"tokoku/config"
)

func main() {
	var inputMenu int
	var cfg = config.ReadConfig()
	var conn = config.ConnectSQL(*cfg)
	

	for inputMenu != 9 {
		fmt.Println("WELCOME TO TOKOKU MENU")
		fmt.Println("1. Login Admin")
		fmt.Println("2. Login Pegawai")
		fmt.Println("9. Exit Program")
		fmt.Print("Masukkan pilihan : ")
		fmt.Scanln(&inputMenu)
		if inputMenu == 1 {
			// buat var ? 
			fmt.Print("Username : ")
			fmt.Scanln(&)
			fmt.Print("Password : ")
			fmt.Scanln(&)

			if err != nil {
				fmt.Println(err.Error())
			}

			if blablabla {
				isLogin := true
				for isLogin {
					fmt.Println("1. Tambah daftar pegawai")
					fmt.Println("2. Hapus daftar pegawai")
					fmt.Println("3. Hapus daftar barang")
					fmt.Println("9. Exit") // atau kembali?
					fmt.Print("Masukkan pilihan : ")
					fmt.Scanln(&inputMenu)

					if inputMenu == 1 {
						// buat var ?
						fmt.Print("Mendaftarkan username : ")
						fmt.Scanln(&)
						fmt.Print("Mendaftarkan password : ")
						fmt.Scanln(&)


						if err != nil {
							fmt.Println(err.Error())
						}
					} else if inputMenu == 2 {
						// buat var ?
						fmt.Print("Hapus id/nama(?) pegawai : ")
						fmt.Scanln(&)
						
					} else if inputMenu == 3 {
						// buat var ?
						fmt.Print("Daftar barang yang ingin dihapus : ")
						fmt.Scanln(&)
					} else if inputMenu == 9 {
						isLogin = false
					}
				}

			} else {
				fmt.Println("Gagal untuk login")
			}

		} else if inputMenu == 2 {
			// buat var ?
			fmt.Print("Username : ")
			fmt.Scanln(&)
			fmt.Print("Password : ")
			fmt.Scanln(&)
		}
	}

}
