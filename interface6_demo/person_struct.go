package interface6_demo

type Person struct {
	name string
}

func (per Person) care(pet Pet) {
	pet.eat()
	pet.sleep()
}
