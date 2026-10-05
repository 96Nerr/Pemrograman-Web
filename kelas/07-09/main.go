package main
import "fmt"


func main() {
	
	// hasil := sapa("dimas", 10)
	// fmt.Println(hasil)

	// fmt.Println("=====================")

	// menghitung("kevin", 10)

	//struct adalah tipe data yang bisa menampung banyak data dengan tipe data berbeda-beda
	kambing := hewan{
		name: "miwa",
		hasWings: false,
		food: "rumput",
	}
	kucing := hewan{}
	kucing.name = "jio"
	kucing.hasWings = false
	kucing.food = "daging"

	println(kambing.name)
	println(kucing.food)


	//array adalah tipe data yang bisa menampung banyak data dengan tipe data sama
	//slice adalah array yang bisa menampung banyak data dengan tipe data sama, tapi bisa berubah-ubah ukurannya
	var list_hewan = []hewan{kambing, kucing}
	for i := range list_hewan {
		fmt.Println(i, list_hewan[i].name)
	}


	//map adalah tipe data yang bisa menampung banyak data dengan tipe data sama, tapi dengan key-value
	harga := map[string]int{
		"apel": 10000,
		"jeruk": 15000,
		"mangga": 20000,
	}
	for key, value := range harga {
		fmt.Println(key, value)
	}

}

type hewan struct {

	name string
	hasWings bool
	food string
}



func sapa(nama string, sifat int) string {
	
	return "hello "+ nama + " :) " + fmt.Sprint(sifat)
}

func menghitung(nama string, b int) string {
	for i := 1; i <= b; i++ {
		fmt.Println("hello "+ nama + " :) " + fmt.Sprint(i))
	}
	return "selesai"
}