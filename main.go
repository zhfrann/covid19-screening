package main

import (
	"fmt"
)

var patient patients

func main() {
	var option_choose int
	var option_valid bool = true

	//defaultDataPatient(&patient)   //aktifkan jika ingin menggunakan data default

	fmt.Println("Menu opsi :")
	fmt.Println("1. Lihat data pasien      2. Masukkan data pasien")
	fmt.Print("Pilih opsi : ")
	fmt.Scan(&option_choose)

	for option_valid {
		if option_choose == 1 {
			showData(patient, patients_length)

			fmt.Println("Menu opsi :")
			fmt.Println("1. Lihat data pasien      2. Masukkan data pasien")
			fmt.Print("Pilih opsi : ")
			fmt.Scan(&option_choose)

		} else if option_choose == 2 {
			insertData(&patient, &patients_length)

			fmt.Println("Menu opsi :")
			fmt.Println("1. Lihat data pasien      2. Masukkan data pasien")
			fmt.Print("Pilih opsi : ")
			fmt.Scan(&option_choose)
		} else {
			option_valid = !option_valid
		}
	}

}

func insertData(patient_array *patients, patient_length *int) {
	var i int

	if *patient_length < NMAX {
		i = *patient_length

		fmt.Printf("Masukkan nama, email, nomor telepon untuk data pasien ke-%d\n", i+1)
		patient_array[i].id = i + 1

		fmt.Scan(&patient_array[i].nama, &patient_array[i].email, &patient_array[i].no_telp)
		*patient_length++
	} else {
		fmt.Print("Data sudah memenuhi kapasitas penyimpanan.")
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
		}
	}
}
