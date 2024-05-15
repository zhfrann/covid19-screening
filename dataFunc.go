package main

import "fmt"

func orderMenuMessage() {
	fmt.Println("Urutkan data dengan tingkat prioritas:")
	fmt.Println("1. Tertinggi\n2.Terendah\n")
	fmt.Print("Pilih opsi : ")
}

func insertData(patient_array *patients, patient_length *int) {
	var i int
	var patient_input string

	if *patient_length < NMAX {
		i = *patient_length

		fmt.Printf("\nMasukkan data untuk pasien ke-%d (Gunakan _ sebagai pengganti spasi)\n", i+1)
		patient_array[i].id = i + 1
		fmt.Print("Masukkan Nama : ")
		fmt.Scan(&patient_array[i].nama)
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
			patient_array[i].suspek = false
		} else if patient_input == "Y" {
			// Pertanyaan 3
			fmt.Println("Apakah gejala yang anda alami cukup berat ?")
			fmt.Print("Masukan (Y = Berat, N = Ringan) : ")
			fmt.Scan(&patient_input)
			if patient_input == "Y" {
				patient_array[i].probable = true
				patient_array[i].suspek = true
				patient_array[i].bobot += 5
			} else if patient_input == "N" {
				patient_array[i].probable = false
				patient_array[i].suspek = true
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

	fmt.Println("\n==== Daftar Data Pasien ====")
	if patient_length == 0 {
		fmt.Println("---- Data kosong ----")
	} else {
		fmt.Printf("%2s %20s %20s %20s %15s %13s %13s %13s %8s\n", "ID", "Nama", "Email", "No. Telp", "Kontak Erat", "Suspek", "Probable", "Konfirmasi", "Bobot")
		for i = 0; i < patient_length; i++ {
			fmt.Printf("%2d %20s %20s %20d", patient_array[i].id, patient_array[i].nama, patient_array[i].email, patient_array[i].no_telp)
			fmt.Printf("%16t %13t %13t %13t %8d\n", patient_array[i].kontak_erat, patient_array[i].suspek, patient_array[i].probable, patient_array[i].konfirmasi, patient_array[i].bobot)
		}
	}
}

func editDataPatient(patient_array *patients, patient_length *int, idx_patient int) {
	fmt.Printf("\nUbah data pasien %s (Index Data Pasien : %d)\n", patient_array[idx_patient].nama, idx_patient+1)

	fmt.Printf("Masukkan nama, email, nomor telepon untuk data pasien ke-%d\n", idx_patient+1)
	fmt.Scan(&patient_array[idx_patient].nama, &patient_array[idx_patient].email, &patient_array[idx_patient].no_telp)
}

func editSymptomsPatient(patient_array *patients, patient_length *int, idx_patient int) {
	var patient_input bool

	fmt.Printf("\nPerbarui data gejala pasien %s (Index Data Pasien : %d)\n", patient_array[idx_patient].nama, idx_patient+1)

	// Pertanyaan 1
	fmt.Println("Apakah anda pernah bertemu dengan orang yang memiliki ciri-ciri seperti ISPA, demam, atau batuk ?")
	fmt.Scan(&patient_input)
	if patient_input {
		patient_array[idx_patient].kontak_erat = true
	} else {
		patient_array[idx_patient].kontak_erat = false
	}

	// Pertanyaan 2
	fmt.Println("apakah anda pernah mengalami gejala seperti demam, batuk, atau flu atau ringan")
	fmt.Scan(&patient_input)
	if patient_input != true {
		patient_array[idx_patient].probable = false
		patient_array[idx_patient].suspek = false
	} else {
		// Pertanyaan 3
		fmt.Println("apakah gejala yang anda alami cukup berat atau ringan")
		fmt.Scan(&patient_input)
		if patient_input {
			patient_array[idx_patient].probable = true
			patient_array[idx_patient].suspek = true
		} else {
			patient_array[idx_patient].probable = false
			patient_array[idx_patient].suspek = true
		}
	}

	// Pertanyaan 4
	fmt.Println("Apakah anda sudah melakukan tes PCR ?")
	fmt.Scan(&patient_input)
	if patient_input != true {
		patient_array[idx_patient].konfirmasi = false
	} else {
		// Pertanyaan 5
		fmt.Println("Apakah hasil tes PCR positif ?")
		fmt.Scan(&patient_input)
		if patient_input {
			patient_array[idx_patient].konfirmasi = true
		} else {
			patient_array[idx_patient].konfirmasi = false
		}
	}
}
