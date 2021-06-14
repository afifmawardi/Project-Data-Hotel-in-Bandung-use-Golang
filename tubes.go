/* Nama : Afifuddin Mawardi
   NIM	: 1301194113
   Nama : M.NOER ICHZAN A.
   NIM	: 1301194114 */

package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)
const N = 10000

func clear() {
	cmd := exec.Command("cmd","/c","cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

type data struct {
	harga int
	luas float64
	fasilitas string
	kelas string
	nomor int
	lokasi string
	}

type pengunjung struct {
	nama string
	umur int
	nohp int
	email string
	}
type datap [N]pengunjung
type apart [N]data

var (
	read string
	i 	int
	pilihtempat string
	pilihharga string
	a,b,c,d,e,q int
	x int
	aa = 1000000
	bb = 2000000
	cc = 2500000
	dd = 3000000
)

var A apart
var B datap

func hargaapart() {

	//prosedur
	
	
	for (pilihharga != "5") {
		clear()
		fmt.Println("===List Harga Apart Berdasarkan Lokasi===\n")
		fmt.Println("1. Ciwidey")
		fmt.Println("2. Dago")
		fmt.Println("3. Punclut")
		fmt.Println("4. Lembang")
		fmt.Println("5. Kembali")
		fmt.Print("\nPilih menu: ")
		fmt.Scanln(&pilihharga)
		
		if (pilihharga == "1") {
			clear()
			fmt.Println("Harga Apart di Ciwidey adalah : Rp ",aa)
			fmt.Scanln(&read)
		} else if (pilihharga == "2") {
			clear()
			fmt.Println("Harga Apart di Dago adalah : Rp ",bb)
			fmt.Scanln(&read)
		} else if (pilihharga == "3") {
			clear()
			fmt.Println("Harga Apart di Punclut adalah : Rp ",cc)
			fmt.Scanln(&read)
		} else if (pilihharga == "4") {
			clear()
			fmt.Println("Harga Apart di Lembang adalah : Rp ",dd)
			fmt.Scanln(&read)
		}
	}
}

func pilihapart(A *apart, i int) {
	
	var jenis string
	
	clear()
	
	fmt.Println("Pilih jenis Apartment: ")
	fmt.Println("")
	fmt.Println("1. Reguler")
	fmt.Println("2. Premium")
	fmt.Println("3. Private")
	fmt.Println("")
	fmt.Print("Pilih Menu :")
	fmt.Scan(&jenis)
	if jenis == "1" {
		A[i].kelas = "Reguler"
		A[i].harga = x + 500000
		A[i].fasilitas = "\n AC\n Kulkas\n TV Flat 32 Inch\n Shower with hot water\n Mini bar\n Bed size king\n Lemari besar\n Kitchen set\n Kompor gas & westafel\n Dispenser (hot / cold / normal)\n"
		A[i].luas = 40 

		} else if jenis == "2" {
		A[i].kelas = "Premium"
		A[i].harga = x + 1000000
		A[i].fasilitas = "\n Cliff view F&B area\n Pool (lagoon pool + infinity pool + kid’s pool)\n 3 on 3 basketball half-court\n Fitness center\n Mini futsal\n Jogging hill track\n Rooftop garden\n Amphitheater\n Access card\n F&B area dan\n CCTV Camera.\n"
		A[i].luas = 60
		} else if jenis == "3" {
		
		A[i].kelas = "Private"
		A[i].harga = x + 2000000
		A[i].fasilitas = "\n Water Heater Instalation \n TV Cable \n Internet Connection \n Children & Adult Swiming Pool\n Covered Parking Area\n Jogging Track\n Sky Garden & Flower Garden\n Water Treatment Plant (WTP)\n Pusat Kebugaran, Fitness, Gym, Spa & Saloon\n Kios\n Alfresco Dining\n 24H Security & CCTV\n Lapangan Terbuka\n Tempat BBQ\n Gedung Perkumpulan\n Ruang Permainan\n Ruang Santai\n Ruang Rapat\n Arena bermain\n Ruang Santai Lantai Atas\n"
		A[i].luas = 80
	}
	time.Sleep(time.Second * 2)
	
	clear()

	fmt.Println("")
	fmt.Println("|--Data dan Pilihan Anda --|\n")
	fmt.Println("\nJenis Apart Pilihan anda : ", A[i].kelas)
	A[i].nomor = i + 200
	fmt.Println("\nNomor Apart Anda :", A[i].nomor)
	fmt.Println("\nTotal Biaya Anda :", A[i].harga)
	fmt.Print("Dengan fasilitas 	:", A[i].fasilitas)
	fmt.Scanln(&read)
	time.Sleep(time.Second * 5)
	

}
func inputdata(B *datap, i int ){
	
	clear()

	fmt.Println(" Input data Pemesan ")
	fmt.Println()
	fmt.Print("Nama :")
	fmt.Scan(&B[i].nama)
	fmt.Print("Umur :")
	fmt.Scan(&B[i].umur)
	fmt.Print("Nomor HP :")
	fmt.Scan(&B[i].nohp)
	fmt.Print("Email :")
	fmt.Scan(&B[i].email)
	
	clear()
	fmt.Println("===Data Pemesan===")
	fmt.Println("Nama				: ",B[i].nama)
	fmt.Println("umur				: ",B[i].umur)
	fmt.Println("Nomor Handphone  \t\t: ",B[i].nohp)
	fmt.Println("Email			 	: ",B[i].email)
	fmt.Print("==================")
	fmt.Scanln(&read)
	time.Sleep(time.Second*3)
	
}
func bookingapart(A *apart, i int) {
	
	// Prosedur

	var ruangan string
	
	clear()

	fmt.Println("Pilih Lokasi Yang di Inginkan ?")
	fmt.Println("")
	fmt.Println("1. Ciwidey")
	fmt.Println("2. Dago")
	fmt.Println("3. Punclut")
	fmt.Println("4. Lembang")
	fmt.Print("Pilih Menu: ")
	fmt.Scan(&ruangan)
	if ruangan == "1" {
		clear()
		A[i].lokasi = "Ciwidey"
		x = aa
		pilihapart(A, i)
	} else if ruangan == "2" {
		clear()
		A[i].lokasi = "Dago"
		x = bb
		pilihapart(A, i)
	} else if ruangan == "3" {
		clear()
		A[i].lokasi = "Punclut"
		x = cc
		pilihapart(A, i)
	} else if ruangan == "4" {
		clear()
		A[i].lokasi = "Lembang"
		x = dd
		pilihapart(A, i)
	}
}


func editdata(B *datap, i int) {
	
	var f string
	var found = false
	clear()
	fmt.Println("Cari nama pemesan :")
	fmt.Scan(&f)
	
		for j:=0; j < i && !found; j++ {
			if B[j].nama == f {
				fmt.Print("Nama :")
				fmt.Scan(&B[j].nama)
				fmt.Print("Umur :")
				fmt.Scan(&B[j].umur)
				fmt.Print("Nomor HP :")
				fmt.Scan(&B[j].nohp)
				fmt.Print("Email :")
				fmt.Scan(&B[j].email)
				found = true
			} else {
			fmt.Println("Maaf data tidak ditemukan")
			}
	}
}

func checkout(B *datap,A *apart, i int) {
	
	var key string
	clear()
	fmt.Println("Cari nama yang ingin di delete :")
	fmt.Scan(&key)
	
		for k:=0; k < i; k++ {
			if B[k].nama == key {
				B[k].nama = ""
				B[k].umur = 0
				B[k].nohp = 0
				B[k].email = ""
			} else {
			fmt.Println("Maaf data tidak ditemukan")
			}
		}
}

func sortnama(B datap,A apart, i int) {

	for l := 0; l < i; l++ {
		t := B[l].nama
		j := l - 1
		for j >= 0 && B[j].nama > t {
			B[j+1].nama = B[j].nama
			j--
		}
		B[j+1].nama = t
	}
	clear()
	for l:=0; l < i; l++ {
		fmt.Println("======DATA  PEMESAN======")
		fmt.Println("|Nama				: ",B[l].nama)
		fmt.Println("=========================")
		fmt.Println("|Umur				: ",B[l].umur)
		fmt.Println("|Nomor Handphone		: ",B[l].nohp)
		fmt.Println("|Email				: ",B[l].email)
		fmt.Println("|Jenis Kamar Pemesan \t\t: ",A[l].kelas)
		fmt.Println("|Lokasi Pemesan      \t\t: ",A[l].lokasi)
		fmt.Println("|No Kamar Pemesan    \t\t: ",A[l].nomor)
		fmt.Println("|Luas Kamar Pemesan  \t\t: ",A[l].luas)
		fmt.Println("=========================")
		fmt.Print()
		fmt.Scanln(&read)
	}
}

func sortumur(B datap,A apart, i int) {
	m := i - 1
	for m > 0 {
		max := 0
		j := 1
		for j < i {
			if B[j].umur < B[max].umur {
				max = j
			}
			j++
		}
		t := B[max].umur
		B[max].umur = B[m].umur
		B[m].umur = t
		m--
	}
		clear()
	for l:=0; l < i; l++ {
		fmt.Println("======DATA  PEMESAN======")
		fmt.Println("|Umur				: ",B[l].umur)
		fmt.Println("=========================")
		fmt.Println("|Nama				: ",B[l].nama)
		fmt.Println("|Nomor Handphone		: ",B[l].nohp)
		fmt.Println("|Email				: ",B[l].email)
		fmt.Println("|Jenis Kamar Pemesan \t\t: ",A[l].kelas)
		fmt.Println("|Lokasi Pemesan      \t\t: ",A[l].lokasi)
		fmt.Println("|No Kamar Pemesan    \t\t: ",A[l].nomor)
		fmt.Println("|Luas Kamar Pemesan  \t\t: ",A[l].luas)
		fmt.Println("=========================")
		fmt.Print()
		fmt.Scanln(&read)
}
}

func carinama(B datap, A apart, i int) {
	// prosedur

	var f string
	
	var found bool = false
	
	clear()

	fmt.Print("Masukkan Nama Pemesan : ")
	fmt.Scan(&f)

	for e:=0; (!found && e < i); e++ {
		if ( f == B[e].nama ) {
			clear()
			fmt.Println("======DATA  PEMESAN======")
			fmt.Println("|Nama					: ",B[e].nama)
			fmt.Println("=========================")
			fmt.Println("|umur					: ",B[e].umur)
			fmt.Println("|Nomor Handphone	 	: ",B[e].nohp)
			fmt.Println("|Email			 		: ",B[e].email)
			fmt.Println("|Jenis Kamar Pemesan   : ",A[e].kelas)
			fmt.Println("|Lokasi Pemesan        : ",A[e].lokasi)
			fmt.Println("|No Kamar Pemesan      : ",A[e].nomor)
			fmt.Println("|Luas Kamar Pemesan    : ",A[e].luas)
			fmt.Println("=========================")
			found = true
			fmt.Scanln(&read)
			f = ""
		} else {
			clear()
			fmt.Print("Maaf, Nama Pemesan tidak Ditemukan")
			fmt.Scanln(&read)
		}
	}
}

func carikamar(B datap, A apart,i int) {
	
	//Prosedur

	var f int
	
	var found = false
	
	clear()
	
	fmt.Print("Masukkan Nomor Kamar : ")
	fmt.Scan(&f)
	
	for z:=0; (!found && z < i); z++ {
		if f == A[z].nomor {
			clear()
			fmt.Println("======DATA  PEMESAN======")
			fmt.Println("|No Kamar Pemesan      : ",A[z].nomor)
			fmt.Println("=========================")
			fmt.Println("|Nama			 : ",B[z].nama)
			fmt.Println("|Umur			 : ",B[z].umur)
			fmt.Println("|Nomor Handphone	: ",B[z].nohp)
			fmt.Println("|Email		: ",B[z].email)
			fmt.Println("|Jenis Kamar Pemesan   : ",A[z].kelas)
			fmt.Println("|Lokasi Pemesan        : ",A[z].lokasi)
			fmt.Println("|Luas Kamar Pemesan    : ",A[z].luas)
			fmt.Print("=========================")
			found = true
			time.Sleep(time.Second*3)
			fmt.Scanln(&read)
			f = 0
		} else {
			clear()
			fmt.Print("Maaf, Pemesan dengan nomor kamar tersebut tidak ditemukan")
			fmt.Scanln(&read)
		}
	}
}

func list(B datap, A apart, i int) {
	
	//Prosedur

	var p string
	
	clear()

	for (p != "3") {
		clear()
		fmt.Println("===================================")
		fmt.Println("| 1. List Menurut Nama            |")
		fmt.Println("| 2. List Menurut Umur            |")
		fmt.Println("| 3. Kembali                      |")
		fmt.Println("===================================\n")
		fmt.Print("Pilih Menu : ")
		fmt.Scan(&p)
			if (p=="1") {
				sortnama(B, A, i)
				fmt.Scanln(&read)
			} else if (p=="2") {
				sortumur(B, A, i)
				fmt.Scanln(&read)
			}
	}
}

func list1(B datap, A apart, i int) {
	
	//Prosedur

	var p string
	
	clear()

	for (p != "3") {
		clear()
		fmt.Println("===================================")
		fmt.Println("| 1. Cari Menurut Kamar           |")
		fmt.Println("| 2. Cari Menurut Nama            |")
		fmt.Println("| 3. Kembali                      |")
		fmt.Println("===================================\n")
		fmt.Print("Pilih Menu : ")
		fmt.Scan(&p)
			if (p=="1") {
				carikamar(B, A, i)
				fmt.Scanln(&read)
			} else if (p=="2") {
				carinama(B, A, i)
				fmt.Scanln(&read)
			}
	}
}
func pilihmenu(A apart, i int) {
	
	var user string

	for user != "7" {
		clear()
		fmt.Println("===================================")
		fmt.Println("| Pilihan Menu                    |")
		fmt.Println("|                                 |")
		fmt.Println("| 1. Booking Kamar                |")
		fmt.Println("| 2. Cek Harga Wisma              |")
		fmt.Println("| 3. Edit data pemesan            |")
		fmt.Println("| 4. Daftar Nama Pemesan          |")
		fmt.Println("| 5. Cari dan Lihat data pemesan  |")
		fmt.Println("| 6. Checkout                     |")
		fmt.Println("| 7. Exit                         |")
		fmt.Println("===================================")
		fmt.Println("")
		fmt.Print("Pilih Menu :")
		fmt.Scan(&user)
		if user == "1" {
			inputdata(&B, i)
			bookingapart(&A, i)
			i = i + 1
			q = q + 1
		} else if user == "2" {
			hargaapart()
		} else if user == "3" {
			editdata(&B, i)
		} else if user == "4" {
			list(B, A, i)
		} else if user == "5" {
			list1(B, A, i)
		} else if user == "6" {
			checkout(&B,&A,i)
			time.Sleep(time.Second * 1)
		} else {
			clear()
			fmt.Println("=======================")
			fmt.Println("|    TERIMA KASIH     |")
			fmt.Println("|  SAMPAI JUMPA LAGI  |")
			fmt.Println("=======================")
			fmt.Scanln(&read)
		}
	}

}

func main() {
	clear()
	i = 0
	q = 0
	pilihmenu(A, i)
	fmt.Scanln(&read)
	clear()
}