package main

import (
	"fmt"
	"os"
)

var patient patients
var patients_length int = 0

func main() {
	var option_choose int
	var option_valid bool = true

	defaultDataPatient(&patient) //aktifkan jika ingin menggunakan data default

	welcomeMessage()
	menuMessage()
	fmt.Scan(&option_choose)

	for option_valid {
		if option_choose == 1 {
			var show_data_option int
			fmt.Println("\nOpsi Penampilan Data : ")
			fmt.Println("1. Tampilkan semua data\n2. Filter Data\n3. Cari Data\n4. Edit Data\n5. Hapus Data\n6. Kembali")
			fmt.Print("Masukan : ")
			fmt.Scan(&show_data_option)

			if show_data_option == 1 {
				var sort_col int
				fmt.Println("\nOpsi Pengurutan :\n1. Data Terakhir Ditambahkan\n2. Nama\n3. Umur\n4. Email\n5. Nomor Telepon\n6. Bobot\n7. Kembali")
				fmt.Print("Masukan : ")
				fmt.Scan(&sort_col)

				if sort_col == 1 {
					var sort_opt int
					fmt.Println("\n1. Ascending\n2. Descending\n3. Kembali")
					fmt.Print("Masukan : ")
					fmt.Scan(&sort_opt)

					if sort_opt == 1 {
						// Asending
						var sorted_patient patients
						sorted_patient = selectionSort(patient, patients_length, sort_opt, sort_col)
						showData(sorted_patient, patients_length)
					} else if sort_opt == 2 {
						// Descending
						var sorted_patient patients
						sorted_patient = selectionSort(patient, patients_length, sort_opt, sort_col)
						showData(sorted_patient, patients_length)
					} else if sort_opt == 3 {
						fmt.Println("\nOpsi Pengurutan :\n1. Data Terakhir Ditambahkan\n2. Nama\n3. Umur\n4. Email\n5. Kembali")
						fmt.Print("Masukan : ")
						fmt.Scan(&sort_col)
					} else {
						for sort_opt != 1 && sort_opt != 2 && sort_opt != 3 {
							fmt.Println("\nMasukan Tidak Valid !")
							fmt.Println("\n1. Ascending\n2. Descending\n3. Kembali")
							fmt.Print("Masukan : ")
							fmt.Scan(&sort_opt)
						}
					}
				} else if sort_col == 2 {
					var sort_opt int
					fmt.Println("\n1. Ascending\n2. Descending\n3. Kembali")
					fmt.Print("Masukan : ")
					fmt.Scan(&sort_opt)

					if sort_opt == 1 {
						// Asending
						var sorted_patient patients
						sorted_patient = selectionSort(patient, patients_length, sort_opt, sort_col)
						showData(sorted_patient, patients_length)
					} else if sort_opt == 2 {
						// Descending
						var sorted_patient patients
						sorted_patient = selectionSort(patient, patients_length, sort_opt, sort_col)
						showData(sorted_patient, patients_length)
					} else if sort_opt == 3 {
						fmt.Println("\nOpsi Pengurutan :\n1. Data Terakhir Ditambahkan\n2. Nama\n3. Umur\n4. Email\n5. Kembali")
						fmt.Print("Masukan : ")
						fmt.Scan(&sort_col)
					} else {
						for sort_opt != 1 && sort_opt != 2 && sort_opt != 3 {
							fmt.Println("\nMasukan Tidak Valid !")
							fmt.Println("\n1. Ascending\n2. Descending\n3. Kembali")
							fmt.Print("Masukan : ")
							fmt.Scan(&sort_opt)
						}
					}
				} else if sort_col == 3 {
					var sort_opt int
					fmt.Println("\n1. Ascending\n2. Descending\n3. Kembali")
					fmt.Print("Masukan : ")
					fmt.Scan(&sort_opt)

					if sort_opt == 1 {
						// Asending
						var sorted_patient patients
						sorted_patient = selectionSort(patient, patients_length, sort_opt, sort_col)
						showData(sorted_patient, patients_length)
					} else if sort_opt == 2 {
						// Descending
						var sorted_patient patients
						sorted_patient = selectionSort(patient, patients_length, sort_opt, sort_col)
						showData(sorted_patient, patients_length)
					} else if sort_opt == 3 {
						fmt.Println("\nOpsi Pengurutan :\n1. Data Terakhir Ditambahkan\n2. Nama\n3. Umur\n4. Email\n5. Kembali")
						fmt.Print("Masukan : ")
						fmt.Scan(&sort_col)
					} else {
						for sort_opt != 1 && sort_opt != 2 && sort_opt != 3 {
							fmt.Println("\nMasukan Tidak Valid !")
							fmt.Println("\n1. Ascending\n2. Descending\n3. Kembali")
							fmt.Print("Masukan : ")
							fmt.Scan(&sort_opt)
						}
					}
				} else if sort_col == 4 {
					var sort_opt int
					fmt.Println("\n1. Ascending\n2. Descending\n3. Kembali")
					fmt.Print("Masukan : ")
					fmt.Scan(&sort_opt)

					if sort_opt == 1 {
						// Asending
						var sorted_patient patients
						sorted_patient = selectionSort(patient, patients_length, sort_opt, sort_col)
						showData(sorted_patient, patients_length)
					} else if sort_opt == 2 {
						// Descending
						var sorted_patient patients
						sorted_patient = selectionSort(patient, patients_length, sort_opt, sort_col)
						showData(sorted_patient, patients_length)
					} else if sort_opt == 3 {
						fmt.Println("\nOpsi Pengurutan :\n1. Data Terakhir Ditambahkan\n2. Nama\n3. Umur\n4. Email\n5. Kembali")
						fmt.Print("Masukan : ")
						fmt.Scan(&sort_col)
					} else {
						for sort_opt != 1 && sort_opt != 2 && sort_opt != 3 {
							fmt.Println("\nMasukan Tidak Valid !")
							fmt.Println("\n1. Ascending\n2. Descending\n3. Kembali")
							fmt.Print("Masukan : ")
							fmt.Scan(&sort_opt)
						}
					}
				} else if sort_col == 5 {
					var sort_opt int
					fmt.Println("\n1. Ascending\n2. Descending\n3. Kembali")
					fmt.Print("Masukan : ")
					fmt.Scan(&sort_opt)

					if sort_opt == 1 {
						// Asending
						var sorted_patient patients
						sorted_patient = selectionSort(patient, patients_length, sort_opt, sort_col)
						showData(sorted_patient, patients_length)
					} else if sort_opt == 2 {
						// Descending
						var sorted_patient patients
						sorted_patient = selectionSort(patient, patients_length, sort_opt, sort_col)
						showData(sorted_patient, patients_length)
					} else if sort_opt == 3 {
						fmt.Println("\nOpsi Pengurutan :\n1. Data Terakhir Ditambahkan\n2. Nama\n3. Umur\n4. Email\n5. Kembali")
						fmt.Print("Masukan : ")
						fmt.Scan(&sort_col)
					} else {
						for sort_opt != 1 && sort_opt != 2 && sort_opt != 3 {
							fmt.Println("\nMasukan Tidak Valid !")
							fmt.Println("\n1. Ascending\n2. Descending\n3. Kembali")
							fmt.Print("Masukan : ")
							fmt.Scan(&sort_opt)
						}
					}
				} else if sort_col == 6 {
					var sort_opt int
					fmt.Println("\n1. Ascending\n2. Descending\n3. Kembali")
					fmt.Print("Masukan : ")
					fmt.Scan(&sort_opt)

					if sort_opt == 1 {
						// Asending
						var sorted_patient patients
						sorted_patient = selectionSort(patient, patients_length, sort_opt, sort_col)
						showData(sorted_patient, patients_length)
					} else if sort_opt == 2 {
						// Descending
						var sorted_patient patients
						sorted_patient = selectionSort(patient, patients_length, sort_opt, sort_col)
						showData(sorted_patient, patients_length)
					} else if sort_opt == 3 {
						fmt.Println("\nOpsi Pengurutan :\n1. Data Terakhir Ditambahkan\n2. Nama\n3. Umur\n4. Email\n5. Kembali")
						fmt.Print("Masukan : ")
						fmt.Scan(&sort_col)
					} else {
						for sort_opt != 1 && sort_opt != 2 && sort_opt != 3 {
							fmt.Println("\nMasukan Tidak Valid !")
							fmt.Println("\n1. Ascending\n2. Descending\n3. Kembali")
							fmt.Print("Masukan : ")
							fmt.Scan(&sort_opt)
						}
					}
				} else if sort_col == 7 {
					fmt.Println("\nOpsi Penampilan Data : ")
					fmt.Println("1. Tampilkan semua data\n2. Filter Data\n3. Cari Data\n4. Edit Data\n5. Hapus Data\n6. Kembali")
					fmt.Print("Masukan : ")
					fmt.Scan(&show_data_option)
				} else {
					for sort_col != 1 && sort_col != 2 {
						fmt.Println("Masukan tidak valid !")
						fmt.Println("\nOpsi Pengurutan :\n1. Data Terakhir Ditambahkan\n2. Nama\n3. Umur\n4. Email\n5. Nomor Telepon\n6. Bobot\n7. Kembali")
						fmt.Print("Masukan : ")
						fmt.Scan(&sort_col)
					}
				}
			} else if show_data_option == 2 {
				// filter data
				var filtered_option_choose int

				fmt.Println("\nFilter data berdasarkan :")
				fmt.Println("1. Gejala\n2. Kembali")
				fmt.Print("Masukan : ")
				fmt.Scan(&filtered_option_choose)

				if filtered_option_choose == 1 {
					// filter data berdasarkan gejala
					var symptoms_option_choose int
					var filtered_patient patients
					var filtered_patient_length int = 0

					fmt.Println("Tampilkan data pasien dengan gejala : ")
					fmt.Println("1. Kontak Erat\n2. Suspect\n3. Probable\n4. Konfirmasi\n5. Kembali")
					fmt.Print("Masukan : ")
					fmt.Scan(&symptoms_option_choose)

					if symptoms_option_choose == 1 {
						// tampilkan data pasien yang kontak erat
						filter_array(patient, &filtered_patient, &filtered_patient_length, 1)
						fmt.Println("\n==== Daftar Data Pasien Dengan Gejala Kontak Erat ====")
						showData(filtered_patient, filtered_patient_length)
					} else if symptoms_option_choose == 2 {
						// tampilkan data pasien yang suspect
						filter_array(patient, &filtered_patient, &filtered_patient_length, 2)
						fmt.Println("\n==== Daftar Data Pasien Dengan Gejala Suspect ====")
						showData(filtered_patient, filtered_patient_length)
					} else if symptoms_option_choose == 3 {
						// tampilkan data pasien yang probable
						filter_array(patient, &filtered_patient, &filtered_patient_length, 3)
						fmt.Println("\n==== Daftar Data Pasien Dengan Gejala Probable ====")
						showData(filtered_patient, filtered_patient_length)
					} else if symptoms_option_choose == 4 {
						// tampilkan data pasien yang konfirmasi
						filter_array(patient, &filtered_patient, &filtered_patient_length, 4)
						fmt.Println("\n==== Daftar Data Pasien Dengan Gejala Terkonfirmasi Covid ====")
						showData(filtered_patient, filtered_patient_length)
					} else if symptoms_option_choose == 5 {
						fmt.Println("\nFilter data berdasarkan :")
						fmt.Println("1. Gejala\n2. Bobot Prioritas\n3. Kembali")
						fmt.Print("Masukan : ")
						fmt.Scan(&filtered_option_choose)
					} else {
						for symptoms_option_choose != 1 && symptoms_option_choose != 2 && symptoms_option_choose != 3 && symptoms_option_choose != 4 {
							fmt.Println("Masukan tidak valid !")
							fmt.Println("Tampilkan data pasien dengan gejala : ")
							fmt.Println("1. Kontak Erat\n2. Suspect\n3. Probable\n4. Konfirmasi\n5. Kembali")
							fmt.Print("Masukan : ")
							fmt.Scan(&symptoms_option_choose)
						}
					}
				} else if filtered_option_choose == 2 {
					fmt.Println("\nOpsi Penampilan Data : ")
					fmt.Println("1. Tampilkan semua data\n2. Filter Data\n3. Cari Data\n4. Edit Data\n5. Hapus Data\n6. Kembali")
					fmt.Print("Masukan : ")
					fmt.Scan(&show_data_option)
				} else {
					for filtered_option_choose != 1 && filtered_option_choose != 2 {
						fmt.Println("Inputan tidak valid !")
						fmt.Println("1. Gejala\n2. Kembali")
						fmt.Print("Masukan : ")
						fmt.Scan(&filtered_option_choose)
					}
				}
			} else if show_data_option == 3 {
				// Cari data pasien
				var search_option int
				fmt.Println("\nCari Berdasarkan\n1. ID\n2. Nama\n3. Umur\n4. Email\n5.Nomor Telepon\n6. Bobot Prioritas\n7.Kembali")
				fmt.Print("Masukan : ")
				fmt.Scan(&search_option)

				if search_option == 1 {
					// Cari pasien berdasarkan ID
					var left, mid, right, id_found, search_data_patient int

					fmt.Println("\nMasukkan ID Pasien : ")
					fmt.Scan(&search_data_patient)

					// Binary Search
					left = patient[0].id
					right = patient[patients_length-1].id
					id_found = -1

					for left <= right && id_found == -1 {
						mid = (left + right) / 2
						if search_data_patient < mid {
							right = mid - 1
						} else if search_data_patient > mid {
							left = mid + 1
						} else {
							id_found = mid
						}
					}

					if id_found == -1 {
						fmt.Println("\nID Pasien tidak ditemukan")
					} else {
						fmt.Printf("\n==== Daftar Data Pasien dengan ID %d ====\n", id_found)
						var idx = id_found - 1
						fmt.Println("+----+----------------------+----------+----------------------+----------------------+-----------------+---------------+---------------+---------------+-------+")
						fmt.Printf("| %-2s | %-20s | %-8s | %-20s | %-20s | %-15s | %-13s | %-13s | %-13s | %-5s |\n", "ID", "Nama", "Umur", "Email", "No. Telp", "Kontak Erat", "suspect", "Probable", "Konfirmasi", "Bobot")
						fmt.Println("+----+----------------------+----------+----------------------+----------------------+-----------------+---------------+---------------+---------------+-------+")
						fmt.Printf("| %-2d | %-20s | %-8d | %-20s | %-20d |", patient[idx].id, patient[idx].nama, patient[idx].umur, patient[idx].email, patient[idx].no_telp)
						fmt.Printf(" %-15t | %-13t | %-13t | %-13t | %-5d |\n", patient[idx].kontak_erat, patient[idx].suspect, patient[idx].probable, patient[idx].konfirmasi, patient[idx].bobot)
						fmt.Println("+----+----------------------+----------+----------------------+----------------------+-----------------+---------------+---------------+---------------+-------+")
						fmt.Printf("Jumlah Data : %d\n", 1)
					}
				} else if search_option == 2 {
					// Cari pasien berdasarkan Nama
				} else if search_option == 3 {
					// Cari pasien berdasarkan Umur
				} else if search_option == 4 {
					// Cari pasien berdasarkan Email
				} else if search_option == 5 {
					// Cari pasien berdasarkan Nomor Telepon
				} else if search_option == 6 {
					// Cari pasien berdasarkan Bobot Prioritas
				} else if search_option == 7 {
					fmt.Println("\nOpsi Penampilan Data : ")
					fmt.Println("1. Tampilkan semua data\n2. Filter Data\n3. Cari Data\n4. Edit Data\n5. Hapus Data\n6. Kembali")
					fmt.Print("Masukan : ")
					fmt.Scan(&show_data_option)
				} else {
					if search_option != 1 && search_option != 2 && search_option != 3 && search_option != 4 && search_option != 5 && search_option != 6 && search_option != 7 {
						fmt.Println("Masukan tidak valid !")
						fmt.Println("Cari Berdasarkan\n1. ID\n2. Nama\n3. Umur\n4. Email\n5.Nomor Telepon\n6. Bobot Prioritas\n7.Kembali")
						fmt.Print("Masukan : ")
						fmt.Scan(&search_option)
					}
				}
			} else if show_data_option == 4 {
				// Edit data pasien
				var id_patient, idx_patient, itr int
				idx_patient = -1
				itr = 0

				fmt.Println("\nEdit Data Pasien dengan ID : ")
				fmt.Scan(&id_patient)

				// Sequential Search
				for itr < patients_length {
					if patient[itr].id == id_patient {
						idx_patient = itr
					}
					itr++
				}

				if idx_patient != -1 {
					editDataPatient(&patient, patients_length, id_patient-1)
					editSymptomsPatient(&patient, patients_length, id_patient-1)
				} else {
					fmt.Println("ID Pasien tidak ditemukan")
				}
			} else if show_data_option == 5 {
				// Hapus data pasien
				var id_delete int

				fmt.Println("\nInputkan data id pasien yang ingin dihapus : ")
				fmt.Scan(&id_delete)

				if deleteData(&patient, &patients_length, id_delete) == true {
					fmt.Printf("Data pasien id %d berhasil dihapus\n", id_delete)
				} else {
					fmt.Printf("Data pasien id %d gagal dihapus\n", id_delete)
				}
			} else if show_data_option == 6 {
				menuMessage()
				fmt.Scan(&option_choose)
			} else {
				for show_data_option != 1 && show_data_option != 2 && show_data_option != 3 && show_data_option != 4 && show_data_option != 5 {
					fmt.Println("Masukan tidak valid !")
					fmt.Println("\nOpsi Penampilan Data : ")
					fmt.Println("1. Tampilkan semua data\n2. Filter Data\n3. Cari Data\n4. Hapus Data\n5. Kembali")
					fmt.Print("Masukan : ")
					fmt.Scan(&show_data_option)
				}
			}
		} else if option_choose == 2 {
			insertData(&patient, &patients_length)
			menuMessage()
			fmt.Scan(&option_choose)
		} else if option_choose == 3 {
			// return
			os.Exit(0)
		} else {
			for option_choose != 1 && option_choose != 2 && option_choose != 3 {
				fmt.Println("Inputan tidak valid !")
				menuMessage()
				fmt.Scan(&option_choose)
			}
		}
	}
}

func menuMessage() {
	fmt.Println("\nMenu Beranda :")
	fmt.Println("1. Lihat data pasien\n2. Masukkan data pasien\n3. Keluar")
	fmt.Print("Pilih opsi : ")
}

func welcomeMessage() {
	fmt.Println("\n============================")
	fmt.Println("====   Selamat Datang   ====")
	fmt.Println("============================")
}
