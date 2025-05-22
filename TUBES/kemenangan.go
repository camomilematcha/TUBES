package main

import "fmt"

func MenuKemenangan() {
	var pilihan int
	for {
		fmt.Println("\n-- Menu Kemenangan --")
		fmt.Println("1. Tambah Hasil Pertandingan")
		fmt.Println("2. Urutkan Berdasarkan Wins")
		fmt.Println("3. Urutkan Berdasarkan Score Difference")
		fmt.Println("4. Tampilkan Tim Terbaik")
		fmt.Println("0. Kembali")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			var id1, id2, s1, s2 int
			fmt.Print("ID Tim 1: ")
			fmt.Scan(&id1)
			fmt.Print("Skor Tim 1: ")
			fmt.Scan(&s1)
			fmt.Print("ID Tim 2: ")
			fmt.Scan(&id2)
			fmt.Print("Skor Tim 2: ")
			fmt.Scan(&s2)
			i1 := BinarySearchByID(id1)
			i2 := BinarySearchByID(id2)
			if i1 == -1 || i2 == -1 {
				fmt.Println("Salah satu ID tidak ditemukan.")
				continue
			}
			if s1 > s2 {
				Teams[i1].Wins++
				Teams[i1].Points += 3
			} else if s2 > s1 {
				Teams[i2].Wins++
				Teams[i2].Points += 3
			} else {
				Teams[i1].Points++
				Teams[i2].Points++
			}
			Teams[i1].ScoreDiff += s1 - s2
			Teams[i2].ScoreDiff += s2 - s1
			fmt.Println("Hasil pertandingan ditambahkan.")
		
		case 2:
			var asc int
			fmt.Print("Ascending (1) / Descending (0): ")
			fmt.Scan(&asc)
			SelectionSortByWins(asc == 1)
			BubbleSortByPoints(asc == 1)
			fmt.Println("\nHasil Setelah Pengurutan Berdasarkan Wins dan Points:")
			for i := 0; i < TeamCount; i++ {
				t := Teams[i]
				fmt.Printf("ID: %d, Nama: %s, Wins: %d, Points: %d, ScoreDiff: %d\n", t.ID, t.Name, t.Wins, t.Points, t.ScoreDiff)
			}

		case 3:
			var asc int
			fmt.Print("Ascending (1) / Descending (0): ")
			fmt.Scan(&asc)
			InsertionSortByScoreDiff(asc == 1)
			BubbleSortByPoints(asc == 1)
			fmt.Println("\nHasil Setelah Pengurutan Berdasarkan ScoreDiff dan Points:")
			for i := 0; i < TeamCount; i++ {
				t := Teams[i]
				fmt.Printf("ID: %d, Nama: %s, Wins: %d, Points: %d, ScoreDiff: %d\n", t.ID, t.Name, t.Wins, t.Points, t.ScoreDiff)
			}

		case 4:
			if TeamCount == 0 {
				fmt.Println("Tidak ada tim.")
				continue
			}
			best := Teams[0]
			for i := 1; i < TeamCount; i++ {
				if Teams[i].Wins > best.Wins || (Teams[i].Wins == best.Wins && Teams[i].ScoreDiff > best.ScoreDiff) {
					best = Teams[i]
				}
			}
			fmt.Println("Tim Terbaik:", best)
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")

			fmt.Println("\nHasil Setelah Pengurutan:") 
			for i := 0; i < TeamCount; i++ { t := Teams[i] 
				fmt.Printf("ID: %d, Nama: %s, Wins: %d, Points: %d, ScoreDiff: %d\n", t.ID, t.Name, t.Wins, t.Points, t.ScoreDiff)
}
		}
	}
}
