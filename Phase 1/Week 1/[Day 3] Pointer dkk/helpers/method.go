package helpers

import (
	"fmt"
	"math/rand"
	"time"
)

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

var TotalDamage int

// struct HERO 1
func CountDamage(damage *Hero) {
	TotalDamage = damage.Weap.Attack + damage.BaseAttack
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(2) == 0 {
		TotalDamage += damage.CriticalDamage
		fmt.Printf("the total damage is %d with critical damage", TotalDamage)
	} else {
		fmt.Printf("the total damage is %d", TotalDamage)
	}
}
