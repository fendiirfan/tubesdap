package main

import (
	"fmt"
	"strconv"
	"time"
)

//record Global
const n = 3000

type data struct {
	NIK, Nama, JenisR                                                        string
	setoranT                                                                 int
	norekening                                                               string
	tanggalpendaftaran                                                       string
	haripendaftaran, bulanpendaftaran, tahunpendaftaran                      int
	tanggaltransaksiterakhiir                                                string
	haritransaksiterakhiir, bulantransaksiterakhiir, tahuntransaksiterakhiir int
}

var dataN [n]data                                 //array global
var menampung [n]data                             //array global
var i int                                         //variabel global
var counterS, counterG, counterP int              //variabel global
var jumlahdatanasabah, jumlahdatanasabahpasif int //variabel global

// untuk menyelesaikan bagian 1
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
func pengubahbulan(bulan int) int {
	t := time.Now()
	if t.Month().String() == "Januari" {
		return 1
	} else if t.Month().String() == "February" {
		return 2
	} else if t.Month().String() == "Maret" {
		return 3
	} else if t.Month().String() == "April" {
		return 4
	} else if t.Month().String() == "Mei" {
		return 5
	} else if t.Month().String() == "June" {
		return 6
	} else if t.Month().String() == "July" {
		return 7
	} else if t.Month().String() == "August" {
		return 8
	} else if t.Month().String() == "September" {
		return 9
	} else if t.Month().String() == "October" {
		return 10
	} else if t.Month().String() == "November" {
		return 11
	} else {
		return 12
	}
}
func validasipilihanmenulogin(login *string) {
	for *login != "1" && *login != "2" && *login != "3" {
		fmt.Print("Ada yang salah!!! Mohon input angka 1 atau 2 atau 3 : ")
		fmt.Scanln(&*login)
	}
}

func validasiNIK(inputnik string) {
	/*
		prosedur untuk memvalidasi nik apakah memenuhi syarat
		1.user harus menginput kata 1/2/3/4/5/6/7/8/9/0 dan cara penyelesaianu menggunakan ASCII
		2.tidak boleh double dan disni sy menggunakan prosedur lain lagi jadi tinggal sy panggil
		3.panjang data nik harus 16
		4.tidak boleh kosong inputannya
	*/

	for validasidoubleNIK(inputnik) == false || len(inputnik) > 16 || len(inputnik) < 16 || inputnik == "" {
		fmt.Print("error!!! Mohon input NIK 16 digit angka & tidak sama dengan nik yang telah terdaftar	: ")
		fmt.Scanln(&inputnik)
	}
	var j int
	j = 0
	for j < len(inputnik) {
		if inputnik[j] < 48 || inputnik[j] > 57 {
			fmt.Print("error!!! Mohon input NIK 16 digit angka & tidak sama dengan nik yang telah terdaftar	: ")
			fmt.Scanln(&inputnik)
			j = 0
		} else {
			j++
		}
	}
}

func validasidoubleNIK(inputnik string) bool {
	//cek dataN[i]
	var cek, j int
	cek = 0
	j = i - 1
	for cek == 0 && j > -1 {
		if dataN[j].NIK == inputnik {
			cek++
		}
		j--
	}
	return cek == 0
}

/* kenapa yang cuma di validasi nama depan? soalnya perintah dari dosen cuma nama depan yang harus diisi
dan secara logika emang tepat*/
func validasinamadepan(namadepan string) {
	for namadepan == "" {
		fmt.Print("Input Nama Depan		: ")
		fmt.Scanln(&namadepan)
	}
}
func validasisetoranawal() {
	for dataN[i].setoranT < 500 {
		fmt.Print("Setoran Awal minimal 500	: ")
		fmt.Scanln(&dataN[i].setoranT)
	}
}
func generatenorekening() {
	if dataN[i].JenisR == "1" {
		if counterS <= 8 {
			counterS = counterS + 1
			dataN[i].norekening = "XYZ-S" + "0" + "0" + strconv.Itoa(counterS)
		} else if counterS <= 98 {
			counterS = counterS + 1
			dataN[i].norekening = "XYZ-S" + "0" + strconv.Itoa(counterS)
		} else {
			counterS = counterS + 1
			dataN[i].norekening = "XYZ-S" + strconv.Itoa(counterS)
		}
	} else if dataN[i].JenisR == "2" {
		if counterG <= 8 {
			counterG = counterG + 1
			dataN[i].norekening = "XYZ-G" + "0" + "0" + strconv.Itoa(counterG)
		} else if counterG <= 98 {
			counterG = counterG + 1
			dataN[i].norekening = "XYZ-G" + "0" + strconv.Itoa(counterG)
		} else {
			counterG = counterG + 1
			dataN[i].norekening = "XYZ-G" + strconv.Itoa(counterG)

		}
	} else {
		if counterP <= 8 {
			counterP = counterP + 1
			dataN[i].norekening = "XYZ-P" + "0" + "0" + strconv.Itoa(counterP)
		} else if counterP <= 98 {
			counterP = counterP + 1
			dataN[i].norekening = "XYZ-P" + "0" + strconv.Itoa(counterP)
		} else {
			counterP = counterP + 1
			dataN[i].norekening = "XYZ-P" + strconv.Itoa(counterP)

		}
	}
	fmt.Print("No rekening nasabah : ")
	fmt.Println(dataN[i].norekening)
	fmt.Println("")
}
func tanggal(ke int, tanggal *string, hari, bulan, tahun *int) {

	time := time.Now()

	*tahun = time.Year()
	*hari = time.Day()
	*bulan = pengubahbulan(*bulan)
	*tanggal = strconv.Itoa(*hari) + "-" + string(*bulan) + "-" + strconv.Itoa(*tahun)

}
func validasiinputanjenisrekening() {
	for dataN[i].JenisR != "1" && dataN[i].JenisR != "2" && dataN[i].JenisR != "3" {
		fmt.Println("Mohon Pilih no jenis rekening yang ingin dibuka yang valid : ")
		fmt.Scanln(&dataN[i].JenisR)
	}
}

/*
Disini tempat pemrosesan penginputan data pada soal 1 (login cs)
*/
// **********************************************************************************************************************
// **********************************************************************************************************************
// **********************************************************************************************************************
func pendaftaran(counterS, counterG, counterP int) {
	var namadepan, namatengah, namabelakang string
	var loopingpenginputandata, inputnik string

	loopingpenginputandata = "1"
	//perulangan akan berhenti jika user g mau input lagi dan pengulangan udah melewati batasan array
	for i < n && loopingpenginputandata == "1" {
		//nama depan belakang dan tengah di inialisasi dengan g ada isi karena kalau di luping lagi data sebelumnya akan terus terpakai
		namadepan = ""
		namatengah = ""
		namabelakang = ""
		fmt.Print("Input NIK			: ")
		fmt.Scanln(&inputnik)
		validasiNIK(inputnik)
		dataN[i].NIK = inputnik

		fmt.Println("-----------------------------------------------")
		fmt.Print("Input Nama Depan		: ")
		fmt.Scanln(&namadepan)
		validasinamadepan(namadepan)
		fmt.Print("Input Nama Tengah		: ")
		fmt.Scanln(&namatengah)
		fmt.Print("Input Nama Belakang		: ")
		fmt.Scanln(&namabelakang)
		dataN[i].Nama = namadepan + " " + namatengah + " " + namabelakang

		fmt.Println("-----------------------------------------------")
		fmt.Println("Jenis rekening : ")
		fmt.Println("1. Silver")
		fmt.Println("2. Gold")
		fmt.Println("3. Platinum")
		fmt.Println("Pilih no jenis rekening yang ingin dibuka : ")
		fmt.Scanln(&dataN[i].JenisR)
		validasiinputanjenisrekening()

		fmt.Println("-----------------------------------------------")
		fmt.Print("Input Setoran Awal	: ")
		fmt.Scanln(&dataN[i].setoranT)
		validasisetoranawal()

		generatenorekening()

		//ini untuk menginput tanggal pendaftaran
		tanggal(i, &dataN[i].tanggalpendaftaran, &dataN[i].haripendaftaran, &dataN[i].bulanpendaftaran, &dataN[i].tahunpendaftaran)
		//untuk menginput tanggal pendaftaran
		tanggal(i, &dataN[i].tanggaltransaksiterakhiir, &dataN[i].haritransaksiterakhiir, &dataN[i].bulantransaksiterakhiir, &dataN[i].tahuntransaksiterakhiir)

		fmt.Println("Masih ingin melanjutkan penginputan?")
		fmt.Println("1. YA")
		fmt.Println("2. TIDAK")
		fmt.Scanln(&loopingpenginputandata)
		i++
		fmt.Println("-----------------------------------------------")
		jumlahdatanasabah++
		fmt.Println(jumlahdatanasabah)
	}

}

// **********************************************************************************************************************
// **********************************************************************************************************************

//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@

// FUNGSI UNTUK NOMOR 2 %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

func validasiinputnomorrekening(nomorrekening string) bool {
	var g int
	g = 0
	for (nomorrekening != dataN[g].norekening) && (g < jumlahdatanasabah-1) {
		g++
	}
	return nomorrekening == dataN[g].norekening
}
func validasiinputnomorrekening2(nomorrekening string) {

	for validasiinputnomorrekening(nomorrekening) == false || nomorrekening == "" || len(nomorrekening) != 8 {
		fmt.Print("Input nomor rekening yang valid : ")
		fmt.Scanln(&nomorrekening)
	}
}
func validasijumlahsetoran(jumlahsetoran, bagianarray int) {
	if dataN[bagianarray].norekening[4] == 83 { // menggunakan ascii XYZ-S001
		for jumlahsetoran > 50000 {
			fmt.Println("Untuk Jenis rekening Silver max setoran 50000")
			fmt.Print("Input Jumlah Setoran : ")
			fmt.Scanln(&jumlahsetoran)
		}
	} else if dataN[bagianarray].norekening[4] == 71 {
		for jumlahsetoran > 200000 {
			fmt.Println("Untuk Jenis rekening Gold max setoran 200000")
			fmt.Print("Input Jumlah Setoran : ")
			fmt.Scanln(&jumlahsetoran)
		}
	} else if dataN[bagianarray].norekening[4] == 80 {
		for jumlahsetoran > 500000 {
			fmt.Println("Untuk Jenis rekening Platinum max setoran 500000")
			fmt.Print("Input Jumlah Setoran : ")
			fmt.Scanln(&jumlahsetoran)
		}
	}
}
func caridatanomorrekening(bagianarray *int, nomorrekening string) {
	var cek bool
	cek = false
	*bagianarray = 0
	for *bagianarray < jumlahdatanasabah && cek == false {
		//BAGIAN UNTUK CEK KALAU UDAH TRUE
		cek = nomorrekening == dataN[*bagianarray].norekening
		*bagianarray++
	}
	*bagianarray--

}
func validasiloopingpenginputanteller(loopingpenginputanteller string) {
	for loopingpenginputanteller != "1" && loopingpenginputanteller != "2" {
		fmt.Println("Pilih input 1 atau 2")
		fmt.Scanln(&loopingpenginputanteller)
	}
}

// jadi saya tinggal manggil-manggil fungsi validasi yang telah sy buat untuk menyelesaikan soal bagian 2
// **********************************************************************************************************************
// ***********************************************************************************************************
//&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&
//&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&
//&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&

func setortunai() {
	var bagianarray int
	var nomorrekening string
	var jumlahsetoran int
	var loopingpenginputanteller string
	loopingpenginputanteller = "1"
	for loopingpenginputanteller == "1" {
		fmt.Print("Input nomor rekening : ")
		fmt.Scanln(&nomorrekening)
		validasiinputnomorrekening2(nomorrekening)
		caridatanomorrekening(&bagianarray, nomorrekening) //disini saya akan mencari array ke berapa yang ada nomor rekeningnya yang di cari
		fmt.Println("-----------------------------------------------")

		fmt.Print("Input jumlah setoran : ")
		fmt.Scanln(&jumlahsetoran)
		validasijumlahsetoran(jumlahsetoran, bagianarray)

		dataN[bagianarray].setoranT = jumlahsetoran + dataN[bagianarray].setoranT

		//menentkan otomatis tanggal

		tanggal(bagianarray, &dataN[bagianarray].tanggaltransaksiterakhiir, &dataN[i].haritransaksiterakhiir, &dataN[i].bulantransaksiterakhiir, &dataN[i].tahuntransaksiterakhiir)
		fmt.Println("Masih ingin melanjutkan penginputan?")
		fmt.Println("1. YA")
		fmt.Println("2. TIDAK")
		fmt.Scanln(&loopingpenginputanteller)
		validasiloopingpenginputanteller(loopingpenginputanteller)
		fmt.Println("-----------------------------------------------")
	}
}

// **********************************************************************************************************************
// **********************************************************************************************************************
//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
//****************************************************************************************************&&&&&&&&&&&&&&&&&&&&

//nomor 3 login manager
//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
func outputdatanasabah(bagianarray int) {
	var z string
	z = "1"
	for z == "1" {
		fmt.Println("=============================================")
		fmt.Println("=============================================")
		fmt.Println("=============================================")
		fmt.Print("Nama 				: ")
		fmt.Println(dataN[bagianarray].Nama)
		fmt.Println("")
		fmt.Print("NIK 				: ")
		fmt.Println(dataN[bagianarray].NIK)
		fmt.Println("")
		if dataN[bagianarray].JenisR == "1" {
			fmt.Println("Jenis Rekening 				: Silver")
		} else if dataN[bagianarray].JenisR == "2" {
			fmt.Println("Jenis Rekening					: Gold")
		} else {
			fmt.Println("Jenis Rekening 				: Platinum")
		}
		fmt.Println("")
		fmt.Print("Saldo Nasabah 			: ")
		fmt.Println(dataN[bagianarray].setoranT)
		fmt.Println("")
		fmt.Print("Tanggal Pendaftaran Nasabah 	: ")
		fmt.Println(dataN[bagianarray].tanggalpendaftaran)
		fmt.Println("")
		fmt.Print("Tanggal Transaksi Terakhir 	: ")
		fmt.Println(dataN[bagianarray].tanggaltransaksiterakhiir)
		fmt.Println("")
		fmt.Println("Apakah Masih ingin melihat data nasabah?")
		fmt.Println("1.YA")
		fmt.Println("2.TIDAK")
		fmt.Scanln(&z)
		for z != "1" && z != "2" {
			fmt.Println("Mohon Input angka 1 atau 2 ")
			fmt.Scanln(&z)
		}
		fmt.Println("")
		fmt.Println("=============================================")
		fmt.Println("=============================================")
		fmt.Println("=============================================")
	}

}
func sortingselectionsort() { //masih salah
	//disini saya melakukan sorting menggunakan teknik selection sort
	var min int
	var t, p int
	const isi = 1
	var tampung [isi]data

	for t = 0; t < jumlahdatanasabah-1; t++ {
		//perlangan untuk array ke-t(0,1,2.....)
		min = t
		//
		for p = 0; p < jumlahdatanasabah; p++ {
			//perulangan untuk mengecek min array ke-t terhadap array selanjutnya
			if dataN[min].Nama > dataN[p].Nama {
				min = p
			}
		}
		//menggunakan array tampung sebagai wadah proses penukaran
		tampung[0] = dataN[t]
		dataN[t] = dataN[min]
		dataN[min] = tampung[0]
	}

}
func outputdatanasabah2() {
	var l int
	l = 0
	for l < jumlahdatanasabah {
		fmt.Println("=============================================")
		fmt.Println("=============================================")
		fmt.Println("=============================================")
		fmt.Print("Nama 				: ")
		fmt.Println(dataN[l].Nama)
		fmt.Println("")
		fmt.Print("NIK 				: ")
		fmt.Println(dataN[l].NIK)
		fmt.Println("")
		if dataN[l].JenisR == "1" {
			fmt.Println("Jenis Rekening 			: Silver")
		} else if dataN[l].JenisR == "2" {
			fmt.Println("Jenis Rekening				: Gold")
		} else {
			fmt.Println("Jenis Rekening 			: Platinum")
		}
		fmt.Println("")
		fmt.Print("Nomor Rekening Nasabah 			: ")
		fmt.Println(dataN[l].norekening)
		fmt.Println("")
		fmt.Print("Saldo Nasabah 			: ")
		fmt.Println(dataN[l].setoranT)
		fmt.Println("")
		fmt.Print("Tanggal Pendaftaran Nasabah 	: ")
		fmt.Println(dataN[l].tanggalpendaftaran)
		fmt.Println("")
		fmt.Print("Tanggal Transaksi Terakhir 	: ")
		fmt.Println(dataN[l].tanggaltransaksiterakhiir)
		fmt.Println("=============================================")
		fmt.Println("=============================================")
		fmt.Println("=============================================")
		fmt.Println("")
		l++
	}
}
func prosesinsertionsort(t, p int, tampung data) {
	if dataN[t].norekening[5] == '0' && dataN[t].norekening[6] == '0' {
		if dataN[t].norekening[7] < dataN[p].norekening[7] {
			tampung = dataN[p]
			dataN[p] = dataN[t]
			dataN[t] = tampung
		}
	} else if dataN[t].norekening[5] == '0' {
		if dataN[t].norekening[6] == dataN[p].norekening[6] {
			if dataN[t].norekening[7] < dataN[p].norekening[7] {
				tampung = dataN[p]
				dataN[p] = dataN[t]
				dataN[t] = tampung
			}
		} else {
			if dataN[t].norekening[6] < dataN[p].norekening[6] {
				tampung = dataN[p]
				dataN[p] = dataN[t]
				dataN[t] = tampung
			}
		}
	}
}
func InsertionSort() { //login manager no 4
	var t, p int
	var tampung data
	fmt.Print("Jumlah Nasabah : ")
	fmt.Println(jumlahdatanasabah, " Nasabah")
	for t = 1; t < jumlahdatanasabah; t++ {
		for p = t - 1; p >= 0; p-- {
			/*
				1. JIKA JENISNYA SAMA
				2.	JIKA 004 ATAU 043 ATAU 432
			*/
			if dataN[t].norekening[4] == 'S' && dataN[p].norekening[4] == 'S' {
				prosesinsertionsort(t, p, tampung)
			} else if dataN[t].norekening[4] == 'S' && dataN[p].norekening[4] == 'G' {
				tampung = dataN[p]
				dataN[p] = dataN[t]
				dataN[t] = tampung
			} else if dataN[t].norekening[4] == 'S' && dataN[p].norekening[4] == 'P' {
				tampung = dataN[p]
				dataN[p] = dataN[t]
				dataN[t] = tampung
			} else if dataN[t].norekening[4] == 'G' && dataN[p].norekening[4] == 'G' {
				prosesinsertionsort(t, p, tampung)
			} else if dataN[t].norekening[4] == 'G' && dataN[p].norekening[4] == 'P' {
				tampung = dataN[p]
				dataN[p] = dataN[t]
				dataN[t] = tampung
			} else if dataN[t].norekening[4] == 'G' && dataN[p].norekening[4] == 'S' {
				//G TERJADI APA"
			} else if dataN[t].norekening[4] == 'P' && dataN[p].norekening[4] == 'G' {
				//G TERJADI APA"
			} else if dataN[t].norekening[4] == 'P' && dataN[p].norekening[4] == 'S' {
				//G TERJADI APA"
			} else if dataN[t].norekening[4] == 'P' && dataN[p].norekening[4] == 'P' {
				prosesinsertionsort(t, p, tampung)
			}
		}
	}
}
func sortingberdasarkantanggal() {
	var t, p int
	var tamp data
	fmt.Print("Jumlah Nasabah Yang Tidak Melakukan Transaksi Lebih Dari 6 Bulan : ")
	fmt.Println(jumlahdatanasabahpasif, " Nasabah")
	for t = 1; t < jumlahdatanasabahpasif; t++ {
		for p = t - 1; p >= 0; p-- {
			//cek tahun apakah sama g.
			if menampung[t].tahuntransaksiterakhiir == menampung[p].tahuntransaksiterakhiir {
				//kalau sama cek lagi apakah bulannya sama g
				if menampung[t].bulantransaksiterakhiir == menampung[p].bulantransaksiterakhiir {
					//kalau sama cek harinya
					if menampung[t].haritransaksiterakhiir < menampung[p].haritransaksiterakhiir {
						//proses penukaran isi array tampung dibantu dengan array tamp
						tamp = menampung[p]
						menampung[p] = menampung[t]
						menampung[t] = tamp
					}
				}
			} else if menampung[t].tahuntransaksiterakhiir < menampung[p].tahuntransaksiterakhiir {
				//kalau yg array tampung[t] lebih kecil dari tampung[p] otomatis tampung[t] lebih dulu dari kecil
				tamp = menampung[p]
				menampung[p] = menampung[t]
				menampung[t] = tamp
			}
		}
	}
}
func outputdatanasabahberdasarkantanggal() {
	var z int
	z = 0
	for z < jumlahdatanasabahpasif {
		fmt.Println("=============================================")
		fmt.Println("=============================================")
		fmt.Println("=============================================")
		fmt.Print("Nama 				: ")
		fmt.Println(menampung[z].Nama)
		fmt.Println("")
		fmt.Print("NIK 				: ")
		fmt.Println(menampung[z].NIK)
		fmt.Println("")
		fmt.Print("Nomor Rekening : ")
		fmt.Print(menampung[z].norekening)
		fmt.Println("")
		if menampung[z].JenisR == "1" {
			fmt.Println("Jenis Rekening 				: Silver")
		} else if menampung[z].JenisR == "2" {
			fmt.Println("Jenis Rekening					: Gold")
		} else {
			fmt.Println("Jenis Rekening 				: Platinum")
		}
		fmt.Println("")
		fmt.Print("Saldo Nasabah 			: ")
		fmt.Println(menampung[z].setoranT)
		fmt.Println("")
		fmt.Print("Tanggal Pendaftaran Nasabah 	: ")
		fmt.Println(menampung[z].tanggalpendaftaran)
		fmt.Println("")
		fmt.Print("Tanggal Transaksi Terakhir 	: ")
		fmt.Println(menampung[z].tanggaltransaksiterakhiir)
		fmt.Println("")

		fmt.Println("")
		fmt.Println("=============================================")
		fmt.Println("=============================================")
		fmt.Println("=============================================")
		z++
	}

}
func mengisikearraymenampung() { //masih salah
	var tahun, bulan, hari, f int
	time := time.Now()
	tahun = time.Year()
	bulan = pengubahbulan(bulan)
	hari = time.Day()

	//masih salah
	for f = 0; f < jumlahdatanasabah; f++ {
		for r = 0; r < jumlahdatanasabah; r++ {
			if dataN[r].tahuntransaksiterakhiir == tahun {
				if bulan == dataN[r].tahuntransaksiterakhiir+6 {
					if hari > dataN[r].haritransaksiterakhiir {
						menampung[f] = dataN[r]
						jumlahdatanasabahpasif++
					}
				} else if bulan > dataN[r].tahuntransaksiterakhiir+6 {
					menampung[f] = dataN[r]
					jumlahdatanasabahpasif++
				}
			} else {
				if tahun >= dataN[r].tahuntransaksiterakhiir+1 {
					menampung[f] = dataN[r]
					jumlahdatanasabahpasif++
				} else {
					if bulan == 7 && dataN[r].bulantransaksiterakhiir == 1 {
						if hari > dataN[r].haritransaksiterakhiir {
							menampung[f] = dataN[r]
							jumlahdatanasabahpasif++
						}
					} else if bulan == 8 && dataN[r].bulantransaksiterakhiir == 2 {
						if hari > dataN[r].haritransaksiterakhiir {
							menampung[f] = dataN[r]
							jumlahdatanasabahpasif++
						}
					} else if bulan == 9 && dataN[r].bulantransaksiterakhiir == 3 {
						if hari > dataN[r].haritransaksiterakhiir {
							menampung[f] = dataN[r]
							jumlahdatanasabahpasif++
						}

					} else if bulan == 10 && dataN[r].bulantransaksiterakhiir == 4 {
						if hari > dataN[r].haritransaksiterakhiir {
							menampung[f] = dataN[r]
							jumlahdatanasabahpasif++
						}

					} else if bulan == 11 && dataN[r].bulantransaksiterakhiir == 5 {
						if hari > dataN[r].haritransaksiterakhiir {
							menampung[f] = dataN[r]
							jumlahdatanasabahpasif++
						}

					} else if bulan == 12 && dataN[r].bulantransaksiterakhiir == 6 {
						if hari > dataN[r].haritransaksiterakhiir {
							menampung[f] = dataN[r]
							jumlahdatanasabahpasif++
						}

					}
				}
			}
		}
	}
	fmt.Println(menampung)
}
func loginmanager() {
	var nomorrekening string
	var bagianarray, r int
	var kondisi, ulang, kondisi1 string
	//paksa masuk dulu
	ulang = "1"
	for ulang == "1" {
		fmt.Println("SELAMAT DATANG TUAN MANAGER")
		fmt.Println("Silahkan Pilih Menu")
		fmt.Println("1.Cek data 1 nasabah tertentu")
		fmt.Println("2.Cek jumlah data dan semua nasabah berdasarkan urutan nama")
		fmt.Println("3.Cek jumlah dan data semua nasabah yang tidak melakukan transaksi lebih dari 6 bulan berturut-turut")
		fmt.Println("4.Cek jumlah dan data nasabah per jenis rekening secara terurut berdasarkan nomor rekening.")
		fmt.Scanln(&kondisi1)
		//untuk pengecekkan/validasi inputan
		for kondisi1 != "1" && kondisi1 != "2" && kondisi1 != "3" && kondisi1 != "4" {
			fmt.Println("Input 1 atau 2 atau 3 atau 4")
			fmt.Scanln(&kondisi1)
		}

		if kondisi1 == "1" {
			fmt.Println("masukan nomor rekening nasabah yang ingin di cek datanya : ")
			fmt.Scanln(&nomorrekening)
			//pemvalidasi apakah nomor rekening udah terdaftar atau belum. kalau belum akan meminta inputan terus
			validasiinputnomorrekening2(nomorrekening)
			//disini saya akan mencari array ke berapa yang ada nomor rekeningnya yang di cari
			caridatanomorrekening(&bagianarray, nomorrekening)
			//-------------------------
			outputdatanasabah(bagianarray)
		} else if kondisi1 == "2" {
			fmt.Println("Ingin Melihat data semua nasabah?")
			fmt.Println("1.YA")
			fmt.Println("2.TIDAK")
			fmt.Scanln(&kondisi)
			//untuk pengecekkan/validasi inputan
			for kondisi != "1" && kondisi != "2" {
				fmt.Println("Input 1 atau 2")
				fmt.Println(&kondisi)
			}
			//jika mau berarti akan menampilkan seluruh data nasabah
			if kondisi == "1" {
				fmt.Print("Jumlah Nasabah Bank XYZ sebanyak : ")
				fmt.Println(jumlahdatanasabah, "Nasabah")
				fmt.Println("")
				sortingselectionsort()
				outputdatanasabah2()
				/*for q = 0; q < jumlahdatanasabah; q++ {
					fmt.Println(dataN[q])
				}*/
			}
		} else if kondisi1 == "3" {
			mengisikearraymenampung()
			//setelah array meanmpung terisi dengan isi array dataN nasabah yg g melakukan transaksi lebih dari 6 bulan maka ksekarang kia sorting berdasarkan urutan
			sortingberdasarkantanggal()
			outputdatanasabahberdasarkantanggal()

		} else {
			InsertionSort()
			outputdatanasabah2()

		}
		fmt.Println("---------------------------------------------------------")
		//untuk mengecek bahwa manager apakah masih mau melakukan sesuatu dimenu manager
		fmt.Println("Masih ingin melakukan sesuatu di menu login?")
		fmt.Println("1.YA")
		fmt.Println("2.TIDAK")
		fmt.Println(&ulang)
		//untuk pengecekkan/validasi inputan
		for ulang != "1" && ulang != "2" {
			fmt.Println("Input 1 atau 2")
			fmt.Println(&ulang)
		}
		fmt.Println("---------------------------------------------------------")

	}
}

// BAGIAN PASSWORD CS,TELLER, DAN MANAGER
func validasiusernamecs(usernameCS, passwordCS string) {
	for usernameCS != "admin" || passwordCS != "admin" {
		fmt.Println("**** Ada yang salah nih ****")
		fmt.Print("Input Username : ")
		fmt.Scan(&usernameCS)
		fmt.Print("Input Password : ")
		fmt.Scan(&passwordCS)
	}
}
func validasiusernameteller(usernameTELLER, passwordTELLER string) {
	for usernameTELLER != "admin" || passwordTELLER != "admin" {
		fmt.Println("**** Ada yang salah nih ****")
		fmt.Print("Input Username : ")
		fmt.Scan(&usernameTELLER)
		fmt.Print("Input Password : ")
		fmt.Scan(&passwordTELLER)
	}
}
func validasiusernamemanager(usernameMANAGER, passwordMANAGER string) {
	for usernameMANAGER != "admin" || passwordMANAGER != "admin" {
		fmt.Println("**** Ada yang salah nih ****")
		fmt.Print("Input Username : ")
		fmt.Scan(&usernameMANAGER)
		fmt.Print("Input Password : ")
		fmt.Scan(&passwordMANAGER)
	}
}

// BAGIAN PASSWORD CS,TELLER, DAN MANAGER

func main() {
	//kamus lokal
	var login string
	var usernameCS, passwordCS string
	var usernameTELLER, passwordTELLER string
	var usernameMANAGER, passwordMANAGER string
	var loppingforever bool

	/*var di luar prosedur pendaftaran karena nanti kalau
	di looping lagi counternya jadi 0 lagi
	*/
	var counterS, counterG, counterP int
	counterS = 0
	counterG = 0
	counterP = 0
	jumlahdatanasabah = 0
	jumlahdatanasabahpasif = 0
	i = 0
	//****************************************************
	fmt.Println("***********  SELAMAT DATANG DI BANK XYZ  ***********")
	/*
		disini saya membuat looping tak terhingga
		karena program akan selalu digunakan
	*/
	loppingforever = true
	for loppingforever == true {
		fmt.Println("1.Login Costumer Service")
		fmt.Println("2.Login Teller")
		fmt.Println("3.Login Manager")
		fmt.Print("Pilih Menu Login : ")
		fmt.Scanln(&login)
		validasipilihanmenulogin(&login)
		fmt.Println("-----------------------------------------------")
		if login == "1" {
			fmt.Print("Input Username : ")
			fmt.Scanln(&usernameCS)
			fmt.Print("Input Password : ")
			fmt.Scanln(&passwordCS)
			validasiusernamecs(usernameCS, passwordCS)
			fmt.Println("-----------------------------------------------")
			pendaftaran(counterS, counterG, counterP)
		} else if login == "2" {
			fmt.Print("Input Username : ")
			fmt.Scanln(&usernameTELLER)
			fmt.Print("Input Password : ")
			fmt.Scanln(&passwordTELLER)
			validasiusernameteller(usernameTELLER, passwordTELLER)
			fmt.Println("-----------------------------------------------")
			setortunai()

		} else if login == "3" {
			fmt.Print("Input Username : ")
			fmt.Scanln(&usernameMANAGER)
			fmt.Print("Input Password : ")
			fmt.Scanln(&passwordMANAGER)
			validasiusernamemanager(usernameMANAGER, passwordMANAGER)
			fmt.Println("-----------------------------------------------")
			loginmanager()
		}
	}
}
