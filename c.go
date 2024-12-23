//SOAL 1
/*
Program uang_jajan
kamus
	type uang : array [0..99] of integer
	i : integer
	n,x,max : integer (n=jum.max ; x=uang jajan ke i ; max=nilai terbesar)
	daftar_uang : uang
Algoritma
	input(n)
	for i <- 0 to n-1 do
		input(x)
		daftar_uang[i] = x
	endfor
	max = daftar_uang[0]
	for i <- 1 to n-1 do
		if max < daftar_uang[i] then
			max <- daftar_uang[i]
		endif
	endfor
	output(max)
endprogram
*/

/*ATAU*/

/*
Program uang_jajan
kamus
	type uang : array [0..99] of integer
	i : integer
	n,max : integer (n=jum.max ; max=nilai terbesar)
	daftar_uang : uang
algoritma
	input(n)
	for i <- 0 to n-1 do
		input(daftar_uang[i])
	endfor
	max = daftar_uang[0]
	for i <- 1 to n-1 do
		if max < daftar_uang[i] then
			max <- daftar_uang[i]
		endif
	endfor
	output(max)
endprogram
*/

/*
var terbesar int
terbesar <- daftar_uang
for i<=n do
	input(input(daftar_uang[i])
	if terbesar < daftar_uang then
		terbesar <- daftar_uang
	endif
endfor
*/

//SOAL 2
/*
function jarak(p1,p2 : titik) -> real
{diterima nilai sumbu dari titik p1 dan p2, untuk mengembalikan jarak dari p1 dan p2}
algoritma
return akar(((p1.x-p2.x)*(p1.x-p2.x)) + ((p1.y-p2.y)*(p1.y-p2.y)))
endfunction

function akar(x : real)->real
{diterima nilai suatu x untuk mengembalikan akar kuadrat }
endfunction

program jarak_titik
kamus
	type titik <
	x,y : real
	>
jarak : real
a,b,c,d,p1,p2 : integer  || p1,p2 : titik

algoritma
input(a,b,c,d) 			|| input(p1.x,p1.y)
p1 = titik{a,b}			|| input(p2.x,p2.y)
p2 = titik{c,d}			||
output(jarak(p1,p2))	|| output(jarak(p1,p2))
endprogram
*/

//SOAL 3 Palindrom

/*
procedure isi(in n int, in/out data : bil)
{I.S. terdefinisi bil.bul n
 F.S. array data berisi sejumlah n bil.bulat
}
kamus
	i : integer
algoritma
	for i=0 to n-1 do
		input(data[i])
	endfor
endprocedure

procedure reverse(in n : integer,arr1 : bill, in/out arr2 : bill)
{

}
kamus									kamus
	i : integer							i,j : integer
algoritma								algoritma
	for i <- 0 to n-1 do					j = n-1
		arr2[(n-1)-i] <- arr1[i]			for i=0 to n-1 do
		z <- z-1								arr2[j] = arr1[i]
	endfor										j = j-1
endprocedure								endfor

function cek(arr : bil,n : integer)-> boolean
{mengembalikan true apabila array arr memiliki pola palindrom, false untuk kondisilainnya}
kamus
	arrV = bill
	i = integer
	cek= boolean
algoritma1
	cek = true
	reverse(arr,n,arrV)
	for i=0 to n-1 do
		cek = cek and arr[i]==arrV[i]
	endfor
	return cek
algoritma2
	reverse(arr,n,arrV)
	cek = true
	i = 0
	while i<=n and cek do
		cek = arr[i]==arrV[i]
		i = i+1
	endwhile
	return cek
algoritma3(tanpa reverse|pakai rekursive)

endfunction

program palindrom
kamus
	type bil:array[0..256] of integer

endprogram


type
*/

// Pertemuan 16 April
/*

Type mahasiswa <
	Nama, nim : string
	Eprt, ipk : real
Semester : integer
>

Constant N : integer = 1000

Type arrmhs : array[0..N-1] of mahasiswa

Procedure mengisiarray(in/out t arrmhs, m integer)
{I.S. -
 F.S. array t berisi m data mahasiswa yang berasal dari masukan}
}
Algoritma
m = 0
input(nim)
while nim!="nome" and m<=N-1 then
	input(t[m].nama,t[m].eprt,t[m].ipk,t[m].semester)
	t[m].nim = nim
	m = m+1
	input(nim)
endwhile
endprocedure

function eprttinggi()->

*/

