package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipcode int
}

func main() {

	jim := person{
		firstName: "jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "mail@mail.com",
			zipcode: 12345,
		},
	}
	//===========Value type
	jim.print()
	jimPointer := &jim //&jim->address to jim struct in memory 
	jimPointer.updateName("jimmy")
	jim.print()
	//SHORTCUT
	jim.updateName("jimmy2")
	jim.print()

	//===========Reference type
	mySlice := []string{"hi", "there", "how", "you", "doing"}
	updateSlice(mySlice)
	fmt.Println(mySlice)
}

//value type, needs pointer
func (pp *person) updateName(newName string){
	(*pp).firstName = newName
}

func (p person) print(){
	fmt.Println("")
	fmt.Printf("%+v", p)
	fmt.Println("")
}

//Reference type, no need pointer
func updateSlice(s []string){
	s[0] = "bye"
}

