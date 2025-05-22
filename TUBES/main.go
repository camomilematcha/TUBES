package main

import "fmt"

func main() {
	var pilihan int
	for {
		fmt.Println("\n=== MENU UTAMA E-SPORT ===")
		fmt.Println("1. Menu Tim")
		fmt.Println("2. Cari Tim")
		fmt.Println("3. Kemenangan")
		fmt.Println("4. Jadwal Pertandingan")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			MenuTim()
		case 2:
			MenuCari()
		case 3:
			MenuKemenangan()
		case 4:
			MenuJadwal()
		case 0:
			fmt.Println("Terima kasih.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
