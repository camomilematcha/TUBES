package main

import "fmt"

func MenuJadwal() {
	var pilihan int
	for {
		fmt.Println("\n-- Menu Jadwal Pertandingan --")
		fmt.Println("1. Tambah Jadwal")
		fmt.Println("2. Tampilkan Jadwal")
		fmt.Println("0. Kembali")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan {
case 1:
	if MatchCount >= MAX_MATCHES {
		fmt.Println("Kapasitas penuh.")
		continue
	}
	var m Match
	fmt.Print("ID Match: ")
	fmt.Scan(&m.ID)
	fmt.Print("ID Tim 1: \n")
	fmt.Scan(&m.Team1ID)
	fmt.Print("ID Tim 2: ")
	fmt.Scan(&m.Team2ID)
	fmt.Print("Tanggal (YYYY-MM-DD): ")
	fmt.Scan(&m.Date)
	fmt.Print("Lokasi: ")
	fmt.Scan(&m.Location)
	if BinarySearchByID(m.Team1ID) == -1 || BinarySearchByID(m.Team2ID) == -1 {
		fmt.Println("ID tim tidak valid.")
		continue
	}
	Matches[MatchCount] = m
	MatchCount++
	fmt.Println("Jadwal ditambahkan.")

		case 2:
			for i := 0; i < MatchCount; i++ {
				m := Matches[i]
				fmt.Printf("Pertandingan %d: Tim %d vs Tim %d\n", i+1, m.Team1ID, m.Team2ID)
				fmt.Println("Tanggal:", m.Date)
				fmt.Println("Lokasi:", m.Location)
				fmt.Println("---------------------------")
			}
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
