package fungsiA

func SortData(T *ArrKantin) {
	for i := 1; i < T.N; i++ {
		j := i
		for j > 0 && T.Data[j].Transaksi > T.Data[j-1].Transaksi {
			T.Data[j], T.Data[j-1] = T.Data[j-1], T.Data[j]
			j--
		}
	}
}
