package main

import (
	"fmt"
)

const NMAX int = 100

type pengeluaran struct {
	kategori        string
	tanggal, jumlah int
}

type arrPengeluaran [NMAX]pengeluaran

func bacaData(A *arrPengeluaran, n *int) {
	fmt.Scan(&*n)
	if *n > NMAX {
		*n = NMAX
	}
	for i := 0; i < *n; i++ {
		fmt.Scan(&A[i].tanggal, &A[i].kategori, &A[i].jumlah)
	}
}

func cetakData(A arrPengeluaran, n int) {
	fmt.Printf("%10s %10s %10s\n", "Tanggal", "Kategori", "Jumlah")
	for i := 0; i < n; i++ {
		fmt.Printf("%10d %10s %10d\n", A[i].tanggal, A[i].kategori, A[i].jumlah)
	}
}

func tambahData(A *arrPengeluaran, n *int) {
	var j, u int
	fmt.Scan(&j)
	u = j + *n
	if u > NMAX {
		u = NMAX
	}
	for i := *n; i < u; i++ {
		fmt.Scan(&A[i].tanggal, &A[i].kategori, &A[i].jumlah)
	}
	*n = u
}

func hitungTotal(A arrPengeluaran, n int) int {
	var total int
	for i := 0; i < n; i++ {
		total += A[i].jumlah
	}
	return total
}

func hitungTotalPerBulan(A arrPengeluaran, n, bulan int) int {
	var total int
	for i := 0; i < n; i++ {
		if A[i].tanggal/100 == bulan {
			total += A[i].jumlah
		}
	}
	return total
}

func cariPengeluaran(A arrPengeluaran, n int, tanggal int, kategori string) int {
	for i := 0; i < n; i++ {
		if tanggal == A[i].tanggal && kategori == A[i].kategori {
			return i
		}
	}
	return -1
}

func hapusPengeluaran(A *arrPengeluaran, n *int) {
	var tgl, index int
	var ktg string
	fmt.Scan(&tgl, &ktg)
	index = cariPengeluaran(*A, *n, tgl, ktg)
	if index != -1 {
		for i := index; i < *n-1; i++ {
			A[i] = A[i+1]
		}
		*n--
		fmt.Println("Data berhasil dihapus")
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func editPengeluaran(A *arrPengeluaran, n int) {
	var tgl, index int
	var ktg string
	fmt.Scan(&tgl, &ktg)
	index = cariPengeluaran(*A, n, tgl, ktg)
	if index != -1 {
		fmt.Println("Masukkan data baru : ")
		fmt.Scan(&A[index].tanggal, &A[index].kategori, &A[index].jumlah)
		fmt.Println("Data berhasil diedit")
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func pengelompokkanPengeluaran(A arrPengeluaran, n int) {
	var kategori string
	fmt.Scan(&kategori)
	fmt.Printf("%10s %10s %10s\n", "Tanggal", "Kategori", "Jumlah")
	for i := 0; i < n; i++ {
		if A[i].kategori == kategori {
			fmt.Printf("%10d %10s %10d\n", A[i].tanggal, A[i].kategori, A[i].jumlah)
		}
	}
}

func selectionSortTgl(A *arrPengeluaran, n int) {
	//ascending
	for pass := 0; pass < n-1; pass++ {
		idx := pass
		for j := pass + 1; j < n; j++ {
			if A[j].tanggal < A[idx].tanggal {
				idx = j
			}
		}
		A[pass], A[idx] = A[idx], A[pass]
	}
}

func selectionSortKtg(A *arrPengeluaran, n int) {
	//ascending/
	for pass := 0; pass < n-1; pass++ {
		idx := pass
		for j := pass + 1; j < n; j++ {
			if A[j].kategori < A[idx].kategori {
				idx = j
			}
		}
		A[pass], A[idx] = A[idx], A[pass]
	}
}

func insertionSortTgl(A *arrPengeluaran, n int) {
	//descending/
	for pass := 1; pass < n; pass++ {
		temp := A[pass]
		j := pass - 1
		for j >= 0 && A[j].tanggal < temp.tanggal {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = temp
	}
}

func insertionSortKtg(A *arrPengeluaran, n int) {
	//descending/
	for pass := 1; pass < n; pass++ {
		temp := A[pass]
		j := pass - 1
		for j >= 0 && A[j].kategori < temp.kategori {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = temp
	}
}

func main() {
	var pilih, n int
	var arr arrPengeluaran

	for {
		fmt.Println("SISTEM PENCATATAN PENGELUARAN RUMAH TANGGA")
		fmt.Println("=========================")
		fmt.Println("           MENU          ")
		fmt.Println("=========================")
		fmt.Println("1. Tambah Pengeluaran")
		fmt.Println("2. Hapus Pengeluaran")
		fmt.Println("3. Edit Pengeluaran")
		fmt.Println("4. Hitung Total Pengeluaran Per Bulan")
		fmt.Println("5. Pengelompokkan Berdasarkan Kategori")
		fmt.Println("6. Urutkan Pengeluaran Berdasarkan Tanggal")
		fmt.Println("7. Urutkan Pengeluaran Berdasarkan Kategori")
		fmt.Println("8. Keluar")
		fmt.Println("=========================")
		fmt.Print("Pilih Menu : ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			tambahData(&arr, &n)
			cetakData(arr, n)
		case 2:
			hapusPengeluaran(&arr, &n)
			cetakData(arr, n)
		case 3:
			editPengeluaran(&arr, n)
			cetakData(arr, n)
		case 4:
			var bulan int
			fmt.Print("Masukkan bulan (YYYYMM): ")
			fmt.Scan(&bulan)
			total := hitungTotalPerBulan(arr, n, bulan)
			fmt.Printf("Total pengeluaran untuk bulan %d adalah: %d\n", bulan, total)
		case 5:
			pengelompokkanPengeluaran(arr, n)
		case 6:
			var input int
			fmt.Println("1. Ascending")
			fmt.Println("2. Descending")
			fmt.Scan(&input)
			if input == 1 {
				selectionSortTgl(&arr, n)
			} else if input == 2 {
				insertionSortTgl(&arr, n)
			}
			cetakData(arr, n)
		case 7:
			var input int
			fmt.Println("1. Ascending")
			fmt.Println("2. Descending")
			fmt.Scan(&input)
			if input == 1 {
				selectionSortTgl(&arr, n)
			} else if input == 2 {
				insertionSortTgl(&arr, n)
			}
			cetakData(arr, n)
		case 8:
			fmt.Println("Program selesai")
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}
