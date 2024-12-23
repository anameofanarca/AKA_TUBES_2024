package main

import "fmt"

const NMAX = 100
const MAXUSERS = 10

type Pengeluaran struct {
	Kategori string
	Jumlah   int
	Tanggal  string
}

type PengeluaranArray [NMAX]Pengeluaran

type User struct {
	Username        string
	Password        string
	PengeluaranList PengeluaranArray
	Count           int
}

var users [MAXUSERS]User
var userCount int

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
)

func main() {
	userCount = 0
	users[userCount] = User{Username: "admin", Password: "admin"}
	userCount++

	for {
		fmt.Println(ColorCyan + "===============================================" + ColorReset)
		fmt.Println(ColorBlue + "                   1. Login" + ColorReset)
		fmt.Println(ColorBlue + "                  2. Register" + ColorReset)
		fmt.Println(ColorBlue + "                   3. Keluar" + ColorReset)
		fmt.Println(ColorCyan + "===============================================" + ColorReset)

		var choice int
		fmt.Print("Pilih menu: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			userIndex := login()
			if userIndex != -1 {
				menu(&users[userIndex])
			}
		case 2:
			register()
		case 3:
			return
		default:
			fmt.Println(ColorRed + "Pilihan tidak valid" + ColorReset)
		}
	}
}

func login() int {
	var username, password string
	fmt.Print("Username: ")
	fmt.Scan(&username)
	fmt.Print("Password: ")
	fmt.Scan(&password)

	for i := 0; i < userCount; i++ {
		if users[i].Username == username && users[i].Password == password {
			fmt.Println(ColorGreen + "Login berhasil!" + ColorReset)
			return i
		}
	}

	fmt.Println(ColorRed + "Username atau password salah" + ColorReset)
	return -1
}

func register() {
	if userCount >= MAXUSERS {
		fmt.Println(ColorRed + "Maksimal jumlah pengguna tercapai!" + ColorReset)
		return
	}

	var username, password string
	fmt.Print("Masukkan username baru: ")
	fmt.Scan(&username)
	fmt.Print("Masukkan password baru: ")
	fmt.Scan(&password)

	users[userCount] = User{Username: username, Password: password}
	userCount++
	fmt.Println(ColorGreen + "Registrasi berhasil!" + ColorReset)
}

func menu(user *User) {
	for {
		fmt.Println(ColorGreen + "===============================================" + ColorReset)
		fmt.Println(ColorCyan + "|               MONEY ORGANIZER               |" + ColorReset)
		fmt.Println(ColorCyan + "|                version 0.0.1                |" + ColorReset)
		fmt.Println(ColorGreen + "|=============================================|" + ColorReset)
		fmt.Println(ColorCyan + "| 1. Tambah Pengeluaran                       |" + ColorReset)
		fmt.Println(ColorCyan + "| 2. Tampilkan Pengeluaran                    |" + ColorReset)
		fmt.Println(ColorCyan + "| 3. Cari Pengeluaran (Sequential Search)     |" + ColorReset)
		fmt.Println(ColorCyan + "| 4. Cari Pengeluaran (Binary Search)         |" + ColorReset)
		fmt.Println(ColorCyan + "| 5. Edit Pengeluaran                         |" + ColorReset)
		fmt.Println(ColorCyan + "| 6. Hapus Pengeluaran                        |" + ColorReset)
		fmt.Println(ColorCyan + "| 7. Urutkan Pengeluaran (Selection Sort)     |" + ColorReset)
		fmt.Println(ColorCyan + "| 8. Urutkan Pengeluaran (Insertion Sort)     |" + ColorReset)
		fmt.Println(ColorCyan + "| 9. Keluar                                   |" + ColorReset)
		fmt.Println(ColorGreen + "===============================================" + ColorReset)

		var choice int
		fmt.Print("Pilih menu: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			user.Count = tambahPengeluaran(&user.PengeluaranList, user.Count)
		case 2:
			tampilkanPengeluaran(user.PengeluaranList, user.Count)
		case 3:
			cariPengeluaranSequential(user.PengeluaranList, user.Count)
		case 4:
			cariPengeluaranBinary(user.PengeluaranList, user.Count)
		case 5:
			user.Count = editPengeluaran(&user.PengeluaranList, user.Count)
		case 6:
			user.Count = hapusPengeluaran(&user.PengeluaranList, user.Count)
		case 7:
			urutkanPengeluaranSelection(&user.PengeluaranList, user.Count)
		case 8:
			urutkanPengeluaranInsertion(&user.PengeluaranList, user.Count)
		case 9:
			return
		default:
			fmt.Println(ColorRed + "Pilihan tidak valid" + ColorReset)
		}
	}
}

func tambahPengeluaran(arr *PengeluaranArray, n int) int {
	if n >= NMAX {
		fmt.Println(ColorRed + "Data pengeluaran penuh!" + ColorReset)
		return n
	}
	var pengeluaran Pengeluaran
	fmt.Print("Masukkan kategori: ")
	fmt.Scan(&pengeluaran.Kategori)
	fmt.Print("Masukkan jumlah: ")
	fmt.Scan(&pengeluaran.Jumlah)
	fmt.Print("Masukkan tanggal (YYYY-MM-DD): ")
	fmt.Scan(&pengeluaran.Tanggal)

	arr[n] = pengeluaran
	n++
	return n
}

func tampilkanPengeluaran(arr PengeluaranArray, n int) {
	fmt.Println("Data Pengeluaran:")
	for i := 0; i < n; i++ {
		fmt.Printf("%d. Kategori: %s, Jumlah: %d, Tanggal: %s\n", i+1, arr[i].Kategori, arr[i].Jumlah, arr[i].Tanggal)
	}
	fmt.Println("Tekan 'Enter' untuk melanjutkan...")
	fmt.Scanln()
	fmt.Scanln()
}

func cariPengeluaranSequential(arr PengeluaranArray, n int) {
	var kategori string
	fmt.Print("Masukkan kategori yang dicari: ")
	fmt.Scan(&kategori)
	found := false
	for i := 0; i < n; i++ {
		if arr[i].Kategori == kategori {
			fmt.Printf("Ditemukan: Kategori: %s, Jumlah: %d, Tanggal: %s\n", arr[i].Kategori, arr[i].Jumlah, arr[i].Tanggal)
			found = true
		}
	}
	if !found {
		fmt.Println(ColorRed + "Pengeluaran tidak ditemukan" + ColorReset)
	}
	fmt.Println("Tekan 'Enter' untuk melanjutkan...")
	fmt.Scanln()
	fmt.Scanln()
}

func cariPengeluaranBinary(arr PengeluaranArray, n int) {
	urutkanPengeluaranInsertion(&arr, n)
	var kategori string
	fmt.Print("Masukkan kategori yang dicari: ")
	fmt.Scan(&kategori)
	left, right := 0, n-1
	found := false
	for left <= right {
		mid := (left + right) / 2
		if arr[mid].Kategori == kategori {
			fmt.Printf("Ditemukan: Kategori: %s, Jumlah: %d, Tanggal: %s\n", arr[mid].Kategori, arr[mid].Jumlah, arr[mid].Tanggal)
			found = true
			break
		} else if arr[mid].Kategori < kategori {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if !found {
		fmt.Println(ColorRed + "Pengeluaran tidak ditemukan" + ColorReset)
	}
	fmt.Println("Tekan 'Enter' untuk melanjutkan...")
	fmt.Scanln()
	fmt.Scanln()
}

func editPengeluaran(arr *PengeluaranArray, n int) int {
	var index int
	fmt.Print("Masukkan nomor pengeluaran yang ingin diubah: ")
	fmt.Scan(&index)
	if index <= 0 || index > n {
		fmt.Println(ColorRed + "Nomor pengeluaran tidak valid" + ColorReset)
		return n
	}
	index--
	var pengeluaran Pengeluaran
	fmt.Print("Masukkan kategori baru: ")
	fmt.Scan(&pengeluaran.Kategori)
	fmt.Print("Masukkan jumlah baru: ")
	fmt.Scan(&pengeluaran.Jumlah)
	fmt.Print("Masukkan tanggal baru (YYYY-MM-DD): ")
	fmt.Scan(&pengeluaran.Tanggal)

	arr[index] = pengeluaran
	return n
}

func hapusPengeluaran(arr *PengeluaranArray, n int) int {
	var index int
	fmt.Print("Masukkan nomor pengeluaran yang ingin dihapus: ")
	fmt.Scan(&index)
	if index <= 0 || index > n {
		fmt.Println(ColorRed + "Nomor pengeluaran tidak valid" + ColorReset)
		return n
	}
	index--
	for i := index; i < n-1; i++ {
		arr[i] = arr[i+1]
	}
	n--
	return n
}

func urutkanPengeluaranSelection(arr *PengeluaranArray, n int) {
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j].Kategori < arr[minIdx].Kategori {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
	fmt.Println(ColorGreen + "Pengeluaran telah diurutkan dengan Selection Sort!" + ColorReset)
}

func urutkanPengeluaranInsertion(arr *PengeluaranArray, n int) {
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j].Kategori > key.Kategori {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	fmt.Println(ColorGreen + "Pengeluaran telah diurutkan dengan Insertion Sort!" + ColorReset)
}
