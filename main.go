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
				// Tampilkan semua data
				fmt.Println("\n==== Daftar Data Pasien ====")
				showData(patient, patients_length)
			} else if show_data_option == 2 {
				// filter data
				for show_data_option == 2 {
					var filtered_option_choose int

					fmt.Println("\nFilter data berdasarkan :")
					fmt.Println("1. Gejala\n2. Bobot Prioritas\n3. Kembali")
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
						// Filter data berdasarkan bobot prioritas
					} else if filtered_option_choose == 3 {
						fmt.Println("\nOpsi Penampilan Data : ")
						fmt.Println("1. Tampilkan semua data\n2. Filter Data\n3. Cari Data\n4. Edit Data\n5. Hapus Data\n6. Kembali")
						fmt.Print("Masukan : ")
						fmt.Scan(&show_data_option)
					} else {
						for filtered_option_choose != 1 && filtered_option_choose != 2 && filtered_option_choose != 3 {
							fmt.Println("Inputan tidak valid !")
							fmt.Println("1. Gejala\n2. Bobot Prioritas\n3. Kembali")
							fmt.Print("Masukan : ")
							fmt.Scan(&filtered_option_choose)
						}
					}
				}
			} else if show_data_option == 3 {
				// Cari data pasien
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
			os.Exit(1)
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
