package main

import "fmt"

const NMAX int = 1000

type KPU struct {
	date            string
	pemilu          arrPemilu
	n, ambang_batas int
}
type arrPemilu [NMAX]calon
type calon struct {
	nama_ketua_calon string
	nama_wakil_calon string
	partai           string
	nomor_urut       string
	suara            int
}

type Pemilih struct {
	guest arrPeserta
	n     int
}
type arrPeserta [NMAX]Peserta
type Peserta struct {
	nama, email, pass, nomor string
	Nama_calon               string
	StatusPemilihan          bool
}

func main() {
	header()
	menu_start()
	menu_login()
}

func header() {
	fmt.Println("* ------------------------------------------- *")
	fmt.Println("*           Aplikasi Pemilihan umum           *")
	fmt.Println("*                  Created by                 *")
	fmt.Println("*                Yusa and Raisya              *")
	fmt.Println("*          Algoritma Pemrograman 2023         *")
	fmt.Println("* ------------------------------------------- *")
}

func menu_login() {
	var s bool
	var kandidat KPU
	var pengguna Pemilih
	var user, date string

	kandidat.date = "10/10/2010"
	kandidat.n = 0
	pengguna.n = 0
	kandidat.ambang_batas = 5
	s = true
	for s {
		fmt.Println("Login: ")
		fmt.Scan(&user)
		if user == "Admin" {
			main_menu_KPU(&kandidat)
		} else if user == "Peserta" {
			main_menu_peserta(&pengguna, &kandidat, date)
		} else if user == "Tanggal" {
			change_date(&date)
		} else if user == "Keluar" {
			s = false
		} else {
			fmt.Println("Mohon beri inputan yang valid")
		}
	}
	fmt.Println("Terima kasih")

}

func change_date(date *string) {
	fmt.Print("Masukkan tanggal baru: ")
	fmt.Scan(&*date)
}

func menu_start() {

	fmt.Println("Admin    Peserta")

}

func main_menu_KPU(pemilu *KPU) {
	var opsi string = ""
	var kondisi int
	for opsi != "10" {

		fmt.Println("* Menu Utama *")
		fmt.Println("1. Tambah nama calon")
		fmt.Println("2. Ubah nama calon")
		fmt.Println("3. Hapus nama calon")
		fmt.Println("4. Pencarian nama calon")
		fmt.Println("5. Tampil4kan data terurut atau tidak")
		fmt.Println("6. Tanggal pemilihan")
		fmt.Println("7. Nilai ambang batas")
		fmt.Println("8. Menampilkan calon yang melewati ambang batas")
		fmt.Println("9. Mencari nama calon berdasarkan no urut dan partai")
		fmt.Println("10. Keluar")
		fmt.Print("Pilih Opsi: ")
		fmt.Scan(&opsi)
		if opsi == "1" {
			// tambah()
			TambahCalon(pemilu)
		} else if opsi == "2" {
			// ubah()
			fmt.Println("Ubah calon: ")
			edit(pemilu, kondisi)
		} else if opsi == "3" {
			// hapus()
			fmt.Print("Hapus calon: ")
			fmt.Scan(&kondisi)
			hapus(pemilu, kondisi)
		} else if opsi == "4" {
			// cari()
			fmt.Print("Tulis nama calon:")
			searchPrint_ketua(*pemilu)

		} else if opsi == "5" {
			// print_data()
			sorting(pemilu)
		} else if opsi == "6" {
			// Lihat tanggal pemilihan
			election_date(&pemilu.date)
		} else if opsi == "7" {
			threshold(pemilu)
		} else if opsi == "8" {
			melampaui_batas(*pemilu)
		} else if opsi == "9" {
			cari_calon(*pemilu)
		}
	}
}

func TambahCalon(p *KPU) {
	var temp calon
	fmt.Print("Nama ketua calon: ")
	fmt.Scan(&temp.nama_ketua_calon)
	fmt.Print("Nama wakil calon: ")
	fmt.Scan(&temp.nama_wakil_calon)
	fmt.Print("Nama partai: ")
	fmt.Scan(&temp.partai)
	fmt.Print("Nomor urut calon: ")
	fmt.Scan(&temp.nomor_urut)
	p.n += 1
	p.pemilu[p.n] = temp
}

func edit(p *KPU, i int) {
	var opsi string
	var ganti string
	fmt.Println("1. Mengubah nama ketua")
	fmt.Println("2. Mengubah nama wakil")
	fmt.Println("3. Mengubah partai")
	fmt.Println("4. Mengubah nomor urut calon")
	fmt.Print("Opsi Anda: ")
	fmt.Scan(&opsi)
	if opsi == "1" {
		fmt.Print("Nama calon ketua baru: ")
		fmt.Scan(&ganti)
		p.pemilu[i].nama_ketua_calon = ganti
	} else if opsi == "2" {
		fmt.Print("Nama calon wakil baru: ")
		fmt.Scan(&ganti)
		p.pemilu[i].nama_wakil_calon = ganti
	} else if opsi == "3" {
		fmt.Print("Nama partai baru: ")
		fmt.Scan(&ganti)
		p.pemilu[i].partai = ganti
	} else if opsi == "4" {
		fmt.Print("Nomor urut calon: ")
		fmt.Scan(&ganti)
		p.pemilu[i].nomor_urut = ganti
	}
}

func hapus(p *KPU, i int) {
	for i < p.n {
		p.pemilu[i] = p.pemilu[i+1]
	}
	p.n -= 1
}

func searchPrint_ketua(p KPU) {
	var i int
	var ketua string

	fmt.Print("Tuliskan nama calon:")
	fmt.Scan(&ketua)
	for i = 0; i < p.n; i++ {
		if p.pemilu[i].nama_ketua_calon == ketua {
			fmt.Println(p.pemilu[i].nama_ketua_calon)
			fmt.Println(p.pemilu[i].nama_wakil_calon)
			fmt.Println(p.pemilu[i].nomor_urut)
			fmt.Println(p.pemilu[i].partai)
			fmt.Println(p.pemilu[i].suara)
		}
	}
}

func election_date(date *string) {
	var opsi string

	fmt.Println("Tanggal pemilihan:", *date)
	fmt.Println("Apakah Anda ingin mengubah tanggal pemilihan? (y/t)")
	fmt.Print("Pilihan Anda: ")
	fmt.Scan(&opsi)
	if opsi == "y" {
		fmt.Print("Masukkan tanggal pemilihan baru (format: dd/mm/yyyy): ")
		fmt.Scan(&*date)
		fmt.Println("Tanggal pemilihan berhasil diubah")
	} else {
		fmt.Println("Tanggal pemilihan batall diubah")
	}
}

func threshold(pemilu *KPU) {
	var opsi string

	fmt.Println("------------------------------------")
	fmt.Println("Ambang batas:", pemilu.ambang_batas)
	fmt.Println("Apakah Anda ingin mengganti? (y/t)")
	fmt.Println("")
	fmt.Print("Jawaban Anda: ")
	fmt.Scan(&opsi)
	if opsi == "y" {
		fmt.Print("Ambang batas baru: ")
		fmt.Scan(&pemilu.ambang_batas)
		fmt.Println("Ambang batas berhasil diganti")
	} else {
		fmt.Println("Ambang batas batal diganti")
	}
}

func melampaui_batas(p KPU) {
	var i int
	for i = 0; i < p.n; i++ {
		if p.pemilu[i].suara >= p.ambang_batas {
			fmt.Println(p.pemilu[i].nama_ketua_calon, p.pemilu[i].nama_wakil_calon)
			fmt.Println(p.pemilu[i].partai, p.pemilu[i].suara)
			fmt.Println("-----------------------------------------------")
		}
	}

}

func sorting(p *KPU) {
	var pass, i int
	var temp calon
	var opsi string
	pass = 1
	for pass < p.n {
		i = pass
		temp = p.pemilu[pass]
		for i > 0 && temp.suara > p.pemilu[i-1].suara {
			p.pemilu[i] = p.pemilu[i-1]
			i--
		}
		p.pemilu[i] = temp
		pass++
	}
	fmt.Println("Apakah ingin ditampilkan? (y/t)")
	fmt.Print("Jawaban Anda: ")
	fmt.Scan(&opsi)
	if opsi == "y" {
		print_sorting(*p)
	}
}

func print_sorting(p KPU) {
	var i int
	for i = 0; i < p.n; i++ {
		fmt.Println(p.pemilu[i].nama_ketua_calon, p.pemilu[i].nama_wakil_calon, p.pemilu[i].nomor_urut, p.pemilu[i].partai, p.pemilu[i].suara)
		fmt.Println()
	}
}

func main_menu_peserta(pengguna *Pemilih, kandidat *KPU, date string) {
	var pilihan string

	fmt.Println("------------------------------------")
	fmt.Println("~~~~~   Selamat datang   ~~~~~")
	fmt.Println("1. Daftar")
	fmt.Println("2. Masuk")
	fmt.Println("3. Keluar")
	fmt.Println("------------------------------------")
	fmt.Print("Pilihan Anda: ")
	fmt.Scan(&pilihan)
	if pilihan == "1" {
		registration_peserta(pengguna)
	} else if pilihan == "2" {
		login_peserta(pengguna, kandidat, date)
	}
	if pilihan != "3" {
		main_menu_peserta(pengguna, kandidat, date)
	}
}

func registration_peserta(pengguna *Pemilih) {
	var name, email, nomor, pass string

	fmt.Println("------------------------------------")
	fmt.Println("Masukkan 'kembali' jika tidak ingin mendaftar")
	fmt.Print("Masukkan nama: ")
	fmt.Scan(&name)
	if name != "kembali" {
		fmt.Print("Masukkan email: ")
		fmt.Scan(&email)
		fmt.Print("Masukkan password: ")
		fmt.Scan(&pass)
		fmt.Print("Masukkan nomor ponsel: ")
		fmt.Scan(&nomor)
		pengguna.n += 1
		pengguna.guest[pengguna.n].nama = name
		pengguna.guest[pengguna.n].email = email
		pengguna.guest[pengguna.n].pass = pass
		pengguna.guest[pengguna.n].nomor = nomor
		fmt.Println("------------------------------------")
		fmt.Println("Akun berhasil dibuat")
	} else {
		fmt.Println("------------------------------------")
		fmt.Println("Akun batal dibuat")
	}
}

func login_peserta(p *Pemilih, kandidat *KPU, date string) {
	var email, pass string
	var status, pilihan int
	var kondisi bool

	kondisi = true
	fmt.Println("------------------------------------")
	fmt.Println("Masukkan 'kembali' jika tidak ingin masuk")
	fmt.Print("Masukkan email: ")
	fmt.Scan(&email)
	if email != "kembali" {
		fmt.Print("Masukkan password: ")
		fmt.Scan(&pass)
		status = email_password(*p, email, pass)
		for status == -1 && email != "kembali" {
			fmt.Println("------------------------------------")
			fmt.Println("Maaf, username atau password Anda kurang tepat")
			fmt.Println("Masukkan 'kembali' jika tidak ingin masuk")
			fmt.Print("Masukkan email: ")
			fmt.Scan(&email)
			if email != "kembali" {
				fmt.Print("Masukkan password: ")
				fmt.Scan(&pass)
				status = email_password(*p, email, pass)
			} else {
				// kondisi = false
			}
		}
		if kondisi {
			pilihan = 0
			menu_Peserta(&p.guest[status], kandidat, &pilihan, date)
		}
	}
}

func email_password(p Pemilih, email, kunci string) int {
	var i, status int

	i = 1
	status = -1
	for i <= p.n && status == -1 {
		if p.guest[i].email == email && p.guest[i].pass == kunci {
			status = i
		}
		i += 1
	}
	return status
}

func menu_Peserta(user *Peserta, kandidat *KPU, pilihan *int, date string) {
	var status int

	for status != 3 && status != 4 {
		fmt.Println("------------------------------------")
		fmt.Println("~~~   Selamat datang,", user.nama, "   ~~~")
		fmt.Println("1. Melihat Calon")
		fmt.Println("2. Logout")
		fmt.Println("------------------------------------")
		fmt.Print("Pilihan Anda (1,2): ")
		fmt.Scan(&status)
		fmt.Println("------------------------------------")
		if status == 1 {
			pemilihan_calon(user, kandidat, date)
		}
	}
}

func pemilihan_calon(user *Peserta, kandidat *KPU, date string) {
	var i int
	var opsi string
	var pemilu 

	for i = 1; i <= kandidat.n; i++ {
		fmt.Println("------------------------------------")
		fmt.Println("Nama ketua calon:	", kandidat.pemilu[i].nama_ketua_calon)
		fmt.Println("Nama wakil calon:	", kandidat.pemilu[i].nama_wakil_calon)
		fmt.Println("Nama partai:		", kandidat.pemilu[i].partai)
		fmt.Println("Nomor urut:		", kandidat.pemilu[i].nomor_urut)
	}
	fmt.Println("------------------------------------")
	if date == kandidat.date && !user.StatusPemilihan {
		fmt.Println("Apakah", user.nama, "ingin memilih? (y/t)")
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&opsi)
		if opsi == "y" {
			pemilihan_peserta(user, kandidat)
		}
	} else if user.StatusPemilihan {
		fmt.Println("Maaf, Anda sudah memilih")
	} else {
		fmt.Println("Maaf, Anda tidak dapat memilih")
	}
	fmt.Println("------------------------------------")
	fmt.Println("Masukkan apa saja untuk kembali ke menu")
	fmt.Scan(&opsi)
}

func pemilihan_peserta(user *Peserta, kandidat *KPU) {
	var pilihan string

	fmt.Print("Pilih pasangan calon nomor urut: ")
	fmt.Scan(&pilihan)
	user.StatusPemilihan = true
	voting_kandidat(kandidat, pilihan)
}

func voting_kandidat(kandidat *KPU, nomor string) {
	var i int
	var status bool

	status = true
	for i = 1; i <= kandidat.n && status; i++ {
		if kandidat.pemilu[i].nomor_urut == nomor {
			kandidat.pemilu[i].suara += 1
			status = false
		}
	}
}

func cari_calon(p KPU) {
	var opsi1, opsi2 string
	fmt.Scan(&opsi1, &opsi2)
	searchPrint_namaCalon(p, opsi1, opsi2)
}

func searchPrint_namaCalon(p KPU, opsi1, opsi2 string) {
	var i int
	for i = 0; i < p.n; i++ {
		if p.pemilu[i].partai == opsi1 && p.pemilu[i].nama_ketua_calon == opsi2 {
			fmt.Println(p.pemilu[i].nama_ketua_calon)
			fmt.Println(p.pemilu[i].nama_wakil_calon)
			fmt.Println(p.pemilu[i].nomor_urut)
			fmt.Println(p.pemilu[i].partai)
			fmt.Println(p.pemilu[i].suara)
		}
	}
}
