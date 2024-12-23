package main

import "fmt"

const NMAX int = 37

type tab [NMAX]int
type Thim struct {
	anggota tab
	panjang int
}

func member(himp Thim, oby int) bool {
	//
	var j bool = false
	var i int = 0
	for i < NMAX && !j {
		if himp.anggota[i] == oby {
			j = true
		}
		i++
	}

	return j
}

func input(himp *Thim) {
	//
	var h int
	himp.panjang = 0

	fmt.Scan(&h)
	for himp.panjang < NMAX && !member(*himp, h) {

		himp.anggota[himp.panjang] = h
		himp.panjang++
		fmt.Scan(&h)
	}
}

func urut(T *Thim) {
	//
	var pass, j, m int

	for pass = 1; pass < T.panjang; pass++ {
		j = pass
		m = T.anggota[j]
		for j > 0 && m < T.anggota[j-1] {
			T.anggota[j] = T.anggota[j-1]
			j--
		}
		T.anggota[j] = m
	}
}

func sama(himp1, himp2 Thim) bool {
	//
	if himp1.panjang != himp2.panjang {
		return false
	} else {
		urut(&himp1)
		urut(&himp2)
		return himp1.anggota == himp2.anggota
	}
}

func cetak(himp Thim) {
	for i := 0; i < himp.panjang; i++ {
		fmt.Print(himp.anggota[i], " ")
	}
	fmt.Println()
}

func main() {
	var set1, set2 Thim
	fmt.Print("Anggota himpunan 1: ")
	input(&set1)
	fmt.Print("Anggota himpunan 2: ")
	input(&set2)
	fmt.Println("Himpunan 1 = Himpunan 2?", sama(set1, set2))
	cetak(set1)
	cetak(set2)
	urut(&set1)
	cetak(set1)
	urut(&set2)
	cetak(set2)
}
