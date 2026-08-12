package main

import "fmt"

// Menukar nilai dua variabel dengan pointer
func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

// Menambahkan item baru ke slice lewat pointer
func updateSlice(slice *[]string, itemBaru string) {
	*slice = append(*slice, itemBaru)
}

// Pass by value (tidak mengubah nilai asli)
func tambahSatuValue(n int) {
	n = n + 1
}

// Pass by pointer (mengubah nilai asli)
func tambahSatuPointer(n *int) {
	*n = *n + 1
}

func main() {
	// Uji fungsi swap
	x, y := 10, 20
	fmt.Println("Sebelum swap:", x, y)
	swap(&x, &y)
	fmt.Println("Setelah swap :", x, y)
	fmt.Println()

	// Uji updateSlice
	daftarHobi := []string{"Coding", "Trekking"}
	fmt.Println("Slice awal :", daftarHobi)
	updateSlice(&daftarHobi, "Membaca")
	fmt.Println("Slice akhir:", daftarHobi)
	fmt.Println()

	// Uji pass by value vs pointer
	angka := 5
	fmt.Println("Nilai awal angka:", angka)

	tambahSatuValue(angka)
	fmt.Println("Setelah pass by value  :", angka)

	tambahSatuPointer(&angka)
	fmt.Println("Setelah pass by pointer:", angka)
}
