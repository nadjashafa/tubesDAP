package main

import "fmt"


// K A M U S
type spec struct {
	tipe, ram, memory string
	harga int
}



const arr = 1000

var i, j, k int
var merkkk arrmerk

type arrmerk struct {
	apple_arr[arr] spec
	samsung_arr[arr] spec
	sony_arr[arr] spec

}


func main() {
	var menu, merk string
	var totapple, totsamsung, totsony int

	

	fmt.Println(" _____________________________________ ")
	fmt.Println("|                                     |")
	fmt.Println("|          -----------------          |")
	fmt.Println("|            W E L C O M E            |")
	fmt.Println("|          -----------------          |")
	fmt.Println("|_____________________________________|")
	fmt.Println(" ")
	fmt.Println(" ")
	


	fmt.Println("1. Menu edit")
	fmt.Println("2. Daftar HP")
	fmt.Println("3. Daftar Stock")
	fmt.Println("4. Search by Spec")
	fmt.Println("5. Transaksi")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Print("Pilih Menu : ")
	fmt.Scanln(&menu)
	fmt.Println(" ")
	fmt.Println(" ")

	if menu == "1" {
		edit(&totapple, &totsamsung, &totsony)
		balik(&totapple, &totsamsung, &totsony)

	} else if menu == "2" {
		daftarhp(totapple, totsamsung, totsony)
		balik(&totapple, &totsamsung, &totsony) 

	} else if menu == "3" {
		daftarstock(totapple, totsamsung, totsony)
		balik(&totapple, &totsamsung, &totsony)

	} else if menu == "4" {
		fmt.Println("Masukan merk hp : ")
		fmt.Scanln(&merk)

		if merk == "apple" {
			cariapple(totapple)
		} else if merk == "samsung" {
			carisamsung(totsamsung)
		} else if merk == "sony" {
			carisony(totsony)
		} else {
			fmt.Println("Merk Tidak Tersedia")
		}

		balik(&totapple, &totsamsung, &totsony)

	} else if menu == "5" {
		transaksi(&totapple, &totsamsung, &totsony)
		balik(&totapple, &totsamsung, &totsony)

	} else {
		fmt.Println("[ MENU TIDAK TERSEDIA ]")
		fmt.Println(" ")
		balik(&totapple, &totsamsung, &totsony)
	}


	fmt.Println(" _____________________________________ ")
	fmt.Println("|                                     |")
	fmt.Println("|          -----------------          |")
	fmt.Println("|            S E L E S A I            |")
	fmt.Println("|          -----------------          |")
	fmt.Println("|_____________________________________|")
	fmt.Println(" ")
	fmt.Println(" ")
}

func balik (totapple, totsamsung, totsony *int) {
	var balik string
	var menu, merk string

	fmt.Println("kembali/selesai? ")
	fmt.Scan(&balik)
	fmt.Println(" ")
	fmt.Println(" ")

	for balik == "kembali" {

		fmt.Println("1. Menu edit")
		fmt.Println("2. Daftar HP")
		fmt.Println("3. Daftar Stock")
		fmt.Println("4. Search by Spec")
		fmt.Println("5. Transaksi")
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Print("Pilih Menu : ")
		fmt.Scanln(&menu)
		fmt.Println(" ")
		fmt.Println(" ")

		if menu == "1" {
			edit(&*totapple, &*totsamsung, &*totsony)

		} else if menu == "2" {
			daftarhp(*totapple, *totsamsung, *totsony) 

		} else if menu == "3" {
			daftarstock(*totapple, *totsamsung, *totsony)

		} else if menu == "4" {
			fmt.Println("Masukan merk hp : ")
			fmt.Scanln(&merk)

			if merk == "apple" {
				cariapple(*totapple)
			} else if merk == "samsung" {
				carisamsung(*totsamsung)
			} else if merk == "sony" {
				carisony(*totsony)
			} else {
				fmt.Println("Merk Tidak Tersedia")
			}

		} else if menu == "5" {
			transaksi(totapple, totsamsung, totsony)
		} else {
			fmt.Println("[ MENU TIDAK TERSEDIA ]")
			fmt.Println(" ")
		}

		fmt.Println("kembali/selesai? ")
		fmt.Scan(&balik)
		fmt.Println(" ")
		fmt.Println(" ")

	}

}

func edit (totapple, totsamsung, totsony *int) {
	var pilihan, pilihan2 string


	fmt.Println("[ Pilih menu edit : ]")
	fmt.Println(" ------------------- ")
	fmt.Println(" ")
	fmt.Println("1. Masukan Data")
	fmt.Println("2. Hapus Data")
	fmt.Println(" ")
	fmt.Print("Pilih menu : ")
	fmt.Scanln(&pilihan)
	fmt.Println(" ")
	fmt.Println(" ")

	if pilihan == "1" {
		masukan(&*totapple, &*totsamsung, &*totsony)
			
	} else if pilihan == "2" {
		hapus(totapple, totsamsung, totsony)
		fmt.Print("Hapus data lagi ? (Y/N) ")
		fmt.Scanln(&pilihan2)
		for (pilihan2 == "Y") || (pilihan2 == "y") {
			hapus(totapple, totsamsung, totsony)
			fmt.Print("Hapus data lagi ? (Y/N) ")
			fmt.Scanln(&pilihan2)
		}

	} else {
		fmt.Println("Menu tidak tersedia")
	}

} 

func daftarhp (totapple, totsamsung, totsony int) {
	var i int
	

	fmt.Println(" _______________ ")
	fmt.Println("|               |")
	fmt.Println("|   A P P L E   |")
	fmt.Println("|_______________|")
	fmt.Println("[ Tipe - RAM -  Memory - Harga ]")
	fmt.Println(" ")
	i = 0
	for i < totapple {
		fmt.Println("- ", merkkk.apple_arr[i].tipe)
		fmt.Println("  ", merkkk.apple_arr[i].ram)
		fmt.Println("  ", merkkk.apple_arr[i].memory)
		fmt.Println("  ", merkkk.apple_arr[i].harga)
		fmt.Println(" ")
		i++
	} 
	

	fmt.Println(" ___________________ ")
	fmt.Println("|                   |")
	fmt.Println("|   S A M S U N G   |")
	fmt.Println("|___________________|")
	fmt.Println("[ Tipe - RAM -  Memory - Harga ]")
	fmt.Println(" ")
	i = 0
	for i < totsamsung {
		fmt.Println("- ", merkkk.samsung_arr[i].tipe)
		fmt.Println("  ", merkkk.samsung_arr[i].ram)
		fmt.Println("  ", merkkk.samsung_arr[i].memory)
		fmt.Println("  ", merkkk.samsung_arr[i].harga)
		fmt.Println(" ")
		i++
	}


	fmt.Println(" _____________ ")
	fmt.Println("|             |")
	fmt.Println("|   S O N Y   |")
	fmt.Println("|_____________|")
	fmt.Println("[ Tipe - RAM -  Memory - Harga ]")
	fmt.Println(" ")
	i = 0
	for i < totsony {
		fmt.Println("- ", merkkk.sony_arr[i].tipe)
		fmt.Println("  ", merkkk.sony_arr[i].ram)
		fmt.Println("  ", merkkk.sony_arr[i].memory)
		fmt.Println("  ", merkkk.sony_arr[i].harga)
		fmt.Println(" ")
		i++
	}


}

func daftarstock (totapple, totsamsung, totsony int) {
	var temp1, i int
	var m1, m2, m3, temp2 string

	fmt.Println(" _______________________ ")
	fmt.Println("|                       |")
	fmt.Println("|      DAFTAR STOCK     |")
	fmt.Println("|_______________________|")
	fmt.Println(" ")
	fmt.Println(" ")

	m1 = "APPLE"
	m2 = "SAMSUNG"
	m3 = "SONY"

	i = 0
	for i < 3 {
		if totapple < totsamsung {
			temp1 = totsamsung
			totsamsung = totapple
			totapple = temp1

			temp2 = m2
			m2 = m1
			m1 = temp2
		}

		if totsamsung < totsony {
			temp1 = totsony
			totsony = totsamsung
			totsamsung = temp1

			temp2 = m3
			m3 = m2
			m2 = temp2
		}

		i++

	}

	fmt.Println(" STOCK ", m1)
	fmt.Println("= ", totapple)
	fmt.Println(" ")
	fmt.Println(" STOCK ", m2)
	fmt.Println("= ", totsamsung)
	fmt.Println(" ")
	fmt.Println(" STOCK ", m3)
	fmt.Println("= ", totsony)
	fmt.Println(" ")

}

// ISI FUNC EDIT
func hp_apple (totapple *int) {
	var apple spec
	var lanjut string
	

	fmt.Println(" ")
	fmt.Print("Masukan Tipe HP : ")
	fmt.Scanln(&apple.tipe)
	fmt.Print("Masukan Tipe RAM : ")
	fmt.Scanln(&apple.ram)
	fmt.Print("Masukan Tipe Memory : ")
	fmt.Scanln(&apple.memory)
	fmt.Print("Masukan Harga : ")
	fmt.Scanln(&apple.harga)

	merkkk.apple_arr[0].tipe = apple.tipe
	merkkk.apple_arr[0].ram = apple.ram
	merkkk.apple_arr[0].memory = apple.memory
	merkkk.apple_arr[0].harga = apple.harga

	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Print("Tambah data lagi (Y/N) ? ")
	fmt.Scanln(&lanjut)

	i = 1
	*totapple = 1
	for (i < arr) && (lanjut == "Y" || lanjut == "y") {
		
		fmt.Println(" ")
		fmt.Print("Masukan Tipe HP : ")
		fmt.Scanln(&apple.tipe)
		fmt.Print("Masukan Tipe RAM : ")
		fmt.Scanln(&apple.ram)
		fmt.Print("Masukan Tipe Memory : ")
		fmt.Scanln(&apple.memory)
		fmt.Print("Masukan Harga : ")
		fmt.Scanln(&apple.harga)

		merkkk.apple_arr[i].tipe = apple.tipe
		merkkk.apple_arr[i].ram = apple.ram
		merkkk.apple_arr[i].memory = apple.memory
		merkkk.apple_arr[i].harga = apple.harga

		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Print("Tambah data lagi (Y/N) ? ")
		fmt.Scanln(&lanjut)

		i = i + 1
		*totapple = *totapple + 1
	
	}

}

func hp_samsung (totsamsung *int) {
	var samsung spec
	var lanjut string
	

	fmt.Println(" ")
	fmt.Print("Masukan Tipe HP : ")
	fmt.Scanln(&samsung.tipe)
	fmt.Print("Masukan Tipe RAM : ")
	fmt.Scanln(&samsung.ram)
	fmt.Print("Masukan Tipe Memory : ")
	fmt.Scanln(&samsung.memory)
	fmt.Print("Masukan Harga : ")
	fmt.Scanln(&samsung.harga)

	merkkk.samsung_arr[0].tipe = samsung.tipe
	merkkk.samsung_arr[0].ram = samsung.ram
	merkkk.samsung_arr[0].memory = samsung.memory
	merkkk.samsung_arr[0].harga = samsung.harga

	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Print("Tambah data lagi (Y/N) ? ")
	fmt.Scanln(&lanjut)
	

	j = 1
	*totsamsung = 1
	for (j < arr) && (lanjut == "Y" || lanjut == "y") {	
	
		fmt.Println(" ")
		fmt.Print("Masukan Tipe HP : ")
		fmt.Scanln(&samsung.tipe)
		fmt.Print("Masukan Tipe RAM : ")
		fmt.Scanln(&samsung.ram)
		fmt.Print("Masukan Tipe Memory : ")
		fmt.Scanln(&samsung.memory)
		fmt.Print("Masukan Harga : ")
		fmt.Scanln(&samsung.harga)

		merkkk.samsung_arr[j].tipe = samsung.tipe
		merkkk.samsung_arr[j].ram = samsung.ram
		merkkk.samsung_arr[j].memory = samsung.memory
		merkkk.samsung_arr[j].harga = samsung.harga
		
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Print("Tambah data lagi (Y/N) ? ")
		fmt.Scanln(&lanjut)		

		j = j + 1
		*totsamsung = *totsamsung + 1
	}

}

func hp_sony (totsony *int) {
	var sony spec
	var lanjut string
	

	fmt.Println(" ")
	fmt.Print("Masukan Tipe HP : ")
	fmt.Scanln(&sony.tipe)
	fmt.Print("Masukan Tipe RAM : ")
	fmt.Scanln(&sony.ram)
	fmt.Print("Masukan Tipe Memory : ")
	fmt.Scanln(&sony.memory)
	fmt.Print("Masukan Harga : ")
	fmt.Scanln(&sony.harga)

	merkkk.sony_arr[0].tipe = sony.tipe
	merkkk.sony_arr[0].ram = sony.ram
	merkkk.sony_arr[0].memory = sony.memory
	merkkk.sony_arr[0].harga = sony.harga

	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Print("Tambah data lagi (Y/N) ? ")
	fmt.Scanln(&lanjut)		

	
	k = 1
	*totsony = 1
	for (k < arr) && (lanjut == "Y" || lanjut == "y") {
	
		fmt.Println(" ")
		fmt.Print("Masukan Tipe HP : ")
		fmt.Scanln(&sony.tipe)
		fmt.Print("Masukan Tipe RAM : ")
		fmt.Scanln(&sony.ram)
		fmt.Print("Masukan Tipe Memory : ")
		fmt.Scanln(&sony.memory)
		fmt.Print("Masukan Harga : ")
		fmt.Scanln(&sony.harga)

		merkkk.sony_arr[k].tipe = sony.tipe
		merkkk.sony_arr[k].ram = sony.ram
		merkkk.sony_arr[k].memory = sony.memory
		merkkk.sony_arr[k].harga = sony.harga	

		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Print("Tambah data lagi (Y/N) ? ")
		fmt.Scanln(&lanjut)		

		k = k + 1
		*totsony = *totsony + 1
			
	}

}

func masukan (totapple, totsamsung, totsony *int) {
	var merk string

	fmt.Println("--- Masukan Merk HP ---")
	fmt.Scanln(&merk)

	if merk == "apple" {
		hp_apple(&*totapple)

	} else if merk == "samsung" {
		hp_samsung(&*totsamsung)
	} else if merk == "sony" {
		hp_sony(&*totsony)
	} else {
		fmt.Println("Merk saat ini belum tersedia")
	}

}

func hapus (totapple, totsamsung, totsony *int) {
	var pilih int
	var merk string
	

	fmt.Print("Pilih merk hp : ")
	fmt.Scanln(&merk)
	if merk == "apple" {
		
		fmt.Println(" _______________ ")
		fmt.Println("|               |")
		fmt.Println("|   A P P L E   |")
		fmt.Println("|_______________|")
		fmt.Println("[ Tipe - RAM -  Memory - Harga ]")
		fmt.Println(" ")
		i = 0
		for i < *totapple {
			fmt.Println("- ", merkkk.apple_arr[i].tipe)
			fmt.Println("  ", merkkk.apple_arr[i].ram)
			fmt.Println("  ", merkkk.apple_arr[i].memory)
			fmt.Println("  ", merkkk.apple_arr[i].harga)
			fmt.Println(" ")
			i++
		} 

		fmt.Println(" ")
		fmt.Print("Pilih data yang akan dihapus ")
		fmt.Scanln(&pilih)

		if pilih == 0 {

			fmt.Println("masukan tidak valid")
		} else {

			pilih = pilih - 1
			merkkk.apple_arr[pilih].tipe = " "
			merkkk.apple_arr[pilih].memory = " "
			merkkk.apple_arr[pilih].ram = " "
			merkkk.apple_arr[pilih].harga = 0
			*totapple--

			fmt.Println("Data berhasil dihapus")
			fmt.Println(" ")
		}
	} else if merk == "samsung" {
		fmt.Println(" ___________________ ")
		fmt.Println("|                   |")
		fmt.Println("|   S A M S U N G   |")
		fmt.Println("|___________________|")
		fmt.Println("[ Tipe - RAM -  Memory - Harga ]")
		fmt.Println(" ")
		i = 0
		for i < *totsamsung {
			fmt.Println("- ", merkkk.samsung_arr[i].tipe)
			fmt.Println("  ", merkkk.samsung_arr[i].ram)
			fmt.Println("  ", merkkk.samsung_arr[i].memory)
			fmt.Println("  ", merkkk.samsung_arr[i].harga)
			fmt.Println(" ")
			i++
		}

		fmt.Println(" ")
		fmt.Print("Pilih data yang akan dihapus ")
		fmt.Scanln(&pilih)
		if pilih == 0 {
			fmt.Println("masukan tidak valid")
		} else {
			pilih = pilih - 1
			merkkk.samsung_arr[pilih].tipe = " "
			merkkk.samsung_arr[pilih].memory = " "
			merkkk.samsung_arr[pilih].ram = " "
			merkkk.samsung_arr[pilih].harga = 0
			*totsamsung--
			fmt.Println("Data berhasil dihapus")
			fmt.Println(" ")
		}
	} else if merk == "sony" {
		fmt.Println(" _____________ ")
		fmt.Println("|             |")
		fmt.Println("|   S O N Y   |")
		fmt.Println("|_____________|")
		fmt.Println("[ Tipe - RAM -  Memory - Harga ]")
		fmt.Println(" ")
		i = 0
		for i < *totsony {
			fmt.Println("  ", merkkk.sony_arr[i].tipe)
			fmt.Println("  ", merkkk.sony_arr[i].ram)
			fmt.Println("  ", merkkk.sony_arr[i].memory)
			fmt.Println("  ", merkkk.sony_arr[i].harga)
			fmt.Println(" ")
			i++
		}

		fmt.Println(" ")
		fmt.Print("Pilih data yang akan dihapus ")
		fmt.Scanln(&pilih)
		if pilih == 0 {
			fmt.Println("masukan tidak valid")
		} else {
			pilih = pilih - 1
			merkkk.sony_arr[pilih].tipe = " "
			merkkk.sony_arr[pilih].memory = " "
			merkkk.sony_arr[pilih].ram = " "
			merkkk.sony_arr[pilih].harga = 0
			*totsony--
			fmt.Println("Data berhasil dihapus")
			fmt.Println(" ")
		}
	} else {
		fmt.Println("merk tidak tersedia")
	}

}

func cariapple(totapple int) {
	var spek, ram, memory string
	

	fmt.Print("Masukan Spesifikasi yang Diinginkan : ")
	fmt.Scanln(&spek)

	if spek == "RAM" || spek == "ram" {
		fmt.Print("RAM : ")
		fmt.Scanln(&ram)
		fmt.Println(" ")
		i = 0
		for i < totapple {
			if merkkk.apple_arr[i].ram == ram {
				fmt.Println("Tipe  : iPhone ", merkkk.apple_arr[i].tipe)
				fmt.Println("Harga : Rp.", merkkk.apple_arr[i].harga)
				fmt.Println(" ")

			} 

			i++
		}
	} else if spek == "Memory" || spek == "memory" {
		fmt.Print("Memory : ")
		fmt.Scanln(&memory)
		fmt.Println(" ")
		i = 0
		for i < totapple {
			if merkkk.apple_arr[i].memory == memory {
				fmt.Println("Tipe  : iPhone ", merkkk.apple_arr[i].tipe)
				fmt.Println("Harga : Rp.", merkkk.apple_arr[i].harga)
				fmt.Println(" ")

			} 

			i++
		}

	} else {
		fmt.Println("--- Menu Tidak Tersedia ---")
	}
	
	 

}




func carisamsung(totsamsung int) {
	var spek, ram, memory string
	

	fmt.Print("Masukan Spesifikasi yang Diinginkan : ")
	fmt.Scanln(&spek)

	if spek == "RAM" || spek == "ram" {
		fmt.Print("RAM : ")
		fmt.Scanln(&ram)
		fmt.Println(" ")
		i = 0
		for i < totsamsung {
			if merkkk.samsung_arr[i].ram == ram {
				fmt.Println("Tipe  : Samsung ", merkkk.samsung_arr[i].tipe)
				fmt.Println("Harga : Rp.", merkkk.samsung_arr[i].harga)
				fmt.Println(" ")

			} 

			i++
		}
	} else if spek == "Memory" || spek == "memory" {
		fmt.Print("Memory : ")
		fmt.Scanln(&memory)
		fmt.Println(" ")
		i = 0
		for i < totsamsung {
			if merkkk.samsung_arr[i].memory == memory {
				fmt.Println("Tipe  : Samsung ", merkkk.samsung_arr[i].tipe)
				fmt.Println("Harga : Rp.", merkkk.samsung_arr[i].harga)
				fmt.Println(" ")

			} 

			i++
		}

	} else {
		fmt.Println("--- Menu Tidak Tersedia ---")
	}
	
	 

}



func carisony(totsony int) {
	var spek, ram, memory string
	

	fmt.Print("Masukan Spesifikasi yang Diinginkan : ")
	fmt.Scanln(&spek)

	if spek == "RAM" || spek == "ram" {
		fmt.Print("RAM : ")
		fmt.Scanln(&ram)
		fmt.Println(" ")
		i = 0
		for i < totsony {
			if merkkk.sony_arr[i].ram == ram {
				fmt.Println("Tipe  : Sony ", merkkk.sony_arr[i].tipe)
				fmt.Println("Harga : Rp.", merkkk.sony_arr[i].harga)
				fmt.Println(" ")

			} 

			i++
		}
	} else if spek == "Memory" || spek == "memory" {
		fmt.Print("Memory : ")
		fmt.Scanln(&memory)
		fmt.Println(" ")
		i = 0
		for i < totsony {
			if merkkk.sony_arr[i].memory == memory {
				fmt.Println("Tipe  : Sony", merkkk.sony_arr[i].tipe)
				fmt.Println("Harga : Rp.", merkkk.sony_arr[i].harga)
				fmt.Println(" ")

			} 

			i++
		}

	} else {
		fmt.Println("--- Menu Tidak Tersedia ---")
	}
	
	 

}


func transaksi (totapple, totsamsung, totsony *int) {
	var merk, kembali string
	var pilbar, total int
	
	fmt.Println("===============================================================")
	fmt.Println("==================== RINCIAN TRANSAKSI ========================")
	fmt.Println("===============================================================")
	fmt.Println(" ")
	fmt.Println(" ")
	total = 0
	kembali = "y"
	for kembali == "y" {
		i = 0
		j = 1
		fmt.Print("Masukan merk HP yang ingin di beli : ")
		fmt.Scanln(&merk)
		if merk == "samsung" {
			for i < *totsamsung {
				fmt.Println(j ,"tipe   : ", merkkk.samsung_arr[i].tipe)
				fmt.Println(" ","ram    : ", merkkk.samsung_arr[i].ram)
				fmt.Println(" ","memory : ", merkkk.samsung_arr[i].memory)
				fmt.Println(" ","harga  : ", merkkk.samsung_arr[i].harga)
				fmt.Println(" ")
				i++
				j++

			}	
		} else if merk == "apple" {
			for i < *totapple {
				fmt.Println(j ,"tipe   : ", merkkk.apple_arr[i].tipe)
				fmt.Println(" ","ram    : ", merkkk.apple_arr[i].ram)
				fmt.Println(" ","memory : ", merkkk.apple_arr[i].memory)
				fmt.Println(" ","harga  : ", merkkk.apple_arr[i].harga)
				fmt.Println(" ")
				i++
				j++
			}
		} else if merk == "sony" {
			for i < *totsony {
				fmt.Println(j ,"tipe    : ", merkkk.sony_arr[i].tipe)
				fmt.Println(" ","ram     : ", merkkk.sony_arr[i].ram)
				fmt.Println(" ","memory  : ", merkkk.sony_arr[i].memory)
				fmt.Println(" ","harga   : ", merkkk.sony_arr[i].harga)
				fmt.Println(" ")
				i++
				j++
			}
		} else {
			fmt.Print("Merk saat ini belum tersedia")
		}
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println("===============================================================")
		fmt.Println("===============================================================")
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Print("Pilih barang yang ingin dibeli : ")
		fmt.Scanln(&pilbar)
		if merk == "samsung" {
				fmt.Println("tipe   : ", merkkk.samsung_arr[pilbar-1].tipe)
				fmt.Println("ram    : ", merkkk.samsung_arr[pilbar-1].ram)
				fmt.Println("memory : ", merkkk.samsung_arr[pilbar-1].memory)
				fmt.Println("harga  : ", merkkk.samsung_arr[pilbar-1].harga)
				fmt.Println(" ")
				total = total + merkkk.samsung_arr[pilbar-1].harga
				*totsamsung--
		} else if merk == "apple" {
				fmt.Println("tipe   : ", merkkk.apple_arr[pilbar-1].tipe)
				fmt.Println("ram    : ", merkkk.apple_arr[pilbar-1].ram)
				fmt.Println("memory : ", merkkk.apple_arr[pilbar-1].memory)
				fmt.Println("harga  : ", merkkk.apple_arr[pilbar-1].harga)
				fmt.Println(" ")
				total = total + merkkk.apple_arr[pilbar-1].harga
				*totapple--
		} else if merk == "sony" {
				fmt.Println("tipe    :", merkkk.sony_arr[pilbar-1].tipe)
				fmt.Println("ram     :", merkkk.sony_arr[pilbar-1].ram)
				fmt.Println("memory  : ", merkkk.sony_arr[pilbar-1].memory)
				fmt.Println("harga   : ", merkkk.sony_arr[pilbar-1].harga)
				fmt.Println(" ")
				total = total + merkkk.sony_arr[pilbar-1].harga
				*totsony--
		}
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println("===============================================================")
		fmt.Println("===============================================================")
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Print("Apakah anda ingin melakukan transaksi lagi (Y/N) : ")
		fmt.Scan(&kembali)
	}
	fmt.Println(" ")
	fmt.Println("Total tagihan : ", total)
	fmt.Println(" ")
}


