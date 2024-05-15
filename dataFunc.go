package main

import "fmt"

func insertData(patient_array *patients, patient_length *int) {
	var i int
	var patient_input bool

	if *patient_length < NMAX {
		i = *patient_length

		fmt.Printf("Masukkan nama, email, nomor telepon untuk data pasien ke-%d\n", i+1)
		patient_array[i].id = i + 1

		fmt.Scan(&patient_array[i].nama, &patient_array[i].email, &patient_array[i].no_telp)

		// Pertanyaan 1
		fmt.Println("Apakah anda pernah bertemu dengan orang yang memiliki ciri-ciri seperti ISPA, demam, atau batuk ?")
		fmt.Scan(&patient_input)
		if patient_input {
			patient_array[i].kontak_erat = true
		} else {
			patient_array[i].kontak_erat = false
		}

		// Pertanyaan 2
		fmt.Println("apakah anda pernah mengalami gejala seperti demam, batuk, atau flu atau ringan")
		fmt.Scan(&patient_input)
		if patient_input != true {
			patient_array[i].probable = false
			patient_array[i].suspek = false
		} else {
			// Pertanyaan 3
			fmt.Println("apakah gejala yang anda alami cukup berat atau ringan")
			fmt.Scan(&patient_input)
			if patient_input {
				patient_array[i].probable = true
				patient_array[i].suspek = true
			} else {
				patient_array[i].probable = false
				patient_array[i].suspek = true
			}
		}

		// Pertanyaan 4
		fmt.Println("Apakah anda sudah melakukan tes PCR ?")
		fmt.Scan(&patient_input)
		if patient_input != true {
			patient_array[i].konfirmasi = false
		} else {
			// Pertanyaan 5
			fmt.Println("Apakah hasil tes PCR positif ?")
			fmt.Scan(&patient_input)
			if patient_input {
				patient_array[i].konfirmasi = true
			} else {
				patient_array[i].konfirmasi = false
			}
		}

		*patient_length++
	} else {
		fmt.Println("Data sudah memenuhi kapasitas penyimpanan.")
	}
}

func showData(patient_array patients, patient_length int) {
	var i int

	fmt.Println("Data Pasien :")
	if patient_length == 0 {
		fmt.Println("---- Data kosong ----")
	} else {
		for i = 0; i < patient_length; i++ {
			fmt.Println(patient_array[i].id, patient_array[i].nama, patient_array[i].email, patient_array[i].no_telp)
			fmt.Println("Kontak erat, suspek, probable, konfirmasi, bobot")
			fmt.Println(patient_array[i].kontak_erat, patient_array[i].suspek, patient_array[i].probable, patient_array[i].konfirmasi, patient_array[i].bobot)
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
