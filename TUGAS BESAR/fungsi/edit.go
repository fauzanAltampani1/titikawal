package fungsiA

import "fmt"

func SequentialSearch(T ArrKantin, nama string) int {
	index := -1
	for i := 0; i < T.N; i++ {
		if nama == T.Data[i].Nama {
			index = i
		}

	}
	return index
}

func EditTenant(T *ArrKantin) {
	var nama string

	fmt.Println("Cari berdasarkan nama : ")
	fmt.Scan(&nama)
	i := SequentialSearch(*T, nama)
	//fmt.Println(T.Data[i].Nama)
	fmt.Print("Ubah Nama : ")
	fmt.Scan(&T.Data[i].Nama)

}

func HapusDataTenant(T *ArrKantin) {
	var pilih int
	fmt.Println("------Hapus Data Tenant------")
	TampilkanDataTenant(T)
	fmt.Print("Masukkan nomor tenant yang ingin dihapus: ")
	fmt.Scan(&pilih)
	if pilih >= 1 && pilih <= T.N {
		// Menggeser data ke kiri untuk menghapus tenant
		for i := pilih - 1; i < T.N-1; i++ {
			T.Data[i] = T.Data[i+1]
		}
		T.N--

		fmt.Println("Data tenant berhasil dihapus.")
	} else {
		fmt.Println("Nomor tenant tidak valid.")
	}

	fmt.Print("Kembali Ke Menu Sebelumnya? (1): ")
	fmt.Scan(&pilih)
	if pilih == 1 {
		fmt.Printf("\x1bc")

	}
}
