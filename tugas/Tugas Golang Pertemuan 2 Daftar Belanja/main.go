//Nama: Naufal Ridho Permana
//NIM: 2404130075

package main
import "fmt"

// Buat program "Daftar Belanja" dengan:
// - struct Item (Nama, Harga, Jumlah)
// - Slice untuk menyimpan item
// - Fungsi untuk tambah, hitung total, tampilkan
// - Validasi: harga dan jumlah harus > 0

type Item struct {
	Nama   string
	Harga  float64
	Jumlah int
}

func tambahItem(daftar *[]Item, nama string, harga float64, jumlah int) error {
	if harga <= 0 || jumlah <= 0 {
		err := fmt.Errorf("gagal menambah '%s': harga dan jumlah harus lebih dari 0", nama)
        fmt.Println(err)
        return err
	}

	newItem := Item{Nama: nama, Harga: harga, Jumlah: jumlah}
	*daftar = append(*daftar, newItem)
	return nil
}

func hitungTotal(daftar []Item) float64 {
	var total float64 = 0
	for _, item := range daftar {
		total += item.Harga * float64(item.Jumlah)
	}
	return total
}

func tampilkan(daftar []Item) {
	fmt.Println("\n=== Daftar Belanja ===")
	if len(daftar) == 0 {
		fmt.Println("Daftar belanja masih kosong.")
		return
	}

	for i, item := range daftar {
		subtotal := item.Harga * float64(item.Jumlah)
		fmt.Printf("%d. %s\n", i+1, item.Nama)
		fmt.Printf("   %d x Rp%.0f = Rp%.0f\n", item.Jumlah, item.Harga, subtotal)
	}

	fmt.Println()
	fmt.Printf("Total Keseluruhan: Rp%.0f\n", hitungTotal(daftar))
}

func main() {
	var daftar []Item

	tambahItem(&daftar, "Beras", 15000, -1)
	tambahItem(&daftar, "Gula", 10000, 1)
	tambahItem(&daftar, "Telur", 25000, 3)

	tampilkan(daftar)

	for _, item := range daftar {
		fmt.Println(item)
	}
}