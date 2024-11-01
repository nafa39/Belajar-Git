package helpers

func (pegawai *Pegawai) NaikkanGaji(persen float64) {
	pegawai.Gaji += (pegawai.Gaji * persen / 100)
}
