package helpers

func TambahPegawai(ID int, Nama string, Umur int, Gaji float64) Pegawai {
	pegawai := Pegawai{
		ID:   ID,
		Nama: Nama,
		Umur: Umur,
		Gaji: Gaji,
	}
	return pegawai
}

func CariPegawai(ID int) *Pegawai {
	pegawai := Pegawai{
		ID: ID,
	}
	return &pegawai
}
