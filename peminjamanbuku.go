package main

import "fmt"

//deklarasi variabel
var (
	Username = "Qheyla"
	Password = "2406356725"
	History  []string
)

type Pengguna struct {
	Nama   string
	Umur   int
	Alamat string
}

var pengguna = Pengguna{
	Nama:   "Qheyla Zavyera Valendro",
	Umur:   17,
	Alamat: "Cibubur, Kota Bogor, Jawa Barat",
}

var daftarBuku = []Buku{
	Buku{
		Nama:   "Pemrograman",
		Jumlah: 10,
	},
	Buku{
		Nama:   "Film",
		Jumlah: 5,
	},
	Buku{
		Nama:   "Printing",
		Jumlah: 20,
	},
}

type Buku struct {
	Nama   string
	Jumlah int
}

func main() {

	// Variabel Input
	var InputUsername, InputPassword string

	fmt.Println("<================================================>")
	fmt.Println("|     <3 SELAMAT DATANG DI PERPUSTAKAAN <3       |")
	fmt.Println("<================================================>")

	// Input Username
	fmt.Print("Masukkan Username Anda: ")
	fmt.Scanf("%s\n", &InputUsername)

	// Input Username
	fmt.Print("Masukkan Password Anda: ")
	fmt.Scanf("%s\n", &InputPassword)

	// Validasi Username & Password
	if InputUsername == Username && InputPassword == Password {
		fmt.Println("<================================================>")
		fmt.Println("|      <3 YEYYY AKHIRNYA LOGIN BERHASIL <3       |")
		fmt.Println("<================================================>")

		for {
			fmt.Println("<================================================>")
			fmt.Println("|       <3  MENU PROGRAM PERPUSTAKAAN  <3        |")
			fmt.Println("<================================================>")
			fmt.Println("|      1. Lihat Informasi Pengguna Program       |")
			fmt.Println("|      2. Lihat Daftar Buku                      |")
			fmt.Println("|      3. Tambah Daftar Buku                     |")
			fmt.Println("|      4. Tambah Peminjaman Buku                 |")
			fmt.Println("|      5. Histori Peminjaman Buku                |")
			fmt.Println("|      6. Keluar dari Program                    |")
			fmt.Println("<================================================>")

			// Input Pilihan Menu
			var pilihan int
			fmt.Print("AYOOO PILIH 1 MENU: ")
			fmt.Scanf("%d\n", &pilihan)

			switch pilihan {
			case 1:
				LihatInformasiPenggunaProgram()
			case 2:
				LihatDaftarBuku()
			case 3:
				TambahDaftarBuku()
			case 4:
				TambahPeminjamanBuku()
			case 5:
				HistoriPeminjamanBuku()
			case 6:
				fmt.Println("<================================================>")
				fmt.Println("| Terima Kasih telah Menggungunakan Program <333 |")
				fmt.Println("<================================================>")
				return
			default:
				fmt.Println("<================================================>")
				fmt.Println("|Yahhh Sangat Disayangkan, opsi anda tidak valid!|")
				fmt.Println("<================================================>")
			}
		}
	} else {
		fmt.Println("Password atau Username Anda Salah, Mohon untuk Coba Lagi!")
	}
}

func LihatInformasiPenggunaProgram() {
	fmt.Println("<================================================>")
	fmt.Println("|             Informasi Pengguna Program         |")
	fmt.Println("<================================================>")
	fmt.Printf("|Nama: %s\n", pengguna.Nama)
	fmt.Printf("|Umur: %d\n", pengguna.Umur)
	fmt.Printf("|Alamat: %s\n", pengguna.Alamat)
	fmt.Println("<================================================>")
}

func LihatDaftarBuku() {
	fmt.Println("<================================================>")
	fmt.Println("|              Daftar Buku Perpustakaan          |")
	fmt.Println("<================================================>")
	for i, buku := range daftarBuku {
		fmt.Printf("%d. Nama: %s - Jumlah: %d\n", i+1, buku.Nama, buku.Jumlah)
	}
	fmt.Println("<================================================>")
}

func TambahDaftarBuku() {
	var bukuBaru Buku
	fmt.Println("<================================================>")
	fmt.Println("|                Tambah Daftar Buku              |")
	fmt.Println("<================================================>")
	fmt.Print("|Masukkan Nama Buku: ")
	fmt.Scanln(&bukuBaru.Nama)
	fmt.Print("|Masukkan Jumlah Buku: ")
	fmt.Scanln(&bukuBaru.Jumlah)

	if bukuBaru.Jumlah <= 0 {
		fmt.Println("|  Maaf, Namun Jumlah harus lebih besar dari 0!  |")
		fmt.Println("<================================================>")
		return
	}

	daftarBuku = append(daftarBuku, bukuBaru)
	History = append(History, fmt.Sprintf("Menambahkan buku: %s - Jumlah: %d", bukuBaru.Nama, bukuBaru.Jumlah))
	fmt.Println("|       Buku berhasil ditambahkan YIPPIEEEEE     |")
	fmt.Println("<================================================>")
}
func TambahPeminjamanBuku() {
	var namaBuku string
	var jumlahPinjam int
	fmt.Println("<================================================>")
	fmt.Println("|              Tambah Peminjaman Buku            |")
	fmt.Println("<================================================>")
	fmt.Print("|Masukkan Nama Buku: ")
	fmt.Scanln(&namaBuku)
	fmt.Print("|Masukkan Jumlah yang ingin Dipinjam: ")
	fmt.Scanln(&jumlahPinjam)

	if jumlahPinjam <= 0 {
		fmt.Println("|  Maaf, Namun Jumlah harus lebih besar dari 0!  |")
		fmt.Println("<================================================>")
		return
	}

	for i, buku := range daftarBuku {
		if buku.Nama == namaBuku {
			if buku.Jumlah >= jumlahPinjam {
				daftarBuku[i].Jumlah -= jumlahPinjam
				History = append(History, fmt.Sprintf("Meminjam buku: %s - Jumlah: %d", namaBuku, jumlahPinjam))
				fmt.Println("| ASIK ASIKKK!!! Peminjaman Anda Berhasil YEY <3 |")
				fmt.Println("<================================================>")
				return
			} else {
				fmt.Println("<================================================>")
				fmt.Println("|  WADOHHH -_- Jumlah buku tidak mencukupi nihh! |")
				fmt.Println("<================================================>")
				return
			}
		}
	}
	fmt.Println("<================================================>")
	fmt.Println("|O-Owww Buku yang Dicari Tidak Ditemukan frenss!!|")
	fmt.Println("<================================================>")
}

func HistoriPeminjamanBuku() {
	if len(History) == 0 {
		fmt.Println("|       YAHHH Histori Peminjaman Belum Adaaa!    |")
		fmt.Println("<================================================>")
		return

	}

	fmt.Println("<===================================================>")
	fmt.Println("|              Histori Peminjaman Buku              |")
	fmt.Println("<===================================================>")
	for _, entry := range History {
		fmt.Println(entry)
	}
	fmt.Println("<===================================================>")
}
