package main

import "fmt"

func MenuCari() {
	var pilihan int
	for {
		fmt.Println("\n-- Menu Cari Tim --")
		fmt.Println("1. Tampilkan Semua Tim")
		fmt.Println("2. Cari Tim berdasarkan Nama")
		fmt.Println("3. Cari Tim berdasarkan ID")
		fmt.Println("0. Kembali")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			for i := 0; i < TeamCount; i++ {
				t := Teams[i]
				fmt.Println("ID:", t.ID)
				fmt.Println("Nama:", t.Name)
				fmt.Println("Region:", t.Region)
				fmt.Println("Points:", t.Points)
				fmt.Println("Wins:", t.Wins)
				fmt.Println("ScoreDiff:", t.ScoreDiff)
				fmt.Println("----------------")
			}
		case 2:
			var name string
			fmt.Print("Nama Tim: ")
			fmt.Scan(&name)
			idx := SequentialSearchByName(name)
			if idx != -1 {
				fmt.Println("Tim ditemukan:", Teams[idx])
			} else {
				fmt.Println("Tim tidak ditemukan.")
			}
		case 3:
			var id int
			fmt.Print("ID Tim: ")
			fmt.Scan(&id)
			idx := BinarySearchByID(id)
			if idx != -1 {
				fmt.Println("Tim ditemukan:", Teams[idx])
			} else {
				fmt.Println("Tim tidak ditemukan.")
			}
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
