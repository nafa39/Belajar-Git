package helpers

import "fmt"

func (pegawai *Pegawai) NaikkanGaji(persen float64) {
	pegawai.Gaji += (pegawai.Gaji * persen / 100)
}

//NON GRADED CHALLENGE

//struct and method

func GetInfo(person *Person) {
	fmt.Println(person.Name)
	fmt.Println(person.Age)
	fmt.Println(person.Job)
}

func AddYear(age *Person) {
	age.Age += 1
	if age.Age >= 50 {
		age.Job = "Retired"
	}
}
