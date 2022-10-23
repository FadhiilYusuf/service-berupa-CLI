package main

import (
	"fmt"
	"os"
	"strconv"
)

type people struct {
	absen               int
	nama                string
	alamat              string
	pekerjaan           string
	alasanMemilihGolang string
}

func cetak_data() []people {
	People := []people{
		{1, "Muhammad Fadhiil Yusuf", "Purwakarta", "Software Engineer", "Karena Ingin meningkatkan Skill dan Ilmu yang belum diketahui"},
		{2, "Ahmad Farizi", "Jakarta", "Back End ", "Karena skill ini diminati banyak perusahaan"},
		{3, "Ferdin Faizi", "Bandung", "Front End", "Bahasa yang sedang dipelajari sangat bermanfaat"},
	}

	return People
}

func main() {

	argument := os.Args
	printData(argument)

}

func printData(arguments []string) {

	startIndex, _ := strconv.Atoi(arguments[1])

	for _, value := range cetak_data() {
		if value.absen == startIndex {
			fmt.Println("***************Data People***********************")
			fmt.Println("Nama:", value.nama)
			fmt.Println("Alamat:", value.alamat)
			fmt.Println("Pekerjaan:", value.pekerjaan)
			fmt.Println("Alasan memilih golang:", value.alasanMemilihGolang)
			fmt.Println("*************************************************")
			return

		}
	}

	fmt.Println("Siswa tidak ditemukan")

}
