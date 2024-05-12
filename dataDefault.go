package main

const NMAX int = 10

type patient_data struct {
	id, no_telp int
	nama, email string
}

type patients [NMAX]patient_data

var patients_length int = 5

func defaultDataPatient(patient_array *patients) {
	patient_array[0].id = 1
	patient_array[0].nama = "Jason Setiawan"
	patient_array[0].email = "jason@gmail.com"
	patient_array[0].no_telp = 62832122376512

	patient_array[1].id = 2
	patient_array[1].nama = "Budi Hartono"
	patient_array[1].email = "budi@gmail.com"
	patient_array[1].no_telp = 6284387874532

	patient_array[2].id = 3
	patient_array[2].nama = "Heri Hartanto"
	patient_array[2].email = "heri@gmail.com"
	patient_array[2].no_telp = 6286511903287

	patient_array[3].id = 4
	patient_array[3].nama = "Sari Meliati"
	patient_array[3].email = "sari@gmail.com"
	patient_array[3].no_telp = 6284388769541

	patient_array[4].id = 5
	patient_array[4].nama = "Alexandra Collin"
	patient_array[4].email = "alex@gmail.com"
	patient_array[4].no_telp = 6285423895678
}
