package helpers

import "fmt"

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

// function pakai pointer change nama jadi budi
func Ubah(pegawai *Pegawai) {
	pegawai.Nama = "Budi"
}

// struct HERO 2
func Battle(attacker *Hero, defender *Hero) {
	fmt.Printf("%s attacks %s\n", attacker.Name, defender.Name)
	defender.IsAttackedBy(attacker)
	if defender.HealthPoint <= 0 {
		fmt.Printf("%s is dead\n", defender.Name)
	} else {
		fmt.Printf("%s is still alive with health point %d\n", defender.Name, defender.HealthPoint)
	}
}
