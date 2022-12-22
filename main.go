package main

import (
	"fmt"
	"time"
	"tokoku/actTransaksi"
	"tokoku/barang"
	"tokoku/config"
	"tokoku/pelanggan"
	"tokoku/transaksi"
	"tokoku/user"
)

func main() {
	var inputMenu int = 1
	var cfg = config.ReadConfig()
	var conn = config.ConnectSQL(*cfg)
	var authMenu = user.AuthMenu{DB: conn}
	var barangMenu = barang.BarangMenu{DB: conn}
	var pelangganMenu = pelanggan.PelangganMenu{DB: conn}
	var TransaksiMenu = transaksi.TransaksiMenu{DB: conn}
	var ActTransaksiMenu = actTransaksi.ActTransaksiMenu{DB: conn}

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
						now := time.Now()
						fmt.Println("## Menu Admin ##")
						fmt.Println("Selamat datang, ", userRes.Nama, "-", now.Format("Jan 2, 2006"))
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
							var idPeg int
							fmt.Println("Berikut daftar pegawai Tokoku :")
							authMenu.ShowUser()
							fmt.Println()
							fmt.Print("Masukan id yang mau dihapus : ")
							fmt.Scanln(&idPeg)

							res, err := authMenu.HapusPegawai(idPeg)
							if err != nil {
								fmt.Println(err.Error())
							}

							if res {
								fmt.Println("Sukses menghapus data pegawai")
							} else {
								fmt.Println("Gagal menghapus data pegawai")
							}

						} else if adminMenu == 3 {
							var idbrg int
							fmt.Println("Berikut daftar barang Tokoku :")
							barangMenu.Showbarang()
							fmt.Println()
							fmt.Print("Masukan id yang mau dihapus : ")
							fmt.Scanln(&idbrg)

							res, err := barangMenu.Hapusbarang(idbrg)
							if err != nil {
								fmt.Println(err.Error())
							}

							if res {
								fmt.Println("Sukses menghapus data barang")
							} else {
								fmt.Println("Gagal menghapus data barang")
							}

						} else if adminMenu == 4 {
							var idPel int
							fmt.Println("Berikut daftar pelanggan Tokoku :")
							pelangganMenu.TampilkanPelanggan()
							fmt.Println()
							fmt.Print("Masukan id yang mau dihapus : ")
							fmt.Scanln(&idPel)

							res, err := pelangganMenu.HapusPelanggan(idPel)
							if err != nil {
								fmt.Println(err.Error())
							}

							if res {
								fmt.Println("Sukses menghapus data pelanggan")
							} else {
								fmt.Println("Gagal menghapus data pelanggan")
							}

						} else if adminMenu == 5 {
							var idTrs int
							fmt.Println("Berikut daftar transaksi Tokoku :")
							TransaksiMenu.TampilkanTransaksi()
							fmt.Println()
							fmt.Print("Masukan id yang mau dihapus : ")
							fmt.Scanln(&idTrs)

							res, err := TransaksiMenu.HapusTransaksi(idTrs)
							if err != nil {
								fmt.Println(err.Error())
							}

							if res {
								fmt.Println("Sukses menghapus data transaksi")
							} else {
								fmt.Println("Gagal menghapus data transaksi")
							}

						} else if adminMenu == 9 {
							isLogin = false
						}

					} else if role > 1 {
						pegawaiMenu := 0
						now := time.Now()
						fmt.Println("## Menu Pegawai ##")
						fmt.Println("Selamat datang, ", userRes.Nama, "-", now.Format("Jan 2, 2006"))
						fmt.Println("1. Data Barang")
						fmt.Println("2. Data Pelanggan")
						fmt.Println("3. Data Transaksi")
						fmt.Println("9. Logout")
						fmt.Print("Masukkan menu : ")
						fmt.Scanln(&pegawaiMenu)

						if pegawaiMenu == 1 {

							brgMenu := 0
							isBrgMenu := true
							for isBrgMenu {
								fmt.Println("## Data Barang ##")
								fmt.Println("1. Tambah Barang")
								fmt.Println("2. Update Barang")
								fmt.Println("3. Lihat Barang")
								fmt.Println("9. Kembali")
								fmt.Print("Masukkan menu : ")
								fmt.Scanln(&brgMenu)

								if brgMenu == 1 {
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

									// newBarang.ID = brgRes
									if brgRes != 0 {
										fmt.Println("Sukses menambahkan barang")
									} else {
										fmt.Println("Gagal menambahkan barang")
									}
								} else if brgMenu == 2 {
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
								} else if brgMenu == 3 {
									fmt.Println("Berikut daftar Barang Tokoku :")
									barangMenu.TampilkanBarang()
									fmt.Print("Tekan enter untuk melanjutkan : ")
									fmt.Scanln()
								} else if brgMenu == 9 {
									isBrgMenu = false
								}
							}

						} else if pegawaiMenu == 2 {
							plgMenu := 0
							isPlgMenu := true
							for isPlgMenu {
								fmt.Println("## Data Pelanggan ##")
								fmt.Println("1. Tambah Pelanggan")
								fmt.Println("2. Lihat Pelanggan")
								fmt.Println("9. Kembali")
								fmt.Print("Masukkan menu : ")
								fmt.Scanln(&plgMenu)

								if plgMenu == 1 {
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

									// newPelanggan.ID = plgRes
									if plgRes != 0 {
										fmt.Println("Sukses menambahkan pelanggan")
									} else {
										fmt.Println("Gagal menambahkan pelanggan")
									}
								} else if plgMenu == 2 {
									fmt.Println("Berikut daftar pelanggan Tokoku :")
									pelangganMenu.TampilkanPelanggan()
									fmt.Print("Tekan enter untuk melanjutkan : ")
									fmt.Scanln()
								} else if plgMenu == 9 {
									isPlgMenu = false
								}
							}

						} else if pegawaiMenu == 3 {
							trsMenu := 0
							isTrsMenu := true
							for isTrsMenu {
								fmt.Println("## Data Transaksi ##")
								fmt.Println("1. Tambah Transaksi")
								fmt.Println("2. Lihat Transaksi")
								fmt.Println("9. Kembali")
								fmt.Print("Masukkan menu : ")
								fmt.Scanln(&trsMenu)

								if trsMenu == 1 {
									var newTransaksi transaksi.Transaksi
									var jwb string
									//Input Transaksi
									fmt.Println("Berikut pelanggan yang tersedia")
									pelangganMenu.Showpelanggan()
									fmt.Print("Masukan id pelanggan : ")
									fmt.Scanln(&newTransaksi.Pelanggan_id)
									newTransaksi.Pegawai = userRes.ID
									trsRes, err := TransaksiMenu.TambahTransaksi(newTransaksi)
									if err != nil {
										fmt.Println(err.Error())
									}
									newTransaksi.ID = trsRes
									if trsRes != 0 {
										fmt.Println("Transaksi : ", newTransaksi.Pelanggan_id)

										//Aktifitas Transaksi
										var newActTransaksi actTransaksi.ActTransaksi
										for jwb != "n" {
											barangMenu.Showbarang()
											fmt.Println()
											fmt.Print("Masukan id barang : ")
											fmt.Scanln(&newActTransaksi.Barang_id)
											fmt.Print("Masukan jumlah barang : ")
											fmt.Scanln(&newActTransaksi.Qty)

											newActTransaksi.Transaksi_id = trsRes
											ActTransaksiMenu.TambahActTransaksi(newActTransaksi)
											TransaksiMenu.UpdateStock(newActTransaksi.Barang_id, newActTransaksi.Qty)
											fmt.Print("Tambah barang ? (y/n) :  ")
											fmt.Scanln(&jwb)

											if jwb == "n" {
												var ctkNota string
												fmt.Print("Cetak Nota ? (y/n) :  ")
												fmt.Scanln(&ctkNota)

												if ctkNota == "y" {
													nota := TransaksiMenu.CetakNota(newActTransaksi.Transaksi_id)
													fmt.Println("NOTA TOKOKU")
													fmt.Println("Tanggal Transaksi : ", nota.Tgl_transaksi)
													fmt.Println("Nomor Transaksi : ", nota.ID)
													fmt.Println("Nama Pegawai : ", nota.Pegawai)
													for i := 0; i < len(nota.Nama_barang); i++ {
														fmt.Println("NO", "\t Nama Barang", "\t Qty")
														fmt.Println(i+1, "\t", nota.Nama_barang[i], "\t", nota.Qty[i])
													}

												}
											}
										}

									} else {
										fmt.Println("Gagal menambahkan Transaksi")
									}

								} else if trsMenu == 2 {
									fmt.Println("Berikut daftar transaksi Tokoku :")
									TransaksiMenu.TampilkanTransaksi()
									fmt.Print("Tekan enter untuk melanjutkan : ")
									fmt.Scanln()
								} else if trsMenu == 9 {
									isTrsMenu = false
								}

							}
						} else if pegawaiMenu == 9 {
							isLogin = false
						}
					}

				}
			}
		}
	}
}
