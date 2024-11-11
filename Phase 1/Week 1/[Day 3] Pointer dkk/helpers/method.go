package helpers

import (
	"fmt"
	"math/rand"
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
func CountDamage(damage *Hero) int {
	TotalDamage = damage.Weap.Attack + damage.BaseAttack
	//rand.Seed(time.Now().UnixNano())
	if rand.Intn(2) == 0 {
		TotalDamage += damage.CriticalDamage
		fmt.Printf("the total damage is %d with critical damage\n", TotalDamage)
	} else {
		fmt.Printf("the total damage is %d\n", TotalDamage)
	}
	return TotalDamage
}

// struct HERO 2
func (villain *Hero) IsAttackedBy(hero *Hero) {
	TotalDamage = CountDamage(hero) - villain.Defence
	if TotalDamage > villain.HealthPoint {
		villain.HealthPoint -= TotalDamage
		fmt.Printf("the total damage received is %d, the hero is dead\n", TotalDamage)
		fmt.Printf("the hero is dead, the remaining health point is %d\n", villain.HealthPoint)
	} else {
		fmt.Printf("the total damage received is %d\n", TotalDamage)
		fmt.Printf("the hero is still alive, the remaining health point is %d\n", villain.HealthPoint)
	}
}
