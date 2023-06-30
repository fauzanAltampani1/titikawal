package fungsiA

import "fmt"

func TambahTransaksi(T *ArrKantin) {
	var JumlahTransaksi int
	var nama string

	fmt.Print("Masukkan Nama Tenant: \n")
	fmt.Scan(&nama)
	fmt.Print("Masukkan Pendapatan: \n")
	fmt.Scan(&JumlahTransaksi)
	i := SequentialSearch(*T, nama)

	if i != -1 {
		pendapatanTenant := JumlahTransaksi
		T.Data[i].Uang += pendapatanTenant
		T.Data[i].Transaksi++
		fmt.Println("Transaksi Berhasil Dicatat")

	} else {
		fmt.Println("Nama Tenant Tidak Ditemukan.")
	}

}

func PendapatanAdmin(T ArrKantin) float64 {
	var TotalPendapatan int = 0
	for i := 0; i < T.N; i++ {
		TotalPendapatan = TotalPendapatan + T.Data[i].Uang
	}
	return float64(TotalPendapatan) * 0.25
}
