package main

import "fmt"

const NMAX int = 50
type ideKonten struct {
	id int
	judul string
	platform  string
	kategori  string
	tanggal   string
	jam       string
	like      int
	komentar  int
	share     int
}
type tabKonten [NMAX]ideKonten

var daftar tabKonten
var nData int = 0
var nextId int = 1

func main() {
	
}

// cari data konten menurut tanggal atau kategori nya
func cariKontenSeq(A tabKonten, N int) {
	var input string
	fmt.Println("Masukan judul/kategori yang diinginkan: ")
	fmt.Scan(&input)
	var found bool = false
	var i int
	for i = 0; i < N; i++{
		if A[i].judul == input || A[i].kategori == input {
			//prosedure tampilkan konten
			found = true
		}
	}
	if !found {
		fmt.Println("Data konten tidak ada")
	}
}

// Mengurutkan data dari paling awal ke akhir, contohnya 01,02,03,....,dst
func selectionSortTgl(A *tabKonten, N int) {
	var i int
	var temp ideKonten
	for i = 0; i < N; i++{
		var min, j int
		min = i
		for j = i+1; j < N; j++ {
			if A[j].tanggal < A[min].tanggal {
				min = j
			}
		}
		temp = A[i]
		A[i] = A[min]
		A[min] = temp
	}
	fmt.Println("Data sudah diurutkan menurut taggalnya")
	// prosedure untuk menampilkan semua data
}

// mMengurutkan engagement dari paling tinggi ke paling rendah
func insertSortEnggmnt(A *tabKonten, N int) {
	var i, j int
	for i = 0; i < N; i++{
		var temp ideKonten
		temp = A[i]
		j = i - 1
		for j >= 0 && totalEngagement(A[j]) < totalEngagement(temp) {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = temp
	}
	fmt.Println("Data sudah diurutkan menurut engagement nya")
}

// Mengembalikan jumlah total engagement
func totalEngagement(egg ideKonten) int{
	return egg.like + egg.komentar + egg.share
}