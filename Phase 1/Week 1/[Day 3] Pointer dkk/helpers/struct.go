package helpers

// Pegawai struct
type Pegawai struct {
	ID   int
	Nama string
	Umur int
	Gaji float64
}

//NON GRADED CHALLENGE

//STRUCT AND METHOD

type Person struct {
	Name string
	Age  int
	Job  string
}

//struct HERO 1

type Hero struct {
	Name           string
	BaseAttack     int
	Defence        int
	CriticalDamage int
	HealthPoint    int
	Weap           Weapon
}

type Weapon struct {
	Attack int
}
