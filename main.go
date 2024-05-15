package main

import (
	"fmt"
)

var patient patients

func main() {
	var option_choose int
	var option_valid bool = true

	defaultDataPatient(&patient) //aktifkan jika ingin menggunakan data default

	menuMessage()
	fmt.Scan(&option_choose)

	for option_valid {
		if option_choose == 1 {
			showData(patient, patients_length)

			menuMessage()
			fmt.Scan(&option_choose)

		} else if option_choose == 2 {
			insertData(&patient, &patients_length)

			menuMessage()
			fmt.Scan(&option_choose)
		} else {
			option_valid = !option_valid
		}
	}
}

func menuMessage() {
	fmt.Println("============================")
	fmt.Println("====   Selamat Datang   ====")
	fmt.Println("============================\n")

	fmt.Println("Menu opsi :")
	fmt.Println("1. Lihat data pasien\n2. Masukkan data pasien\n")
	fmt.Print("Pilih opsi : ")
}
