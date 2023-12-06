package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// tambahBarang adalah fungsi untuk menambahkan barang ke daftar barang
func tambahBarang() {
	readerData := bufio.NewReader(os.Stdin)

fmt.Printf("┌────────────────────────────────────────────────┐\n")
fmt.Printf("│               TAMBAH NAMA BARANG               │\n")
fmt.Printf("└────────────────────────────────────────────────┘\n\n")
fmt.Printf(" Input Barang \n\n")


	// Mengambil input nama barang
	fmt.Print("    Masukan nama barang  : ")
	nama, _ := readerData.ReadString('\n')
	nama = strings.TrimSpace(nama)

	// Mengambil input harga barang
	fmt.Print("    Masukan harga barang : ")
	harga, _ := readerData.ReadString('\n')
	harga = strings.TrimSpace(harga)

	// Validasi harga
	if _, err := strconv.ParseFloat(harga, 64); err != nil {
		fmt.Println("Harga harus berupa angka.")
		return
	}

	// Mengambil input stok barang
	fmt.Print("    Masukan stok barang  : ")
	stok, _ := readerData.ReadString('\n')
	stok = strings.TrimSpace(stok)

	// Validasi stok
	if _, err := strconv.Atoi(stok); err != nil {
		fmt.Println("Stok harus berupa angka.")
		return
	}

	fmt.Printf("\n")

	fmt.Println("└────────────────────────────────────────────────┘")

	// memasukan data ke dalam slice
	dataBarang := []string{nama, harga, stok}

	// Membuka file atau membuat file baru dan menambahkan data barang
	file, err := os.OpenFile("data-barang.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Gagal membuka file:", err)
		return
	}
	// file ditutup setelah fungsinya selesai dieksekusi
	defer file.Close()

	// memisahkan data-data di dalam slice menggunakan |
	_, err = file.WriteString(strings.Join(dataBarang, " | ") + "\n")
	if err != nil {
		fmt.Println("Gagal menulis ke file:", err)
		return
	}

	// menampilkan pesan berhasil jika data berhasil ditambahkan ke dalam file txt
	fmt.Printf("Berhasil menambahkan barang dengan nama %s, harga %s, dan stok %s.\n", nama, harga, stok)
}

// tampilanMenu menampilkan tampilan menu aplikasi
func tampilanMenu() {
	fmt.Println("┌────────────────────────────────────────────────┐")
	fmt.Println("|             APLIKASI DAFTAR BELANJA            │")
	fmt.Println("├────────────────────────────────────────────────┤")
	fmt.Println("|   PILIHAN MENU                                 |")
	fmt.Println("│   1. Tambah Nama Barang                        │")
	fmt.Println("│   2. Cari Kategori                             │")
	fmt.Println("│   3. Hapus Kategori                            │")
	fmt.Println("│   4. Tambah Belanjaan                          │")
	fmt.Println("│   5. Cari Belanjaan                            │")
	fmt.Println("│   6. Hapus Belanjaan                           │")
	fmt.Println("│   7. Hitung Total Belanjaan                    │")
	fmt.Println("│   8. Keluar                                    │")
	fmt.Println("|                                                |")
	fmt.Println("└────────────────────────────────────────────────┘")
}

func main() {
	// Menambahkan pesan keluar yang lebih baik
	defer fmt.Println("Terima kasih telah menggunakan aplikasi. Semoga hari anda menyenangkan!")

	// Menampilkan menu
	tampilanMenu()

	// Simulasi pemilihan menu (gantilah ini sesuai dengan kebutuhan)
	pilihan := 1
	switch pilihan {
	case 1:
		tambahBarang()
	case 2:
		// Implementasi menu lainnya
	case 3:
		// Implementasi menu lainnya
	case 4:
		// Implementasi menu lainnya
	case 5:
		// Implementasi menu lainnya
	case 6:
		// Implementasi menu lainnya
	case 7:
		// Implementasi menu lainnya
	case 8:
		// Keluar dari program
		return
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}
