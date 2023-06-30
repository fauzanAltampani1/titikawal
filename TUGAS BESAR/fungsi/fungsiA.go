package fungsiA

import "fmt"

const NMAX = 100

type Kantin struct {
	Nama      string
	Uang      int
	Transaksi int
}

type ArrKantin struct {
	Data [NMAX]Kantin
	N    int
}

func TambahDatatenant(T *ArrKantin) {
	n := T.N + 1
	i := T.N
	fmt.Print("Nama Tenant:")
	fmt.Scan(&T.Data[i].Nama)
	T.N = n
	fmt.Println("Mantap!,Data Tenant Telah ditambahkan")
	fmt.Println()
}

func TampilkanDataTenant(T *ArrKantin) {
	var pilih string

	fmt.Println("------Tampilkan Data Tenant------")
	for i := 0; i < T.N; i++ {
		fmt.Printf("No: %d\n", i+1)
		fmt.Printf("Nama Tenant: %s\n", T.Data[i].Nama)
		fmt.Printf("Jumlah Uang: %d\n", T.Data[i].Uang)
		fmt.Printf("Jumlah Transaksi: %d\n", T.Data[i].Transaksi)
		fmt.Println("-----------------------")
		fmt.Println()
	}

	if T.N == 0 {
		fmt.Println("Belum ada Tenant yang ditambahkan")
		fmt.Print("Kembali Ke Menu Sebelumnya ? (y):")
		fmt.Scan(&pilih)
		if pilih == "y" {
			fmt.Print("\033c")
			return
		} else {
			fmt.Println("Menu yang anda pilih tidak valid")
			fmt.Print("Kembali Ke Menu Sebelumnya ?(y):")
			fmt.Scan(&pilih)
			if pilih == "y" {
				fmt.Print("\033c")
				return

			}

		}

	}
}
