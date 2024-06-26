package main

import "fmt"

func patientMenuMessage() {
	fmt.Println("\nMenu Opsi:")
	fmt.Println("1. Urutkan\n2. Filter Data\n3. Cari Pasien\n4. Kembali ke beranda")
	fmt.Print("Pilih opsi : ")
}

func insertData(patient_array *patients, patient_length *int) {
	var i int
	var patient_input string

	if *patient_length < NMAX {
		i = *patient_length

		fmt.Printf("\nMasukkan data untuk pasien ke-%d (Gunakan _ sebagai pengganti spasi)\n", i+1)
		// patient_array[i].id = i + 1
		patient_array[i].id = patient_array[*patient_length-1].id + 1
		patient_array[i].bobot = 0
		fmt.Print("Masukkan Nama : ")
		fmt.Scan(&patient_array[i].nama)
		fmt.Print("Masukkan Umur : ")
		fmt.Scan(&patient_array[i].umur)
		fmt.Print("Masukkan Email : ")
		fmt.Scan(&patient_array[i].email)
		fmt.Print("Masukkan Nomor Telepon : ")
		fmt.Scan(&patient_array[i].no_telp)

		// Pertanyaan 1
		fmt.Println("Apakah anda pernah bertemu dengan orang yang memiliki ciri-ciri seperti ISPA, demam, atau batuk ?")
		fmt.Print("Masukan (Y/N) : ")
		fmt.Scan(&patient_input)
		if patient_input == "Y" {
			patient_array[i].kontak_erat = true
			patient_array[i].bobot += 1
		} else if patient_input == "N" {
			patient_array[i].kontak_erat = false
		} else {
			for patient_input != "Y" && patient_input != "N" {
				fmt.Println("Masukan hanya boleh Y atau N !")
				fmt.Println("Apakah anda pernah bertemu dengan orang yang memiliki ciri-ciri seperti ISPA, demam, atau batuk ?")
				fmt.Print("Masukan (Y/N) : ")
				fmt.Scan(&patient_input)
			}
		}

		// Pertanyaan 2
		fmt.Println("Apakah anda pernah mengalami gejala seperti demam, batuk, atau flu ?")
		fmt.Print("Masukan (Y/N) : ")
		fmt.Scan(&patient_input)
		if patient_input == "N" {
			patient_array[i].probable = false
			patient_array[i].suspect = false
		} else if patient_input == "Y" {
			// Pertanyaan 3
			fmt.Println("Apakah gejala yang anda alami cukup berat ?")
			fmt.Print("Masukan (Y = Berat, N = Ringan) : ")
			fmt.Scan(&patient_input)
			if patient_input == "Y" {
				patient_array[i].probable = true
				patient_array[i].suspect = true
				patient_array[i].bobot += 5
			} else if patient_input == "N" {
				patient_array[i].probable = false
				patient_array[i].suspect = true
				patient_array[i].bobot += 2
			} else {
				for patient_input != "Y" && patient_input != "N" {
					fmt.Println("Masukan hanya boleh Y (gejala berat) atau N (gejala ringan) !")
					fmt.Println("Apakah gejala yang anda alami cukup berat ?")
					fmt.Print("Masukan (Y = Berat, Y = Ringan) : ")
					fmt.Scan(&patient_input)
				}
			}
		} else {
			for patient_input != "Y" && patient_input != "N" {
				fmt.Println("Masukan hanya boleh Y atau N !")
				fmt.Println("Apakah anda pernah mengalami gejala seperti demam, batuk, atau flu ?")
				fmt.Print("Masukan (Y/N) : ")
				fmt.Scan(&patient_input)
			}
		}

		// Pertanyaan 4
		fmt.Println("Apakah anda sudah melakukan tes PCR ?")
		fmt.Print("Masukan (Y/N) : ")
		fmt.Scan(&patient_input)
		if patient_input == "N" {
			patient_array[i].konfirmasi = false
		} else if patient_input == "Y" {
			// Pertanyaan 5
			fmt.Println("Apakah hasil tes PCR positif ?")
			fmt.Print("Masukan (Y/N) : ")
			fmt.Scan(&patient_input)
			if patient_input == "Y" {
				patient_array[i].konfirmasi = true
				patient_array[i].bobot += 7
			} else if patient_input == "N" {
				patient_array[i].konfirmasi = false
			} else {
				for patient_input != "Y" && patient_input != "N" {
					fmt.Println("Masukan hanya boleh Y atau N !")
					fmt.Println("Apakah hasil tes PCR positif ? (Y/N)")
					fmt.Print("Masukan (Y/N) : ")
					fmt.Scan(&patient_input)
				}
			}
		} else {
			for patient_input != "Y" && patient_input != "N" {
				fmt.Println("Masukan hanya boleh Y atau N !")
				fmt.Println("Apakah anda sudah melakukan tes PCR ?")
				fmt.Print("Masukan (Y/N) : ")
				fmt.Scan(&patient_input)
			}
		}

		*patient_length++
	} else {
		fmt.Println("Data sudah memenuhi kapasitas penyimpanan. Hapus data terlebih dahulu untuk menambah data baru.")
	}
}

func showData(patient_array patients, patient_length int) {
	var i int

	if patient_length == 0 {
		fmt.Println("---- Data kosong ----")
	} else {
		fmt.Println("+----+----------------------+----------+----------------------+----------------------+-----------------+---------------+---------------+---------------+-------+")
		fmt.Printf("| %-2s | %-20s | %-8s | %-20s | %-20s | %-15s | %-13s | %-13s | %-13s | %-5s |\n", "ID", "Nama", "Umur", "Email", "No. Telp", "Kontak Erat", "suspect", "Probable", "Konfirmasi", "Bobot")
		fmt.Println("+----+----------------------+----------+----------------------+----------------------+-----------------+---------------+---------------+---------------+-------+")
		for i = 0; i < patient_length; i++ {
			fmt.Printf("| %-2d | %-20s | %-8d | %-20s | %-20d |", patient_array[i].id, patient_array[i].nama, patient_array[i].umur, patient_array[i].email, patient_array[i].no_telp)
			fmt.Printf(" %-15t | %-13t | %-13t | %-13t | %-5d |\n", patient_array[i].kontak_erat, patient_array[i].suspect, patient_array[i].probable, patient_array[i].konfirmasi, patient_array[i].bobot)
		}
		fmt.Println("+----+----------------------+----------+----------------------+----------------------+-----------------+---------------+---------------+---------------+-------+")
	}
	fmt.Printf("Jumlah Data : %d\n", patient_length)
}

func assign_value(filtered_patient_array *patients, patient_array patients, idx, itr int) {
	filtered_patient_array[idx].id = patient_array[itr].id
	filtered_patient_array[idx].nama = patient_array[itr].nama
	filtered_patient_array[idx].no_telp = patient_array[itr].no_telp
	filtered_patient_array[idx].umur = patient_array[itr].umur
	filtered_patient_array[idx].email = patient_array[itr].email
	filtered_patient_array[idx].kontak_erat = patient_array[itr].kontak_erat
	filtered_patient_array[idx].suspect = patient_array[itr].suspect
	filtered_patient_array[idx].probable = patient_array[itr].probable
	filtered_patient_array[idx].konfirmasi = patient_array[itr].konfirmasi
	filtered_patient_array[idx].bobot = patient_array[itr].bobot
}

func filter_array(patient_array patients, filtered_patient_array *patients, filtered_patient_length *int, filteredOption int) {
	var itr int = 0
	var idx int = 0

	// Sequential Search
	for itr < patients_length {
		if filteredOption == 1 && patient[itr].kontak_erat == true {
			assign_value(*&filtered_patient_array, patient_array, idx, itr)
			idx++
		} else if filteredOption == 2 && patient[itr].suspect == true {
			assign_value(*&filtered_patient_array, patient_array, idx, itr)
			idx++
		} else if filteredOption == 3 && patient[itr].probable {
			assign_value(*&filtered_patient_array, patient_array, idx, itr)
			idx++
		} else if filteredOption == 4 && patient[itr].konfirmasi {
			assign_value(*&filtered_patient_array, patient_array, idx, itr)
			idx++
		}
		itr++
	}

	*filtered_patient_length = idx
}

func deleteData(patient_array *patients, patient_length *int, id_patient int) bool {
	var itr, id_found int

	id_found = -1

	// Sequential Search
	for itr = 0; itr < *patient_length; itr++ {
		if patient_array[itr].id == id_patient {
			id_found = id_patient
		}
	}

	// if (id_patient <= 0) || id_patient > patient_array[*patient_length-1].id {
	// 	fmt.Println("Data pasien yang ingin dihapus tidak ada didalam rentang data pasien")
	// 	return false
	// }

	if id_found == -1 {
		fmt.Println("Data pasien tidak ditemukan")
		return false
	} else {
		var newLength, indexLoop int
		indexLoop = 0

		for indexLoop = 0; indexLoop < *patient_length; indexLoop++ {
			if patient_array[indexLoop].id != id_patient {
				assign_value(*&patient_array, *patient_array, newLength, indexLoop)
				newLength++
			}
		}

		if newLength < *patient_length {
			*patient_length = newLength
			return true
		} else {
			return false
		}
	}
}

func editDataPatient(patient_array *patients, patient_length int, idx_patient int) {
	fmt.Printf("\nUbah data pasien %s (Index Data Pasien : %d)\n", patient_array[idx_patient].nama, idx_patient+1)

	fmt.Printf("Masukkan data untuk pasien ke-%d (Gunakan _ sebagai pengganti spasi)\n", idx_patient+1)
	fmt.Print("Masukkan Nama : ")
	fmt.Scan(&patient_array[idx_patient].nama)
	fmt.Print("Masukkan Umur : ")
	fmt.Scan(&patient_array[idx_patient].umur)
	fmt.Print("Masukkan Email : ")
	fmt.Scan(&patient_array[idx_patient].email)
	fmt.Print("Masukkan Nomor Telepon : ")
	fmt.Scan(&patient_array[idx_patient].no_telp)
}

func editSymptomsPatient(patient_array *patients, patient_length int, idx_patient int) {
	var patient_input string

	fmt.Printf("\nPerbarui data gejala pasien %s (Index Data Pasien : %d)\n", patient_array[idx_patient].nama, idx_patient+1)
	patient_array[idx_patient].bobot = 0
	// Pertanyaan 1
	fmt.Println("Apakah anda pernah bertemu dengan orang yang memiliki ciri-ciri seperti ISPA, demam, atau batuk ?")
	fmt.Print("Masukan (Y/N) : ")
	fmt.Scan(&patient_input)
	if patient_input == "Y" {
		patient_array[idx_patient].kontak_erat = true
		patient_array[idx_patient].bobot += 1
	} else if patient_input == "N" {
		patient_array[idx_patient].kontak_erat = false
	} else {
		for patient_input != "Y" && patient_input != "N" {
			fmt.Println("Masukan hanya boleh Y atau N !")
			fmt.Println("Apakah anda pernah bertemu dengan orang yang memiliki ciri-ciri seperti ISPA, demam, atau batuk ?")
			fmt.Print("Masukan (Y/N) : ")
			fmt.Scan(&patient_input)
		}
	}

	// Pertanyaan 2
	fmt.Println("Apakah anda pernah mengalami gejala seperti demam, batuk, atau flu ?")
	fmt.Print("Masukan (Y/N) : ")
	fmt.Scan(&patient_input)
	if patient_input == "N" {
		patient_array[idx_patient].probable = false
		patient_array[idx_patient].suspect = false
	} else if patient_input == "Y" {
		// Pertanyaan 3
		fmt.Println("Apakah gejala yang anda alami cukup berat ?")
		fmt.Print("Masukan (Y = Berat, N = Ringan) : ")
		fmt.Scan(&patient_input)
		if patient_input == "Y" {
			patient_array[idx_patient].probable = true
			patient_array[idx_patient].suspect = true
			patient_array[idx_patient].bobot += 5
		} else if patient_input == "N" {
			patient_array[idx_patient].probable = false
			patient_array[idx_patient].suspect = true
			patient_array[idx_patient].bobot += 2
		} else {
			for patient_input != "Y" && patient_input != "N" {
				fmt.Println("Masukan hanya boleh Y (gejala berat) atau N (gejala ringan) !")
				fmt.Println("Apakah gejala yang anda alami cukup berat ?")
				fmt.Print("Masukan (Y = Berat, Y = Ringan) : ")
				fmt.Scan(&patient_input)
			}
		}
	} else {
		for patient_input != "Y" && patient_input != "N" {
			fmt.Println("Masukan hanya boleh Y atau N !")
			fmt.Println("Apakah anda pernah mengalami gejala seperti demam, batuk, atau flu ?")
			fmt.Print("Masukan (Y/N) : ")
			fmt.Scan(&patient_input)
		}
	}

	// Pertanyaan 4
	fmt.Println("Apakah anda sudah melakukan tes PCR ?")
	fmt.Print("Masukan (Y/N) : ")
	fmt.Scan(&patient_input)
	if patient_input == "N" {
		patient_array[idx_patient].konfirmasi = false
	} else if patient_input == "Y" {
		// Pertanyaan 5
		fmt.Println("Apakah hasil tes PCR positif ?")
		fmt.Print("Masukan (Y/N) : ")
		fmt.Scan(&patient_input)
		if patient_input == "Y" {
			patient_array[idx_patient].konfirmasi = true
			patient_array[idx_patient].bobot += 7
		} else if patient_input == "N" {
			patient_array[idx_patient].konfirmasi = false
		} else {
			for patient_input != "Y" && patient_input != "N" {
				fmt.Println("Masukan hanya boleh Y atau N !")
				fmt.Println("Apakah hasil tes PCR positif ? (Y/N)")
				fmt.Print("Masukan (Y/N) : ")
				fmt.Scan(&patient_input)
			}
		}
	} else {
		for patient_input != "Y" && patient_input != "N" {
			fmt.Println("Masukan hanya boleh Y atau N !")
			fmt.Println("Apakah anda sudah melakukan tes PCR ?")
			fmt.Print("Masukan (Y/N) : ")
			fmt.Scan(&patient_input)
		}
	}
}

func selectionSort(patient_array patients, patient_length, sort_opt, sort_col int) patients {
	var idx, pass, i int
	var temp patient_data

	if sort_col == 1 {
		// Kolom ID
		for pass = 0; pass < patient_length-1; pass++ {
			idx = pass
			if sort_opt == 1 {
				// Ascending
				for i = pass + 1; i < patient_length; i++ {
					if patient_array[idx].id > patient_array[i].id {
						idx = i
					}
				}
				temp = patient_array[idx]
				patient_array[idx] = patient_array[pass]
				patient_array[pass] = temp
			} else if sort_opt == 2 {
				// Descending
				for i = pass + 1; i < patient_length; i++ {
					if patient_array[idx].id < patient_array[i].id {
						idx = i
					}
				}
				temp = patient_array[idx]
				patient_array[idx] = patient_array[pass]
				patient_array[pass] = temp
			}
		}
	} else if sort_col == 2 {
		// Kolom Nama
		for pass = 0; pass < patient_length-1; pass++ {
			idx = pass
			if sort_opt == 1 {
				// Ascending
				for i = pass + 1; i < patient_length; i++ {
					if patient_array[idx].nama > patient_array[i].nama {
						idx = i
					}
				}
				temp = patient_array[idx]
				patient_array[idx] = patient_array[pass]
				patient_array[pass] = temp
			} else if sort_opt == 2 {
				// Descending
				for i = pass + 1; i < patient_length; i++ {
					if patient_array[idx].nama < patient_array[i].nama {
						idx = i
					}
				}
				temp = patient_array[idx]
				patient_array[idx] = patient_array[pass]
				patient_array[pass] = temp
			}
		}
	} else if sort_col == 3 {
		// Kolom Umur
		for pass = 0; pass < patient_length-1; pass++ {
			idx = pass
			if sort_opt == 1 {
				// Ascending
				for i = pass + 1; i < patient_length; i++ {
					if patient_array[idx].umur > patient_array[i].umur {
						idx = i
					}
				}
				temp = patient_array[idx]
				patient_array[idx] = patient_array[pass]
				patient_array[pass] = temp
			} else if sort_opt == 2 {
				// Descending
				for i = pass + 1; i < patient_length; i++ {
					if patient_array[idx].umur < patient_array[i].umur {
						idx = i
					}
				}
				temp = patient_array[idx]
				patient_array[idx] = patient_array[pass]
				patient_array[pass] = temp
			}
		}
	} else if sort_col == 4 {
		// Kolom Email
		for pass = 0; pass < patient_length-1; pass++ {
			idx = pass
			if sort_opt == 1 {
				// Ascending
				for i = pass + 1; i < patient_length; i++ {
					if patient_array[idx].email > patient_array[i].email {
						idx = i
					}
				}
				temp = patient_array[idx]
				patient_array[idx] = patient_array[pass]
				patient_array[pass] = temp
			} else if sort_opt == 2 {
				// Descending
				for i = pass + 1; i < patient_length; i++ {
					if patient_array[idx].email < patient_array[i].email {
						idx = i
					}
				}
				temp = patient_array[idx]
				patient_array[idx] = patient_array[pass]
				patient_array[pass] = temp
			}
		}
	} else if sort_col == 5 {
		// Kolom Nomor Telepon
		for pass = 0; pass < patient_length-1; pass++ {
			idx = pass
			if sort_opt == 1 {
				// Ascending
				for i = pass + 1; i < patient_length; i++ {
					if patient_array[idx].no_telp > patient_array[i].no_telp {
						idx = i
					}
				}
				temp = patient_array[idx]
				patient_array[idx] = patient_array[pass]
				patient_array[pass] = temp
			} else if sort_opt == 2 {
				// Descending
				for i = pass + 1; i < patient_length; i++ {
					if patient_array[idx].no_telp < patient_array[i].no_telp {
						idx = i
					}
				}
				temp = patient_array[idx]
				patient_array[idx] = patient_array[pass]
				patient_array[pass] = temp
			}
		}
	} else if sort_col == 6 {
		// Kolom Bobot
		for pass = 0; pass < patient_length-1; pass++ {
			idx = pass
			if sort_opt == 1 {
				// Ascending
				for i = pass + 1; i < patient_length; i++ {
					if patient_array[idx].bobot > patient_array[i].bobot {
						idx = i
					}
				}
				temp = patient_array[idx]
				patient_array[idx] = patient_array[pass]
				patient_array[pass] = temp
			} else if sort_opt == 2 {
				// Descending
				for i = pass + 1; i < patient_length; i++ {
					if patient_array[idx].bobot < patient_array[i].bobot {
						idx = i
					}
				}
				temp = patient_array[idx]
				patient_array[idx] = patient_array[pass]
				patient_array[pass] = temp
			}
		}
	}
	return patient_array
}
