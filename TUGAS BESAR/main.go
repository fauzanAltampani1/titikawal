package main

import (
	"fmt"
	FungsiA "kantin/fungsi"
)

func main() {
	FungsiA.Header()

	var choice int
	fmt.Print("Pilih angka menu: ")
	fmt.Scan(&choice)
	var structKantin FungsiA.ArrKantin

	for choice != 4 {
		fmt.Print("\033c")
		switch choice {
		case 1:
			// Tambahkan logika untuk menu 1 di sini
			FungsiA.Menu1(&structKantin)

			fmt.Println()
			fmt.Print("\033c")

			fmt.Println()

			FungsiA.Header()
		case 2:
			fmt.Println("Anda memilih menu 2: Tambahkan data transaksi")
			// Tambahkan logika untuk menu 2 di sini
			FungsiA.Menu2(&structKantin)
			fmt.Println()
			//fmt.Printf("\x1bc")
			fmt.Print("\033c")
			fmt.Println()

			FungsiA.Header()
		case 3:
			fmt.Println("Anda memilih menu 3: Total keuntungan admin")
			// Tambahkan logika untuk menu 3 di sini
			fmt.Println("Total Keuntungan Admin Saat ini", FungsiA.PendapatanAdmin(structKantin))
			FungsiA.Header()
		default:
			fmt.Println("Menu yang anda pilih tidak valid.")
		}
		fmt.Print("Pilih angka menu: ")
		fmt.Scan(&choice)
	}
	fmt.Println("Aplikasi ditutup")
}
