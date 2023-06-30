package fungsiA

import (
	"fmt"
)

func Menu1(tenant *ArrKantin) {
	HeaderTenant()
	var Select int

	fmt.Print(" Pilih Menu:")
	fmt.Scan(&Select) // 1

	for Select != 4 {
		fmt.Print("\033c")
		if Select == 1 {
			TambahDatatenant(tenant)

			HeaderTenant()
		} else if Select == 2 {
			TampilkanDataTenant(tenant)

			HeaderTenant()
		} else if Select == 3 {
			fmt.Println("Sunting data tenant")
			TampilkanDataTenant(tenant)

			fmt.Print("Ubah atau Hapus Data atau kembali ke menu (1/2/3)? =")
			fmt.Scan(&Select)
			if Select == 1 {
				EditTenant(tenant)
				HeaderTenant()
			} else if Select == 2 {
				HapusDataTenant(tenant)
				HeaderTenant()
			} else if Select == 3 {
				HeaderTenant()
			}

		} else {
			fmt.Println("Menu yang anda pilih tidak valid.")
		}
		fmt.Print(" Pilih Menu:")
		fmt.Scan(&Select)
	}
}

func Menu2(tenant *ArrKantin) {
	HeaderTransaksi()
	var Pilih int
	fmt.Print(" Pilih Menu:")
	fmt.Scan(&Pilih) // 1

	for Pilih != 4 {
		fmt.Print("\033c")
		if Pilih == 1 {

			TambahTransaksi(tenant)
			HeaderTransaksi()
		} else if Pilih == 2 {
			TampilkanDataTenant(tenant)

			HeaderTransaksi()
		} else if Pilih == 3 {

			SortData(tenant)
			TampilkanDataTenant(tenant)
			HeaderTenant()

		} else {
			fmt.Println("Menu yang anda pilih tidak valid.")
		}
		fmt.Print(" Pilih Menu:")
		fmt.Scan(&Pilih)
	}
}
