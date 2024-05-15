package main

const NMAX int = 10

type patient_data struct {
	id, no_telp, umur, bobot                  int
	nama, email                               string
	probable, suspek, konfirmasi, kontak_erat bool
}

type patients [NMAX]patient_data

func defaultDataPatient(patient_array *patients) {
	patient_array[0].id = 1
	patient_array[0].nama = "Jason_Setiawan"
	patient_array[0].email = "jason@gmail.com"
	patient_array[0].no_telp = 62832122376512
	patient_array[0].umur = 21
	patient_array[0].kontak_erat = true
	patient_array[0].suspek = true
	patient_array[0].probable = false
	patient_array[0].konfirmasi = false
	patient_array[0].bobot = 3

	patient_array[1].id = 2
	patient_array[1].nama = "Budi_Hartono"
	patient_array[1].email = "budi@gmail.com"
	patient_array[1].no_telp = 6284387874532
	patient_array[1].umur = 20
	patient_array[1].kontak_erat = true
	patient_array[1].suspek = true
	patient_array[1].probable = true
	patient_array[1].konfirmasi = false
	patient_array[1].bobot = 6

	patient_array[2].id = 3
	patient_array[2].nama = "Heri_Hartanto"
	patient_array[2].email = "heri@gmail.com"
	patient_array[2].no_telp = 6286511903287
	patient_array[2].umur = 25
	patient_array[2].kontak_erat = true
	patient_array[2].suspek = true
	patient_array[2].probable = false
	patient_array[2].konfirmasi = false
	patient_array[2].bobot = 3

	patient_array[3].id = 4
	patient_array[3].nama = "Sari_Meliati"
	patient_array[3].email = "sari@gmail.com"
	patient_array[3].no_telp = 6284388769541
	patient_array[3].umur = 22
	patient_array[3].kontak_erat = true
	patient_array[3].suspek = false
	patient_array[3].probable = false
	patient_array[3].konfirmasi = false
	patient_array[3].bobot = 1

	patient_array[4].id = 5
	patient_array[4].nama = "Alexandra_Collin"
	patient_array[4].email = "alex@gmail.com"
	patient_array[4].no_telp = 6285423895678
	patient_array[4].umur = 31
	patient_array[4].kontak_erat = false
	patient_array[4].suspek = false
	patient_array[4].probable = false
	patient_array[4].konfirmasi = true
	patient_array[4].bobot = 7

	patients_length = 5
}
