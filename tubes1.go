package main

import (
	"fmt"
	"strings"
)

// Tipe bentukan untuk menyimpan informasi barang
type Barang struct {
	Kode  string
	Nama  string
	Harga int
	Stok  int
}

// Konstanta untuk kapasitas maksimum array
const MAX_BARANG = 100

// Array statis untuk menyimpan data barang
var dataBarang [MAX_BARANG]Barang
var jumlahBarang int

// Fungsi utama
func main() {
	//jumlahBarang = 0 // Inisialisasi jumlah barang
	//// Menambahkan beberapa data barang untuk contoh
	//tambahBarang("B001", "Sabun", 3000, 50, &jumlahBarang, &dataBarang)
	//tambahBarang("B002", "Shampoo", 12000, 20, &jumlahBarang, &dataBarang)
	//tambahBarang("B003", "Pasta Gigi", 8000, 30, &jumlahBarang, &dataBarang)

	// Menampilkan menu
	menu()
}

// Fungsi untuk menampilkan menu
func menu() {
	/*	I.S. -
		F.S. menampilkan pilihan pengelolaan dan mengarahkan ke pengelolaan barang yang dipilih.
	*/
	var pilihan int

	for pilihan != 6 {
		fmt.Println("\nAplikasi Pengelolaan Data Barang di Supermarket")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Hapus Barang")
		fmt.Println("3. Update Barang")
		fmt.Println("4. Cari Barang")
		fmt.Println("5. Tampilkan Barang")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih menu: ")

		fmt.Scan(&pilihan)

		if pilihan == 1 {
			tambahBarangMenu()
		} else if pilihan == 2 {
			hapusBarangMenu()
		} else if pilihan == 3 {
			updateBarangMenu()
		} else if pilihan == 4 {
			cariBarangMenu()
		} else if pilihan == 5 {
			tampilkanBarangMenu()
		} else if pilihan == 6 {
			fmt.Println("Keluar dari program.")
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// Fungsi untuk menambah barang
func tambahBarang(kode, nama string, harga, stok int, jumlahBarang *int, dataBarang *[MAX_BARANG]Barang) {
	/*	I.S. terdefinisi kode dan nama string, harga dan stok bilangan bulat
		F.S.
	*/
	if *jumlahBarang < MAX_BARANG {
		dataBarang[*jumlahBarang] = Barang{kode, nama, harga, stok}
		*jumlahBarang++
	} else {
		fmt.Println("Kapasitas penyimpanan barang penuh.")
	}
}

func tambahBarangMenu() {
	var kode, nama string
	var harga, stok int

	fmt.Print("Masukkan kode barang: ")
	fmt.Scan(&kode)
	fmt.Print("Masukkan nama barang: ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan harga barang: ")
	fmt.Scan(&harga)
	fmt.Print("Masukkan stok barang: ")
	fmt.Scan(&stok)

	tambahBarang(kode, nama, harga, stok, &jumlahBarang, &dataBarang)
}

// Fungsi untuk mencari indeks barang berdasarkan kode (Sequential Search)
func cariBarang(kode, nama string, jumlahBarang int, dataBarang *[MAX_BARANG]Barang) int {
	for i := 0; i < jumlahBarang; i++ {
		if dataBarang[i].Kode == kode && dataBarang[i].Nama == nama {
			return i
		}
	}
	return -1
}

// Fungsi untuk mencari barang (Binary Search)
func binarySearchBarang(kode string, jumlahBarang int, dataBarang *[MAX_BARANG]Barang) int {
	low, high := 0, jumlahBarang-1
	for low <= high {
		mid := (low + high) / 2
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
	fmt.Print("Masukkan kode barang yang dicari: ")
	fmt.Scan(&kode)
	fmt.Print("Masukkan nama barang yang dicari: ")
	fmt.Scan(&nama)

	index := cariBarang(kode, nama, jumlahBarang, &dataBarang)
	if index != -1 {
		fmt.Printf("Barang ditemukan: %+v\n", dataBarang[index])
	} else {
		fmt.Println("Barang tidak ditemukan.")
	}
}

// Fungsi untuk menghapus barang berdasarkan kode
func hapusBarang(kode, nama string, jumlahBarang *int, dataBarang *[MAX_BARANG]Barang) bool {
	index := cariBarang(kode, nama, *jumlahBarang, dataBarang)
	if index != -1 {
		for i := index; i < *jumlahBarang-1; i++ {
			dataBarang[i] = dataBarang[i+1]
		}
		*jumlahBarang--
		return true
	}
	return false
}

func hapusBarangMenu() {
	var kode, nama string
	fmt.Print("Masukkan kode barang yang akan dihapus: ")
	fmt.Scan(&kode)
	fmt.Print("Masukkan nama barang yang akan dihapus: ")
	fmt.Scan(&nama)

	if hapusBarang(kode, nama, &jumlahBarang, &dataBarang) {
		fmt.Println("Barang berhasil dihapus.")
	} else {
		fmt.Println("Barang tidak ditemukan.")
	}
}

// Fungsi untuk mengupdate barang
func updateBarang(kode, nama1, nama string, harga, stok int, jumlahBarang int, dataBarang *[MAX_BARANG]Barang) bool {
	index := cariBarang(kode, nama1, jumlahBarang, dataBarang)
	if index != -1 {
		dataBarang[index] = Barang{kode, nama, harga, stok}
		return true
	}
	return false
}

func updateBarangMenu() {
	var kode, nama1, nama string
	var harga, stok int

	fmt.Print("Masukkan kode barang yang akan diupdate: ")
	fmt.Scan(&kode)
	fmt.Print("Masukkan nama barang yang akan diupdate: ")
	fmt.Scan(&nama1)
	fmt.Print("Masukkan nama barang baru: ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan harga barang baru: ")
	fmt.Scan(&harga)
	fmt.Print("Masukkan stok barang baru: ")
	fmt.Scan(&stok)

	if updateBarang(kode, nama1, nama, harga, stok, jumlahBarang, &dataBarang) {
		fmt.Println("Barang berhasil diupdate.")
	} else {
		fmt.Println("Barang tidak ditemukan.")
	}
}

// Fungsi untuk menampilkan semua barang
func tampilkanBarang(jumlahBarang int, dataBarang *[MAX_BARANG]Barang) {
	fmt.Println("\nDaftar Barang:")
	for i := 0; i < jumlahBarang; i++ {
		fmt.Println(dataBarang[i].Kode, dataBarang[i].Nama, dataBarang[i].Harga, dataBarang[i].Stok)
	}
}

func tampilkanBarangMenu() {
	var pilihan int
	fmt.Println("\nPilih kategori pengurutan:")
	fmt.Println("1. Kode (Ascending)")
	fmt.Println("2. Kode (Descending)")
	fmt.Println("3. Nama (Ascending)")
	fmt.Println("4. Nama (Descending)")
	fmt.Println("5. Harga (Ascending)")
	fmt.Println("6. Harga (Descending)")
	fmt.Println("7. Stok (Ascending)")
	fmt.Println("8. Stok (Descending)")
	fmt.Print("Pilih: ")

	fmt.Scan(&pilihan)

	if pilihan == 1 {
		selectionSort("kode", true, jumlahBarang, &dataBarang)
	} else if pilihan == 2 {
		selectionSort("kode", false, jumlahBarang, &dataBarang)
	} else if pilihan == 3 {
		insertionSort("nama", true, jumlahBarang, &dataBarang)
	} else if pilihan == 4 {
		insertionSort("nama", false, jumlahBarang, &dataBarang)
	} else if pilihan == 5 {
		selectionSort("harga", true, jumlahBarang, &dataBarang)
	} else if pilihan == 6 {
		selectionSort("harga", false, jumlahBarang, &dataBarang)
	} else if pilihan == 7 {
		insertionSort("stok", true, jumlahBarang, &dataBarang)
	} else if pilihan == 8 {
		insertionSort("stok", false, jumlahBarang, &dataBarang)
	} else {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	tampilkanBarang(jumlahBarang, &dataBarang)
}

// Fungsi Selection Sort
func selectionSort(kategori string, ascending bool, jumlahBarang int, dataBarang *[MAX_BARANG]Barang) {
	for i := 0; i < jumlahBarang-1; i++ {
		idx := i
		for j := i + 1; j < jumlahBarang; j++ {
			if (ascending && compare(dataBarang[j], dataBarang[idx], kategori) < 0) ||
				(!ascending && compare(dataBarang[j], dataBarang[idx], kategori) > 0) {
				idx = j
			}
		}
		dataBarang[i], dataBarang[idx] = dataBarang[idx], dataBarang[i]
	}
}

// Fungsi Insertion Sort
func insertionSort(kategori string, ascending bool, jumlahBarang int, dataBarang *[MAX_BARANG]Barang) {
	for i := 1; i < jumlahBarang; i++ {
		key := dataBarang[i]
		j := i - 1

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
	if kategori == "kode" {
		return strings.Compare(a.Kode, b.Kode)
	} else if kategori == "nama" {
		return strings.Compare(a.Nama, b.Nama)
	} else if kategori == "harga" {
		if a.Harga < b.Harga {
			return -1
		} else if a.Harga > b.Harga {
			return 1
		} else {
			return 0
		}
	} else if kategori == "stok" {
		if a.Stok < b.Stok {
			return -1
		} else if a.Stok > b.Stok {
			return 1
		} else {
			return 0
		}
	}
	return 0
}
