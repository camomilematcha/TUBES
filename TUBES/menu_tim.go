package main

import "fmt"

func MenuTim() {
	var pilihan int
	for {
		fmt.Println("\n-- Menu Tim --")
		fmt.Println("1. Tambah Tim")
		fmt.Println("2. Edit Tim")
		fmt.Println("3. Hapus Tim")
		fmt.Println("0. Kembali")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			var t Team
			fmt.Print("ID: ")
			fmt.Scan(&t.ID)
			fmt.Print("Nama: ")
			fmt.Scan(&t.Name)
			fmt.Print("Region: ")
			fmt.Scan(&t.Region)
			if TeamCount < MAX_TEAMS {
				Teams[TeamCount] = t
				TeamCount++
				fmt.Println("Tim berhasil ditambahkan.")
			} else {
				fmt.Println("Kapasitas tim penuh.")
			}
		case 2:
			var id int
			fmt.Print("ID Tim yang ingin diedit: ")
			fmt.Scan(&id)
			idx := BinarySearchByID(id)
			if idx == -1 {
				fmt.Println("Tim tidak ditemukan.")
				continue
			}
			var name, region string
			var points int
			fmt.Print("Nama baru: ")
			fmt.Scan(&name)
			fmt.Print("Region baru: ")
			fmt.Scan(&region)
			fmt.Print("Points baru: ")
			fmt.Scan(&points)
			Teams[idx].Name = name
			Teams[idx].Region = region
			Teams[idx].Points = points
			fmt.Println("Tim berhasil diperbarui.")
		case 3:
			var id int
			fmt.Print("ID Tim yang ingin dihapus: ")
			fmt.Scan(&id)
			idx := BinarySearchByID(id)
			if idx == -1 {
				fmt.Println("Tim tidak ditemukan.")
				continue
			}
			for i := idx; i < TeamCount-1; i++ {
				Teams[i] = Teams[i+1]
			}
			TeamCount--
			fmt.Println("Tim berhasil dihapus.")
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
