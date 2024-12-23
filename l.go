package main

import "fmt"

type barang struct {
	kode string
	nama string
	harga int
	stok int
}

const max int = 100

type daftarbarang [max]barang

var data barang
var n int

func main(){

}

func tambah(kode, nama string, harga stok int, A *daftarbarang,n *int){
	if *n < max {
		daftarbarang[*n] = barang{kode, nama, harga , stok}
		*n++
	} else {
		fmt.Println("Kapasitas penuh")
	}
}

func cariseq(kode, nama string, A daftarbarang, n int) int {
	var i int
	
	for i=0;i<n;i++{
		if daftarbarang[i].kode == kode && daftarbarang[i].nama == nama {
			return i
		}
	}
	return -1
}

func binnary(kode, nama string, A daftarbarang, n int) int {
	var kanan, kiri, mid int
	kanan = n-1
	kiri = 0
	
	if kanan>=kiri {
		mid = (kanan +kiri)/2
		if A[mid].nama == nama && A[mid].kode == kode{
			return mid
		} else if A[i].kode < kode {
			kiri = mid +1
		} else {
			kanan = mid -1
		}
	}
	
	return -1
}

func hapus(A *daftarbarang, n *int, nama, kode string) bool {
	var index, i int
	
	index = binnary(kode, nama, *A, *n)
	if index!=-1 {
		for i = index;i<*n,i++{
			A[i]= A[i+1]
		}
		i--
		return true
	}
	return false
}

func update(kode1, nama1, kode, nama string, harga,stok int, A daftarbarang, n int) bool{
	var index int
	
	index = cariseq(kode1,nama1,A,n)
	
	if index!=-1{
		daftarbarang[index] = barang{kode, nama, harga,stok}
		return true
	}
	return false
}

