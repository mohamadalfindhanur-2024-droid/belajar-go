package main

import "fmt"

func main() {
	// Deklarasi variabel dasar
	var nama string = "Alfin"
	var umur int = 20
	var ipk float64 = 3.85
	var statusAktif bool = true
	var hobi = []string{"Coding", "Naik Gunung"}

	fmt.Println("=== Data Mahasiswa ===")
	fmt.Println("Nama:", nama)
	fmt.Println("Umur:", umur)
	fmt.Println("IPK:", ipk)
	fmt.Println("Status Aktif:", statusAktif)
	fmt.Println("Hobi:", hobi)
	fmt.Println()

	// Inisialisasi map nilai
	nilaiMhs := map[string]int{
		"Budi": 85,
		"Siti": 90,
	}

	// Tambah data baru
	nilaiMhs["Andi"] = 78

	// Cek keberadaan data
	nilai, ada := nilaiMhs["Siti"]
	if ada {
		fmt.Println("Nilai Siti adalah:", nilai)
	}

	// Hapus data
	delete(nilaiMhs, "Budi")

	// Tampilkan seluruh isi map
	fmt.Println("\nDaftar Nilai Akhir:")
	for k, v := range nilaiMhs {
		fmt.Printf("- %s : %d\n", k, v)
	}
}
