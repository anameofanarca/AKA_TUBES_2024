package main

import (
	"fmt"
)

// Tipe bentukan untuk menyimpan informasi barang
type Barang struct {
	Kode  string
	Nama  string
	Harga int
	Stok  int
}

// Tipe bentukan untuk array barang
type DaftarBarang [MAX_BARANG]Barang

// Konstanta untuk kapasitas maksimum array
const MAX_BARANG = 100

// Array statis untuk menyimpan data barang
var dataBarang DaftarBarang
var jumlahBarang int

// Fungsi utama
func main() {
	// Inisialisasi jumlah barang dan menambahkan contoh data barang
	jumlahBarang = 0
	tambahBarang("B002", "Shampoo", 12000, 20, &jumlahBarang, &dataBarang)
	tambahBarang("B001", "Sabun", 3000, 50, &jumlahBarang, &dataBarang)
	tambahBarang("B003", "Pasta_Gigi", 8000, 30, &jumlahBarang, &dataBarang)

	// Menampilkan menu
	menu()
}

// Fungsi untuk menampilkan menu
func menu() {
	var pilihan int

	for pilihan != 6 {
		fmt.Println("=========================================================")
		fmt.Println("     Aplikasi Pengelolaan Data Barang di Supermarket")
		fmt.Println("=========================================================")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Hapus Barang")
		fmt.Println("3. Update Barang")
		fmt.Println("4. Cari Barang")
		fmt.Println("5. Tampilkan Barang")
		fmt.Println("6. Keluar")

		fmt.Println("=========================================================")

		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)
		fmt.Println("=========================================================")
		fmt.Println()

		switch pilihan {
		case 1:
			tambahBarangMenu()
		case 2:
			hapusBarangMenu()
		case 3:
			updateBarangMenu()
		case 4:
			cariBarangMenu()
		case 5:
			tampilkanBarangMenu()
		case 6:
			fmt.Println("=========================================================")
			fmt.Println("                   Keluar dari program.                  ")
			fmt.Println("                       Terimakasih                       ")
			fmt.Println("=========================================================")
		default:
			fmt.Println("=========================================================")
			fmt.Println("                   Pilihan tidak valid.                  ")
			fmt.Println("=========================================================")
			fmt.Println()
		}
	}
}

// Fungsi untuk menambah barang
func tambahBarang(kode, nama string, harga, stok int, jumlahBarang *int, dataBarang *DaftarBarang) {
	if *jumlahBarang < MAX_BARANG {
		dataBarang[*jumlahBarang] = Barang{kode, nama, harga, stok}
		*jumlahBarang++
	} else {
		fmt.Println("=========================================================")
		fmt.Println("           Kapasitas penyimpanan barang penuh.           ")
		fmt.Println("=========================================================")
		fmt.Println()
	}
}

func tambahBarangMenu() {
	var kode, nama string
	var harga, stok int

	fmt.Println("=========================================================")
	fmt.Println("                    Menu Tambah Barang                   ")
	fmt.Println("=========================================================")
	fmt.Println("Masukkan kode barang (String kapital dan integer)")
	fmt.Println("Contoh : B001")
	fmt.Print(": ")
	fmt.Scan(&kode)
	fmt.Println()
	fmt.Println("Masukkan nama barang (String)")
	fmt.Println("Gunakan '_' (underscore) sebagai pengganti spasi")
	fmt.Println("Contoh : Mie_Instan")
	fmt.Print(": ")
	fmt.Scan(&nama)
	fmt.Println()
	fmt.Println("Masukkan harga barang (Integer)")
	fmt.Println("Contoh : 16000")
	fmt.Print(": ")
	fmt.Scan(&harga)
	fmt.Println()
	fmt.Println("Masukkan stok barang (Integer)")
	fmt.Println("Contoh : 60")
	fmt.Print(": ")
	fmt.Scan(&stok)
	fmt.Println("=========================================================")
	fmt.Println()

	tambahBarang(kode, nama, harga, stok, &jumlahBarang, &dataBarang)
}

// Fungsi untuk mencari indeks barang berdasarkan kode dan nama (Sequential Search)
func cariBarang(kode, nama string, jumlahBarang int, dataBarang *DaftarBarang) int {
	var i int

	for i = 0; i < jumlahBarang; i++ {
		if dataBarang[i].Kode == kode && dataBarang[i].Nama == nama {
			return i
		}
	}
	return -1
}

// Fungsi untuk mencari indeks barang berdasarkan kode dan nama (Binary Search)
func binarySearchBarang(kode, nama string, jumlahBarang int, dataBarang *DaftarBarang) int {
	var low, high, mid int
	low = 0
	high = jumlahBarang - 1
	selectionSort(kode, true, jumlahBarang, dataBarang)
	for low <= high {
		mid = (low + high) / 2
		if dataBarang[mid].Kode == kode {
			return mid
		} else if dataBarang[mid].Kode < kode {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func cariBarangMenu() {
	var kode, nama string
	var pilih, index int
	fmt.Println("=========================================================")
	fmt.Println("                    Menu Cari Barang                     ")
	fmt.Println("=========================================================")
	fmt.Println("Masukkan kode barang (String kapital dan integer)")
	fmt.Println("Contoh : B001")
	fmt.Print(": ")
	fmt.Scan(&kode)
	fmt.Println()
	fmt.Println("Masukkan nama barang (String)")
	fmt.Println("Gunakan '_' (underscore) sebagai pengganti spasi")
	fmt.Println("Contoh : Mie_Instan")
	fmt.Print(": ")
	fmt.Scan(&nama)
	fmt.Println()
	fmt.Println("Pilih Cara Pencarian")
	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search")
	fmt.Print(": ")
	fmt.Scan(&pilih)
	fmt.Println("=========================================================")

	fmt.Println()
	fmt.Println("=========================================================")
	fmt.Println("                    Status Pencarian                     ")
	fmt.Println("_________________________________________________________")

	switch pilih {
	case 1:
		index = cariBarang(kode, nama, jumlahBarang, &dataBarang)
	case 2:
		index = binarySearchBarang(kode, nama, jumlahBarang, &dataBarang)
	default:
		fmt.Println("                  Pilihan Cara tidak valid.              ")
		fmt.Println("_________________________________________________________")
		fmt.Println("=========================================================")

		return
	}

	if index != -1 {
		fmt.Println("                     Barang ditemukan                    ")
		fmt.Println("_________________________________________________________")
		fmt.Printf("%-10s %-15s %-10s %-10s\n", "Kode", "Nama", "Harga", "Stok")
		fmt.Printf("%-10s %-15s %-10d %-10d\n", dataBarang[index].Kode, dataBarang[index].Nama, dataBarang[index].Harga, dataBarang[index].Stok)
	} else {
		fmt.Println("                  Barang tidak ditemukan.                ")
		fmt.Println("_________________________________________________________")

	}
	fmt.Println("=========================================================")
	fmt.Println()
}

// Fungsi untuk menghapus barang berdasarkan kode dan nama
func hapusBarang(kode, nama string, jumlahBarang *int, dataBarang *DaftarBarang) bool {

	var index, i int
	index = binarySearchBarang(kode, nama, *jumlahBarang, dataBarang)
	if index != -1 {
		for i = index; i < *jumlahBarang-1; i++ {
			dataBarang[i] = dataBarang[i+1]
		}
		*jumlahBarang--
		return true
	}
	return false
}

func hapusBarangMenu() {
	var kode, nama string

	fmt.Println("=========================================================")
	fmt.Println("                    Menu Hapus Barang                    ")
	fmt.Println("=========================================================")
	fmt.Println("Masukkan kode barang (String kapital dan integer)")
	fmt.Println("Contoh : B001")
	fmt.Print(": ")
	fmt.Scan(&kode)
	fmt.Println()
	fmt.Println("Masukkan nama barang (String)")
	fmt.Println("Gunakan '_' (underscore) sebagai pengganti spasi")
	fmt.Println("Contoh : Mie_Instan")
	fmt.Print(": ")
	fmt.Scan(&nama)
	fmt.Println("=========================================================")
	fmt.Println()

	fmt.Println("=========================================================")
	fmt.Println("                    Status Penghapusan                   ")
	fmt.Println("_________________________________________________________")
	if hapusBarang(kode, nama, &jumlahBarang, &dataBarang) {
		fmt.Println("                 Barang berhasil dihapus.                ")
		fmt.Println("_________________________________________________________")

	} else {
		fmt.Println("                  Barang tidak ditemukan.                ")
		fmt.Println("_________________________________________________________")
	}
	fmt.Println("=========================================================")
	fmt.Println()

}

// Fungsi untuk mengupdate barang
func updateBarang(kode1, kode, nama1, nama string, harga, stok int, jumlahBarang int, dataBarang *DaftarBarang) bool {
	var index int

	index = cariBarang(kode1, nama1, jumlahBarang, dataBarang)
	if index != -1 {
		dataBarang[index] = Barang{kode, nama, harga, stok}
		return true
	}
	return false
}

func updateBarangMenu() {
	var kode, kode1, nama1, nama string
	var harga, stok int

	fmt.Println("=========================================================")
	fmt.Println("                    Menu Update Barang                   ")
	fmt.Println("=========================================================")
	fmt.Println("                Barang Yang Akan Di Update               ")
	fmt.Println("_________________________________________________________")
	fmt.Println("Masukkan kode barang (String kapital dan integer)")
	fmt.Println("Contoh : B001")
	fmt.Print(": ")
	fmt.Scan(&kode1)
	fmt.Println()
	fmt.Println("Masukkan nama barang (String)")
	fmt.Println("Gunakan '_' (underscore) sebagai pengganti spasi")
	fmt.Println("Contoh : Mie_Instan")
	fmt.Print(": ")
	fmt.Scan(&nama1)
	fmt.Println()
	fmt.Println("_________________________________________________________")
	fmt.Println("                       Barang Baru                       ")
	fmt.Println("_________________________________________________________")
	fmt.Println("Masukkan kode barang (String kapital dan integer)")
	fmt.Println("Contoh : B001")
	fmt.Print(": ")
	fmt.Scan(&kode)
	fmt.Println()
	fmt.Println("Masukkan nama barang baru (String)")
	fmt.Println("Gunakan '_' (underscore) sebagai pengganti spasi")
	fmt.Println("Contoh : Mie_Instan")
	fmt.Print(": ")
	fmt.Scan(&nama)
	fmt.Println()
	fmt.Println("Masukkan harga barang baru (Integer)")
	fmt.Println("Contoh : 16000")
	fmt.Print(": ")
	fmt.Scan(&harga)
	fmt.Println()
	fmt.Println("Masukkan stok barang baru (Integer)")
	fmt.Println("Contoh : 60")
	fmt.Print(": ")
	fmt.Scan(&stok)
	fmt.Println("_________________________________________________________")
	fmt.Println("=========================================================")
	fmt.Println()
	fmt.Println("=========================================================")
	fmt.Println("                      Status Update                      ")
	fmt.Println("_________________________________________________________")
	if updateBarang(kode1, kode, nama1, nama, harga, stok, jumlahBarang, &dataBarang) {
		fmt.Println("                 Barang berhasil diupdate.               ")
		fmt.Println("_________________________________________________________")

	} else {
		fmt.Println("                  Barang tidak ditemukan.                ")
		fmt.Println("_________________________________________________________")
	}
	fmt.Println("=========================================================")
	fmt.Println()
}

// Fungsi untuk menampilkan semua barang
func tampilkanBarang(jumlahBarang int, dataBarang *DaftarBarang) {
	var i int

	fmt.Println("=========================================================")
	fmt.Println("                      Daftar Barang:                     ")
	fmt.Println("_________________________________________________________")
	fmt.Printf("%-10s %-15s %-10s %-10s\n", "Kode", "Nama", "Harga", "Stok")
	for i = 0; i < jumlahBarang; i++ {
		fmt.Printf("%-10s %-15s %-10d %-10d\n", dataBarang[i].Kode, dataBarang[i].Nama, dataBarang[i].Harga, dataBarang[i].Stok)
	}
	fmt.Println("=========================================================")
	fmt.Println()
}

func tampilkanBarangMenu() {
	var pilihan int
	fmt.Println("=========================================================")
	fmt.Println("                Pilih kategori pengurutan:               ")
	fmt.Println("=========================================================")
	fmt.Println("1. Kode (Ascending)")
	fmt.Println("2. Kode (Descending)")
	fmt.Println("3. Nama (Ascending)")
	fmt.Println("4. Nama (Descending)")
	fmt.Println("5. Harga (Ascending)")
	fmt.Println("6. Harga (Descending)")
	fmt.Println("7. Stok (Ascending)")
	fmt.Println("8. Stok (Descending)")
	fmt.Println("9. Tampilkan tanpa pengurutan")
	fmt.Println("=========================================================")
	fmt.Print("Pilih: ")

	fmt.Scan(&pilihan)

	fmt.Println("=========================================================")
	fmt.Println()

	switch pilihan {
	case 1:
		selectionSort("kode", true, jumlahBarang, &dataBarang)
	case 2:
		selectionSort("kode", false, jumlahBarang, &dataBarang)
	case 3:
		insertionSort("nama", true, jumlahBarang, &dataBarang)
	case 4:
		insertionSort("nama", false, jumlahBarang, &dataBarang)
	case 5:
		selectionSort("harga", true, jumlahBarang, &dataBarang)
	case 6:
		selectionSort("harga", false, jumlahBarang, &dataBarang)
	case 7:
		insertionSort("stok", true, jumlahBarang, &dataBarang)
	case 8:
		insertionSort("stok", false, jumlahBarang, &dataBarang)
	case 9:
		// Do nothing, tampilkan langsung
	default:
		fmt.Println("                   Pilihan tidak valid.                  ")
		return
	}

	tampilkanBarang(jumlahBarang, &dataBarang)
}

// Fungsi Selection Sort
func selectionSort(kategori string, ascending bool, jumlahBarang int, dataBarang *DaftarBarang) {
	var temp Barang
	var i, idx, j int

	for i = 0; i < jumlahBarang-1; i++ {
		idx = i
		for j = i + 1; j < jumlahBarang; j++ {
			if (ascending && compare(dataBarang[j], dataBarang[idx], kategori) < 0) || (!ascending && compare(dataBarang[j], dataBarang[idx], kategori) > 0) {
				idx = j
			}
		}
		temp = dataBarang[i]
		dataBarang[i] = dataBarang[idx]
		dataBarang[idx] = temp
	}
}

// Fungsi Insertion Sort
func insertionSort(kategori string, ascending bool, jumlahBarang int, dataBarang *DaftarBarang) {
	var i, j int
	var key Barang

	for i = 1; i < jumlahBarang; i++ {
		key = dataBarang[i]
		j = i - 1

		for j >= 0 && ((ascending && compare(dataBarang[j], key, kategori) > 0) ||
			(!ascending && compare(dataBarang[j], key, kategori) < 0)) {
			dataBarang[j+1] = dataBarang[j]
			j--
		}
		dataBarang[j+1] = key
	}
}

// Fungsi pembanding untuk sorting
func compare(a, b Barang, kategori string) int {
	switch kategori {
	case "kode":
		return compareStrings(a.Kode, b.Kode)
	case "nama":
		return compareStrings(a.Nama, b.Nama)
	case "harga":
		if a.Harga < b.Harga {
			return -1
		} else if a.Harga > b.Harga {
			return 1
		} else {
			return 0
		}
	case "stok":
		if a.Stok < b.Stok {
			return -1
		} else if a.Stok > b.Stok {
			return 1
		} else {
			return 0
		}
	default:
		return 0
	}
}

// Fungsi pembanding manual untuk string
func compareStrings(a, b string) int {

	var lenA, lenB, minLen int
	var i int

	lenA = len(a)
	lenB = len(b)
	minLen = lenA
	if lenB < minLen {
		minLen = lenB
	}

	for i = 0; i < minLen; i++ {
		if a[i] < b[i] {
			return -1
		} else if a[i] > b[i] {
			return 1
		}
	}

	if lenA < lenB {
		return -1
	} else if lenA > lenB {
		return 1
	} else {
		return 0
	}
}
