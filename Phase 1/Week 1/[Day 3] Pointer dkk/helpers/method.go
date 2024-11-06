package helpers

func (pegawai *Pegawai) NaikkanGaji(persen float64) {
	pegawai.Gaji += (pegawai.Gaji * persen / 100)
}

//NON GRADED CHALLENGE

//struct and method

func (person *Person) GetInfo() {
	person.Name = "Bambang"
	person.Age = 20
	person.Job = "Gambler"
}

func AddYear(age *Person) {
	age.Age += 1
	if age.Age >= 50 {
		age.Job = "Retired"
	}
}
