package main

import "fmt"

type ideKonten struct {
	id        int
	judul     string
	platform  string
	kategori  string
	tanggal   string
	jam       string
	like      int
	komentar  int
	share     int
}

const NMAX int = 50

type tabKonten [NMAX]ideKonten

var daftar tabKonten
var nData int = 0
var nextId int = 1

func main() {
	var pilihan int

	for {
		fmt.Println("\n=== APLIKASI MANAJEMEN KONTEN ===")
		fmt.Println("1. Tambah Konten")
		fmt.Println("2. Ubah Konten")
		fmt.Println("3. Hapus Konten")
		fmt.Println("4. Cari Konten (Sequential Search)")
		fmt.Println("5. Cari Konten (Binary Search)")
		fmt.Println("6. Urutkan Berdasarkan Tanggal")
		fmt.Println("7. Urutkan Berdasarkan Engagement")
		fmt.Println("8. Tampilkan Semua Konten")
		fmt.Println("9. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahKonten(&daftar, &nData)
		case 2:
			ubahKonten(&daftar, nData)
		case 3:
			hapusKonten(&daftar, &nData)
		case 4:
			cariKontenSeq(daftar, nData)
		case 5:
			cariKontenBin(daftar, nData)
		case 6:
			selectionSortTgl(&daftar, nData)
		case 7:
			insertSortEngagement(&daftar, nData)
		case 8:
			tampilkanSemua(daftar, nData)
		case 9:
			fmt.Println("Program selesai")
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func tambahKonten(A *tabKonten, N *int) {
	var temp ideKonten
	temp.id = nextId
	
	if *N >= NMAX {
		fmt.Println("Data konten sudah penuh")
		return
	}
	fmt.Print("Judul: ")
	fmt.Scan(&temp.judul)
	fmt.Print("Platform: ")
	fmt.Scan(&temp.platform)
	fmt.Print("Kategori: ")
	fmt.Scan(&temp.kategori)
	fmt.Print("Tanggal (YYYY-MM-DD): ")
	fmt.Scan(&temp.tanggal)
	fmt.Print("Jam (HH:MM): ")
	fmt.Scan(&temp.jam)
	fmt.Print("Like: ")
	fmt.Scan(&temp.like)
	fmt.Print("Komentar: ")
	fmt.Scan(&temp.komentar)
	fmt.Print("Share: ")
	fmt.Scan(&temp.share)

	A[*N] = temp
	*N++
	nextId++

	fmt.Println("Konten berhasil ditambahkan.")
}

func ubahKonten(A *tabKonten, N int) {
	var idUbah int
	fmt.Print("Masukkan ID konten yang ingin diubah: ")
	fmt.Scan(&idUbah)

	for i := 0; i < N; i++ {
		if A[i].id == idUbah {
			fmt.Println("Masukkan data baru:")
			fmt.Print("Judul: ")
			fmt.Scan(&A[i].judul)
			fmt.Print("Platform: ")
			fmt.Scan(&A[i].platform)
			fmt.Print("Kategori: ")
			fmt.Scan(&A[i].kategori)
			fmt.Print("Tanggal (YYYY-MM-DD): ")
			fmt.Scan(&A[i].tanggal)
			fmt.Print("Jam (HH:MM): ")
			fmt.Scan(&A[i].jam)
			fmt.Print("Like: ")
			fmt.Scan(&A[i].like)
			fmt.Print("Komentar: ")
			fmt.Scan(&A[i].komentar)
			fmt.Print("Share: ")
			fmt.Scan(&A[i].share)
			fmt.Println("Data berhasil diubah.")
			return
		}
	}
	fmt.Println("ID tidak ditemukan.")
}

func hapusKonten(A *tabKonten, N *int) {
	var idHapus int
	fmt.Print("Masukkan ID konten yang ingin dihapus: ")
	fmt.Scan(&idHapus)

	for i := 0; i < *N; i++ {
		if A[i].id == idHapus {
			for j := i; j < *N-1; j++ {
				A[j] = A[j+1]
			}
			*N--
			fmt.Println("Konten berhasil dihapus.")
			return
		}
	}
	fmt.Println("ID tidak ditemukan.")
}

func tampilkanSemua(A tabKonten, N int) {
	if N == 0 {
		fmt.Println("Belum ada konten.")
		return
	}
	fmt.Println("ID\tJudul\t\tPlatform\tKategori\tTanggal\t\tJam\tEngagement")
	for i := 0; i < N; i++ {
		total := totalEngagement(A[i])
		fmt.Printf("%d\t%15s\t\t%15s\t\t%15s\t\t%15s\t%15s\t%d\n",
			A[i].id, A[i].judul, A[i].platform, A[i].kategori, A[i].tanggal, A[i].jam, total)
	}
}

//untuk mencari konten menggunakan sequential search
func cariKontenSeq(A tabKonten, N int) {
	var input string
	var i int
	fmt.Print("Masukan judul/kategori yang diinginkan: ")
	fmt.Scan(&input)
	var found bool = false
	for i = 0; i < N; i++ {
		if A[i].judul == input || A[i].kategori == input {
			if !found {
				fmt.Println("Ditemukan:")
			}
			found = true
			fmt.Printf("ID: %10d | Judul: %10s | Kategori: %10s | Engagement: %10d\n",
				A[i].id, A[i].judul, A[i].kategori, totalEngagement(A[i]))
		}
	}
	if !found {
		fmt.Println("Data konten tidak ada.")
	}
}

// Mencari konten menggunakan binary search nya berdasarkan judul yang sudah teurut
func cariKontenBin(A tabKonten, N int) {
	var input string
	fmt.Scan(&input)

	selectionSortJudul(&A, N)

	var right, mid, left int
	var found bool = false

	right = N-1
	left = 0
	
	for left <= right {
		mid = (right+left)/2
		if A[mid].judul == input {
			fmt.Println("Konten ketemu nih")
			fmt.Printf("ID: %d | Judul: %s | Kategori: %s | Engagement: %d\n",
				A[mid].id, A[mid].judul, A[mid].kategori, totalEngagement(A[mid]))
			found = true
			break
		} else if input < A[mid].judul{
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if !found {
		fmt.Println("Konten tidak ada")
	}
}

//prosedure data konten nya terurut, supaya binaryy search berjalan
func selectionSortJudul(A *tabKonten, N int) {
	var i, j, min int
	var temp ideKonten
	for i = 0; i < N-1; i++{
		min = i
		for j = 1+1; j < N; j++{
			if A[j].judul < A[min].judul {
				min = j
			}
		}
		temp = A[i]
		A[i] = A[min]
		A[min] = temp
	}
}

//untuk mengurutkan daftar konten berdasarkan tanggal
func selectionSortTgl(A *tabKonten, N int) {
	var i, j, min int
	var temp ideKonten
	for i = 0; i < N-1; i++ {
		min = i
		for j = i + 1; j < N; j++ {
			if A[j].tanggal < A[min].tanggal {
				min = j
			}
		}
		temp = A[i]
		A[i] = A[min]
		A[min] = temp
	}
	fmt.Println("Data sudah diurutkan menurut tanggal.")
}

//untuk mengurutkan daftar konten berdasrkan tingkat engagement
func insertSortEngagement(A *tabKonten, N int) {
	var i, j int
	var temp ideKonten
	for i = 1; i < N; i++ {
		temp = A[i]
		j = i - 1
		for j >= 0 && totalEngagement(A[j]) < totalEngagement(temp) {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = temp
	}
	fmt.Println("Data sudah diurutkan menurut engagement.")
}

//menghitung total dan mengembalika nilai engagement totalnya
func totalEngagement(egg ideKonten) int {
	return egg.like + egg.komentar + egg.share
}
